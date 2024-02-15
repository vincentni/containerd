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

	containersapi "github.com/containerd/containerd/api/services/containers/v1"
	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/protobuf"
	"github.com/containerd/typeurl/v2"
)

type remoteContainers struct {
	client containersapi.ContainersClient
}

var _ containers.Store = &remoteContainers{}

// NewRemoteContainerStore returns the container Store connected with the provided client
func NewRemoteContainerStore(client containersapi.ContainersClient) containers.Store {
	return &remoteContainers{
		client: client,
	}
}

func (r *remoteContainers) Get(ctx context.Context, id string) (containers.Container, error) {
	resp, err := r.client.Get(ctx, &containersapi.GetContainerRequest{
		ID: id,
	})
	if err != nil {
		return containers.Container{}, errdefs.FromGRPC(err)
	}

	return containerFromProto(resp.Container), nil
}

func containerFromProto(containerpb *containersapi.Container) containers.Container {
	var runtime containers.RuntimeInfo
	if containerpb.Runtime != nil {
		runtime = containers.RuntimeInfo{
			Name:    containerpb.Runtime.Name,
			Options: containerpb.Runtime.Options,
		}
	}
	extensions := make(map[string]typeurl.Any)
	for k, v := range containerpb.Extensions {
		v := v
		extensions[k] = v
	}
	return containers.Container{
		ID:          containerpb.ID,
		Labels:      containerpb.Labels,
		Image:       containerpb.Image,
		Runtime:     runtime,
		Spec:        containerpb.Spec,
		Snapshotter: containerpb.Snapshotter,
		SnapshotKey: containerpb.SnapshotKey,
		CreatedAt:   protobuf.FromTimestamp(containerpb.CreatedAt),
		UpdatedAt:   protobuf.FromTimestamp(containerpb.UpdatedAt),
		Extensions:  extensions,
		SandboxID:   containerpb.Sandbox,
	}
}
