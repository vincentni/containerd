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

	eventsapi "github.com/containerd/containerd/api/services/events/v1"
	"github.com/containerd/containerd/events"
	"github.com/containerd/containerd/protobuf"
)

// EventService handles the publish, forward and subscribe of events.
type EventService interface {
	events.Subscriber
}

// NewEventServiceFromClient returns a new event service which communicates
// over a GRPC connection.
func NewEventServiceFromClient(client eventsapi.EventsClient) EventService {
	return &eventRemote{
		client: client,
	}
}

type eventRemote struct {
	client eventsapi.EventsClient
}

func (e *eventRemote) Subscribe(ctx context.Context, filters ...string) (ch <-chan *events.Envelope, errs <-chan error) {
	var (
		evq  = make(chan *events.Envelope)
		errq = make(chan error, 1)
	)

	errs = errq
	ch = evq

	session, err := e.client.Subscribe(ctx, &eventsapi.SubscribeRequest{
		Filters: filters,
	})
	if err != nil {
		errq <- err
		close(errq)
		return
	}

	go func() {
		defer close(errq)

		for {
			ev, err := session.Recv()
			if err != nil {
				errq <- err
				return
			}

			select {
			case evq <- &events.Envelope{
				Timestamp: protobuf.FromTimestamp(ev.Timestamp),
				Namespace: ev.Namespace,
				Topic:     ev.Topic,
				Event:     ev.Event,
			}:
			case <-ctx.Done():
				if cerr := ctx.Err(); cerr != context.Canceled {
					errq <- cerr
				}
				return
			}
		}
	}()

	return ch, errs
}
