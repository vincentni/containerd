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

	api "github.com/containerd/containerd/api/services/namespaces/v1"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/namespaces"
)

// NewNamespaceStoreFromClient returns a new namespace store
func NewNamespaceStoreFromClient(client api.NamespacesClient) namespaces.Store {
	return &remoteNamespaces{client: client}
}

type remoteNamespaces struct {
	client api.NamespacesClient
}

func (r *remoteNamespaces) Labels(ctx context.Context, namespace string) (map[string]string, error) {
	var req api.GetNamespaceRequest
	req.Name = namespace

	resp, err := r.client.Get(ctx, &req)
	if err != nil {
		return nil, errdefs.FromGRPC(err)
	}

	return resp.Namespace.Labels, nil
}

func (r *remoteNamespaces) List(ctx context.Context) ([]string, error) {
	var req api.ListNamespacesRequest

	resp, err := r.client.List(ctx, &req)
	if err != nil {
		return nil, errdefs.FromGRPC(err)
	}

	var namespaces []string

	for _, ns := range resp.Namespaces {
		namespaces = append(namespaces, ns.Name)
	}

	return namespaces, nil
}
