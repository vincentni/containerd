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
	"time"

	"github.com/containerd/containerd/platforms"
	"google.golang.org/grpc"
)

type clientOpts struct {
	defaultns       string
	defaultRuntime  string
	defaultPlatform platforms.MatchComparer
	services        *services
	dialOptions     []grpc.DialOption
	callOptions     []grpc.CallOption
	timeout         time.Duration
}

// ClientOpt allows callers to set options on the containerd client
type ClientOpt func(c *clientOpts) error
