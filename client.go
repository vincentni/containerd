/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package containerd

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	containersapi "github.com/containerd/containerd/api/services/containers/v1"
	eventsapi "github.com/containerd/containerd/api/services/events/v1"
	imagesapi "github.com/containerd/containerd/api/services/images/v1"
	namespacesapi "github.com/containerd/containerd/api/services/namespaces/v1"
	"github.com/containerd/containerd/api/services/tasks/v1"
	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/defaults"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/events"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/pkg/dialer"
	"github.com/containerd/containerd/platforms"
	"github.com/containerd/typeurl/v2"
	"github.com/opencontainers/runtime-spec/specs-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	const prefix = "types.containerd.io"
	// register TypeUrls for commonly marshaled external types
	major := strconv.Itoa(specs.VersionMajor)
	typeurl.Register(&specs.Spec{}, prefix, "opencontainers/runtime-spec", major, "Spec")
	typeurl.Register(&specs.Process{}, prefix, "opencontainers/runtime-spec", major, "Process")
	typeurl.Register(&specs.LinuxResources{}, prefix, "opencontainers/runtime-spec", major, "LinuxResources")
	typeurl.Register(&specs.WindowsResources{}, prefix, "opencontainers/runtime-spec", major, "WindowsResources")
}

// New returns a new containerd client that is connected to the containerd
// instance provided by address
func New(address string, opts ...ClientOpt) (*Client, error) {
	var copts clientOpts
	for _, o := range opts {
		if err := o(&copts); err != nil {
			return nil, err
		}
	}
	if copts.timeout == 0 {
		copts.timeout = 10 * time.Second
	}

	c := &Client{
		defaultns: copts.defaultns,
	}

	if copts.defaultRuntime != "" {
		c.runtime = copts.defaultRuntime
	} else {
		c.runtime = defaults.DefaultRuntime
	}

	if copts.defaultPlatform != nil {
		c.platform = copts.defaultPlatform
	} else {
		c.platform = platforms.Default()
	}

	if copts.services != nil {
		c.services = *copts.services
	}
	if address != "" {
		backoffConfig := backoff.DefaultConfig
		backoffConfig.MaxDelay = 3 * time.Second
		connParams := grpc.ConnectParams{
			Backoff: backoffConfig,
		}
		gopts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.FailOnNonTempDialError(true),
			grpc.WithConnectParams(connParams),
			grpc.WithContextDialer(dialer.ContextDialer),
			grpc.WithReturnConnectionError(),
		}
		if len(copts.dialOptions) > 0 {
			gopts = copts.dialOptions
		}
		gopts = append(gopts, grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(defaults.DefaultMaxRecvMsgSize),
			grpc.MaxCallSendMsgSize(defaults.DefaultMaxSendMsgSize)))
		if len(copts.callOptions) > 0 {
			gopts = append(gopts, grpc.WithDefaultCallOptions(copts.callOptions...))
		}
		if copts.defaultns != "" {
			unary, stream := newNSInterceptors(copts.defaultns)
			gopts = append(gopts, grpc.WithChainUnaryInterceptor(unary))
			gopts = append(gopts, grpc.WithChainStreamInterceptor(stream))
		}

		connector := func() (*grpc.ClientConn, error) {
			ctx, cancel := context.WithTimeout(context.Background(), copts.timeout)
			defer cancel()
			conn, err := grpc.DialContext(ctx, dialer.DialAddress(address), gopts...)
			if err != nil {
				return nil, fmt.Errorf("failed to dial %q: %w", address, err)
			}
			return conn, nil
		}
		conn, err := connector()
		if err != nil {
			return nil, err
		}
		c.conn, c.connector = conn, connector
	}
	if copts.services == nil && c.conn == nil {
		return nil, fmt.Errorf("no grpc connection or services is available: %w", errdefs.ErrUnavailable)
	}

	// check namespace labels for default runtime
	if copts.defaultRuntime == "" && c.defaultns != "" {
		if label, err := c.GetLabel(context.Background(), defaults.DefaultRuntimeNSLabel); err != nil {
			return nil, err
		} else if label != "" {
			c.runtime = label
		}
	}

	return c, nil
}

// Client is the client to interact with containerd and its various services
// using a uniform interface
type Client struct {
	services
	connMu    sync.Mutex
	conn      *grpc.ClientConn
	runtime   string
	defaultns string
	platform  platforms.MatchComparer
	connector func() (*grpc.ClientConn, error)
}

// Reconnect re-establishes the GRPC connection to the containerd daemon
func (c *Client) Reconnect() error {
	if c.connector == nil {
		return fmt.Errorf("unable to reconnect to containerd, no connector available: %w", errdefs.ErrUnavailable)
	}
	c.connMu.Lock()
	defer c.connMu.Unlock()
	c.conn.Close()
	conn, err := c.connector()
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

// GetLabel gets a label value from namespace store
// If there is no default label, an empty string returned with nil error
func (c *Client) GetLabel(ctx context.Context, label string) (string, error) {
	ns, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		if c.defaultns == "" {
			return "", err
		}
		ns = c.defaultns
	}

	srv := c.NamespaceService()
	labels, err := srv.Labels(ctx, ns)
	if err != nil {
		return "", err
	}

	value := labels[label]
	return value, nil
}

// Subscribe to events that match one or more of the provided filters.
//
// Callers should listen on both the envelope and errs channels. If the errs
// channel returns nil or an error, the subscriber should terminate.
//
// The subscriber can stop receiving events by canceling the provided context.
// The errs channel will be closed and return a nil error.
func (c *Client) Subscribe(ctx context.Context, filters ...string) (ch <-chan *events.Envelope, errs <-chan error) {
	return c.EventService().Subscribe(ctx, filters...)
}

// Close closes the clients connection to containerd
func (c *Client) Close() error {
	c.connMu.Lock()
	defer c.connMu.Unlock()
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// NamespaceService returns the underlying Namespaces Store
func (c *Client) NamespaceService() namespaces.Store {
	if c.namespaceStore != nil {
		return c.namespaceStore
	}
	c.connMu.Lock()
	defer c.connMu.Unlock()
	return NewNamespaceStoreFromClient(namespacesapi.NewNamespacesClient(c.conn))
}

// ContainerService returns the underlying container Store
func (c *Client) ContainerService() containers.Store {
	if c.containerStore != nil {
		return c.containerStore
	}
	c.connMu.Lock()
	defer c.connMu.Unlock()
	return NewRemoteContainerStore(containersapi.NewContainersClient(c.conn))
}

// TaskService returns the underlying TasksClient
func (c *Client) TaskService() tasks.TasksClient {
	if c.taskService != nil {
		return c.taskService
	}
	c.connMu.Lock()
	defer c.connMu.Unlock()
	return tasks.NewTasksClient(c.conn)
}

// ImageService returns the underlying image Store
func (c *Client) ImageService() images.Store {
	if c.imageStore != nil {
		return c.imageStore
	}
	c.connMu.Lock()
	defer c.connMu.Unlock()
	return NewImageStoreFromClient(imagesapi.NewImagesClient(c.conn))
}

// EventService returns the underlying event service
func (c *Client) EventService() EventService {
	if c.eventService != nil {
		return c.eventService
	}
	c.connMu.Lock()
	defer c.connMu.Unlock()
	return NewEventServiceFromClient(eventsapi.NewEventsClient(c.conn))
}
