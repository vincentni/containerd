// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/events/v1/events.proto

package events

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsClient interface {
	// Subscribe to a stream of events, possibly returning only that match any
	// of the provided filters.
	//
	// Unlike many other methods in containerd, subscribers will get messages
	// from all namespaces unless otherwise specified. If this is not desired,
	// a filter can be provided in the format 'namespace==<namespace>' to
	// restrict the received events.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Events_SubscribeClient, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Events_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Events_ServiceDesc.Streams[0], "/containerd.services.events.v1.Events/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_SubscribeClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type eventsSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventsSubscribeClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
// All implementations must embed UnimplementedEventsServer
// for forward compatibility
type EventsServer interface {
	// Subscribe to a stream of events, possibly returning only that match any
	// of the provided filters.
	//
	// Unlike many other methods in containerd, subscribers will get messages
	// from all namespaces unless otherwise specified. If this is not desired,
	// a filter can be provided in the format 'namespace==<namespace>' to
	// restrict the received events.
	Subscribe(*SubscribeRequest, Events_SubscribeServer) error
	mustEmbedUnimplementedEventsServer()
}

// UnimplementedEventsServer must be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (UnimplementedEventsServer) Subscribe(*SubscribeRequest, Events_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventsServer) mustEmbedUnimplementedEventsServer() {}

// UnsafeEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServer will
// result in compilation errors.
type UnsafeEventsServer interface {
	mustEmbedUnimplementedEventsServer()
}

func RegisterEventsServer(s grpc.ServiceRegistrar, srv EventsServer) {
	s.RegisterService(&Events_ServiceDesc, srv)
}

func _Events_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Subscribe(m, &eventsSubscribeServer{stream})
}

type Events_SubscribeServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type eventsSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventsSubscribeServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

// Events_ServiceDesc is the grpc.ServiceDesc for Events service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Events_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.events.v1.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Events_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/containerd/containerd/api/services/events/v1/events.proto",
}
