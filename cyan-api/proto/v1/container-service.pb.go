// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: cyan-api/api/v1/container-service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContainerCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *UserIdentity `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name     string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Parent   *Container    `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	MetaData *Meta         `protobuf:"bytes,4,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
}

func (x *ContainerCreateRequest) Reset() {
	*x = ContainerCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerCreateRequest) ProtoMessage() {}

func (x *ContainerCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerCreateRequest.ProtoReflect.Descriptor instead.
func (*ContainerCreateRequest) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerCreateRequest) GetUserId() *UserIdentity {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ContainerCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerCreateRequest) GetParent() *Container {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *ContainerCreateRequest) GetMetaData() *Meta {
	if x != nil {
		return x.MetaData
	}
	return nil
}

type ContainerGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *UserIdentity `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Path     string        `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	MaxDepth int32         `protobuf:"varint,3,opt,name=max_depth,json=maxDepth,proto3" json:"max_depth,omitempty"`
}

func (x *ContainerGetRequest) Reset() {
	*x = ContainerGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerGetRequest) ProtoMessage() {}

func (x *ContainerGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerGetRequest.ProtoReflect.Descriptor instead.
func (*ContainerGetRequest) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerGetRequest) GetUserId() *UserIdentity {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ContainerGetRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ContainerGetRequest) GetMaxDepth() int32 {
	if x != nil {
		return x.MaxDepth
	}
	return 0
}

type ContainerUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  *UserIdentity `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Path    string        `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	NewName string        `protobuf:"bytes,3,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	NewMeta *Meta         `protobuf:"bytes,4,opt,name=new_meta,json=newMeta,proto3" json:"new_meta,omitempty"`
}

func (x *ContainerUpdateRequest) Reset() {
	*x = ContainerUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerUpdateRequest) ProtoMessage() {}

func (x *ContainerUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerUpdateRequest.ProtoReflect.Descriptor instead.
func (*ContainerUpdateRequest) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{2}
}

func (x *ContainerUpdateRequest) GetUserId() *UserIdentity {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ContainerUpdateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ContainerUpdateRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

func (x *ContainerUpdateRequest) GetNewMeta() *Meta {
	if x != nil {
		return x.NewMeta
	}
	return nil
}

type ContainerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *UserIdentity `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Path   string        `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ContainerDeleteRequest) Reset() {
	*x = ContainerDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerDeleteRequest) ProtoMessage() {}

func (x *ContainerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerDeleteRequest.ProtoReflect.Descriptor instead.
func (*ContainerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{3}
}

func (x *ContainerDeleteRequest) GetUserId() *UserIdentity {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *ContainerDeleteRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ContainerCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created   bool       `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Container *Container `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *ContainerCreateResponse) Reset() {
	*x = ContainerCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerCreateResponse) ProtoMessage() {}

func (x *ContainerCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerCreateResponse.ProtoReflect.Descriptor instead.
func (*ContainerCreateResponse) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{4}
}

func (x *ContainerCreateResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *ContainerCreateResponse) GetContainer() *Container {
	if x != nil {
		return x.Container
	}
	return nil
}

type ContainerGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Container *Container   `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	Inners    []*Container `protobuf:"bytes,2,rep,name=inners,proto3" json:"inners,omitempty"`
}

func (x *ContainerGetResponse) Reset() {
	*x = ContainerGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerGetResponse) ProtoMessage() {}

func (x *ContainerGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerGetResponse.ProtoReflect.Descriptor instead.
func (*ContainerGetResponse) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{5}
}

func (x *ContainerGetResponse) GetContainer() *Container {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *ContainerGetResponse) GetInners() []*Container {
	if x != nil {
		return x.Inners
	}
	return nil
}

type ContainerUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated   bool       `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
	Container *Container `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *ContainerUpdateResponse) Reset() {
	*x = ContainerUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerUpdateResponse) ProtoMessage() {}

func (x *ContainerUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerUpdateResponse.ProtoReflect.Descriptor instead.
func (*ContainerUpdateResponse) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{6}
}

func (x *ContainerUpdateResponse) GetUpdated() bool {
	if x != nil {
		return x.Updated
	}
	return false
}

func (x *ContainerUpdateResponse) GetContainer() *Container {
	if x != nil {
		return x.Container
	}
	return nil
}

type ContainerDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *ContainerDeleteResponse) Reset() {
	*x = ContainerDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerDeleteResponse) ProtoMessage() {}

func (x *ContainerDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerDeleteResponse.ProtoReflect.Descriptor instead.
func (*ContainerDeleteResponse) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{7}
}

