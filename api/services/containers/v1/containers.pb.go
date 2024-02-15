//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/services/containers/v1/containers.proto

package containers

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the user-specified identifier.
	//
	// This field may not be updated.
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Labels provides an area to include arbitrary data on containers.
	//
	// The combined size of a key/value pair cannot exceed 4096 bytes.
	//
	// Note that to add a new value to this field, read the existing set and
	// include the entire result in the update call.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Image contains the reference of the image used to build the
	// specification and snapshots for running this container.
	//
	// If this field is updated, the spec and rootfs needed to updated, as well.
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	// Runtime specifies which runtime to use for executing this container.
	Runtime *Container_Runtime `protobuf:"bytes,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// Spec to be used when creating the container. This is runtime specific.
	Spec *anypb.Any `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	// Snapshotter specifies the snapshotter name used for rootfs
	Snapshotter string `protobuf:"bytes,6,opt,name=snapshotter,proto3" json:"snapshotter,omitempty"`
	// SnapshotKey specifies the snapshot key to use for the container's root
	// filesystem. When starting a task from this container, a caller should
	// look up the mounts from the snapshot service and include those on the
	// task create request.
	//
	// Snapshots referenced in this field will not be garbage collected.
	//
	// This field is set to empty when the rootfs is not a snapshot.
	//
	// This field may be updated.
	SnapshotKey string `protobuf:"bytes,7,opt,name=snapshot_key,json=snapshotKey,proto3" json:"snapshot_key,omitempty"`
	// CreatedAt is the time the container was first created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UpdatedAt is the last time the container was mutated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Extensions allow clients to provide zero or more blobs that are directly
	// associated with the container. One may provide protobuf, json, or other
	// encoding formats. The primary use of this is to further decorate the
	// container object with fields that may be specific to a client integration.
	//
	// The key portion of this map should identify a "name" for the extension
	// that should be unique against other extensions. When updating extension
	// data, one should only update the specified extension using field paths
	// to select a specific map key.
	Extensions map[string]*anypb.Any `protobuf:"bytes,10,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Sandbox ID this container belongs to.
	Sandbox string `protobuf:"bytes,11,opt,name=sandbox,proto3" json:"sandbox,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescGZIP(), []int{0}
}

func (x *Container) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Container) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Container) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Container) GetRuntime() *Container_Runtime {
	if x != nil {
		return x.Runtime
	}
	return nil
}

func (x *Container) GetSpec() *anypb.Any {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Container) GetSnapshotter() string {
	if x != nil {
		return x.Snapshotter
	}
	return ""
}

func (x *Container) GetSnapshotKey() string {
	if x != nil {
		return x.SnapshotKey
	}
	return ""
}

func (x *Container) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Container) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Container) GetExtensions() map[string]*anypb.Any {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *Container) GetSandbox() string {
	if x != nil {
		return x.Sandbox
	}
	return ""
}

type GetContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetContainerRequest) Reset() {
	*x = GetContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainerRequest) ProtoMessage() {}

func (x *GetContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainerRequest.ProtoReflect.Descriptor instead.
func (*GetContainerRequest) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescGZIP(), []int{1}
}

func (x *GetContainerRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Container *Container `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *GetContainerResponse) Reset() {
	*x = GetContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContainerResponse) ProtoMessage() {}

func (x *GetContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContainerResponse.ProtoReflect.Descriptor instead.
func (*GetContainerResponse) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescGZIP(), []int{2}
}

func (x *GetContainerResponse) GetContainer() *Container {
	if x != nil {
		return x.Container
	}
	return nil
}

type Container_Runtime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the runtime.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Options specify additional runtime initialization options.
	Options *anypb.Any `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Container_Runtime) Reset() {
	*x = Container_Runtime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container_Runtime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container_Runtime) ProtoMessage() {}

func (x *Container_Runtime) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container_Runtime.ProtoReflect.Descriptor instead.
func (*Container_Runtime) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Container_Runtime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container_Runtime) GetOptions() *anypb.Any {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_github_com_containerd_containerd_api_services_containers_v1_containers_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x06,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x50, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5c, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x07, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x53, 0x0a, 0x0f, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x32, 0x84, 0x01, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x76, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescData = file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDesc
)

func file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDescData
}

var file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_goTypes = []interface{}{
	(*Container)(nil),             // 0: containerd.services.containers.v1.Container
	(*GetContainerRequest)(nil),   // 1: containerd.services.containers.v1.GetContainerRequest
	(*GetContainerResponse)(nil),  // 2: containerd.services.containers.v1.GetContainerResponse
	nil,                           // 3: containerd.services.containers.v1.Container.LabelsEntry
	(*Container_Runtime)(nil),     // 4: containerd.services.containers.v1.Container.Runtime
	nil,                           // 5: containerd.services.containers.v1.Container.ExtensionsEntry
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_depIdxs = []int32{
	3,  // 0: containerd.services.containers.v1.Container.labels:type_name -> containerd.services.containers.v1.Container.LabelsEntry
	4,  // 1: containerd.services.containers.v1.Container.runtime:type_name -> containerd.services.containers.v1.Container.Runtime
	6,  // 2: containerd.services.containers.v1.Container.spec:type_name -> google.protobuf.Any
	7,  // 3: containerd.services.containers.v1.Container.created_at:type_name -> google.protobuf.Timestamp
	7,  // 4: containerd.services.containers.v1.Container.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 5: containerd.services.containers.v1.Container.extensions:type_name -> containerd.services.containers.v1.Container.ExtensionsEntry
	0,  // 6: containerd.services.containers.v1.GetContainerResponse.container:type_name -> containerd.services.containers.v1.Container
	6,  // 7: containerd.services.containers.v1.Container.Runtime.options:type_name -> google.protobuf.Any
	6,  // 8: containerd.services.containers.v1.Container.ExtensionsEntry.value:type_name -> google.protobuf.Any
	1,  // 9: containerd.services.containers.v1.Containers.Get:input_type -> containerd.services.containers.v1.GetContainerRequest
	2,  // 10: containerd.services.containers.v1.Containers.Get:output_type -> containerd.services.containers.v1.GetContainerResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_init() }
func file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_init() {
	if File_github_com_containerd_containerd_api_services_containers_v1_containers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContainerResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container_Runtime); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_services_containers_v1_containers_proto = out.File
	file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_goTypes = nil
	file_github_com_containerd_containerd_api_services_containers_v1_containers_proto_depIdxs = nil
}