func (x *ContainerDeleteResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MetaData *Meta  `protobuf:"bytes,3,opt,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[8]
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
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{8}
}

func (x *Container) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Container) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Container) GetMetaData() *Meta {
	if x != nil {
		return x.MetaData
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_cyan_api_api_v1_container_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_cyan_api_api_v1_container_service_proto_rawDescGZIP(), []int{9}
}

func (x *Meta) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_cyan_api_api_v1_container_service_proto protoreflect.FileDescriptor

var file_cyan_api_api_v1_container_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x79, 0x61, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x79, 0x61, 0x6e, 0x1a,
	0x22, 0x63, 0x79, 0x61, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x79,
	0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x73, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x79, 0x61, 0x6e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x44, 0x65, 0x70, 0x74, 0x68, 0x22, 0x9b, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x4d, 0x65, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x62, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x06, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x79, 0x61,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x06, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x73, 0x22, 0x62, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x79,
	0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x5c, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x69, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09,
	0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xd2, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x63, 0x79, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x79,
	0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e,
	0x63, 0x79, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x79, 0x61, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x79, 0x61, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x63, 0x79,
	0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x79, 0x61, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6e,
	0x44, 0x2f, 0x63, 0x79, 0x61, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63,
	0x79, 0x61, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cyan_api_api_v1_container_service_proto_rawDescOnce sync.Once
	file_cyan_api_api_v1_container_service_proto_rawDescData = file_cyan_api_api_v1_container_service_proto_rawDesc
)

func file_cyan_api_api_v1_container_service_proto_rawDescGZIP() []byte {
	file_cyan_api_api_v1_container_service_proto_rawDescOnce.Do(func() {
		file_cyan_api_api_v1_container_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cyan_api_api_v1_container_service_proto_rawDescData)
	})
	return file_cyan_api_api_v1_container_service_proto_rawDescData
}

var file_cyan_api_api_v1_container_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cyan_api_api_v1_container_service_proto_goTypes = []interface{}{
	(*ContainerCreateRequest)(nil),  // 0: cyan.ContainerCreateRequest
	(*ContainerGetRequest)(nil),     // 1: cyan.ContainerGetRequest
	(*ContainerUpdateRequest)(nil),  // 2: cyan.ContainerUpdateRequest
	(*ContainerDeleteRequest)(nil),  // 3: cyan.ContainerDeleteRequest
	(*ContainerCreateResponse)(nil), // 4: cyan.ContainerCreateResponse
	(*ContainerGetResponse)(nil),    // 5: cyan.ContainerGetResponse
	(*ContainerUpdateResponse)(nil), // 6: cyan.ContainerUpdateResponse
	(*ContainerDeleteResponse)(nil), // 7: cyan.ContainerDeleteResponse
	(*Container)(nil),               // 8: cyan.Container
	(*Meta)(nil),                    // 9: cyan.Meta
	nil,                             // 10: cyan.Meta.MetaEntry
	(*UserIdentity)(nil),            // 11: cyan.UserIdentity
}
var file_cyan_api_api_v1_container_service_proto_depIdxs = []int32{
	11, // 0: cyan.ContainerCreateRequest.user_id:type_name -> cyan.UserIdentity
	8,  // 1: cyan.ContainerCreateRequest.parent:type_name -> cyan.Container
	9,  // 2: cyan.ContainerCreateRequest.meta_data:type_name -> cyan.Meta
	11, // 3: cyan.ContainerGetRequest.user_id:type_name -> cyan.UserIdentity
	11, // 4: cyan.ContainerUpdateRequest.user_id:type_name -> cyan.UserIdentity
	9,  // 5: cyan.ContainerUpdateRequest.new_meta:type_name -> cyan.Meta
	11, // 6: cyan.ContainerDeleteRequest.user_id:type_name -> cyan.UserIdentity
	8,  // 7: cyan.ContainerCreateResponse.container:type_name -> cyan.Container
	8,  // 8: cyan.ContainerGetResponse.container:type_name -> cyan.Container
	8,  // 9: cyan.ContainerGetResponse.inners:type_name -> cyan.Container
	8,  // 10: cyan.ContainerUpdateResponse.container:type_name -> cyan.Container
	9,  // 11: cyan.Container.meta_data:type_name -> cyan.Meta
	10, // 12: cyan.Meta.meta:type_name -> cyan.Meta.MetaEntry
	0,  // 13: cyan.ContainerService.CreateContainer:input_type -> cyan.ContainerCreateRequest
	1,  // 14: cyan.ContainerService.GetContainers:input_type -> cyan.ContainerGetRequest
	2,  // 15: cyan.ContainerService.UpdateContainer:input_type -> cyan.ContainerUpdateRequest
	3,  // 16: cyan.ContainerService.DeleteContainer:input_type -> cyan.ContainerDeleteRequest
	4,  // 17: cyan.ContainerService.CreateContainer:output_type -> cyan.ContainerCreateResponse
	5,  // 18: cyan.ContainerService.GetContainers:output_type -> cyan.ContainerGetResponse
	6,  // 19: cyan.ContainerService.UpdateContainer:output_type -> cyan.ContainerUpdateResponse
	7,  // 20: cyan.ContainerService.DeleteContainer:output_type -> cyan.ContainerDeleteResponse
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cyan_api_api_v1_container_service_proto_init() }
func file_cyan_api_api_v1_container_service_proto_init() {
	if File_cyan_api_api_v1_container_service_proto != nil {
		return
	}
	file_cyan_api_api_v1_user_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cyan_api_api_v1_container_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerCreateRequest); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerGetRequest); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerUpdateRequest); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerDeleteRequest); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerCreateResponse); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerGetResponse); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerUpdateResponse); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerDeleteResponse); i {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_cyan_api_api_v1_container_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_cyan_api_api_v1_container_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cyan_api_api_v1_container_service_proto_goTypes,
		DependencyIndexes: file_cyan_api_api_v1_container_service_proto_depIdxs,
		MessageInfos:      file_cyan_api_api_v1_container_service_proto_msgTypes,
	}.Build()
	File_cyan_api_api_v1_container_service_proto = out.File
	file_cyan_api_api_v1_container_service_proto_rawDesc = nil
	file_cyan_api_api_v1_container_service_proto_goTypes = nil
	file_cyan_api_api_v1_container_service_proto_depIdxs = nil
}
