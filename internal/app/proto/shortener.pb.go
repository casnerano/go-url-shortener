// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: internal/app/server/grpc/proto/shortener.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetShortURLRequest) Reset() {
	*x = GetShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortURLRequest) ProtoMessage() {}

func (x *GetShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortURLRequest.ProtoReflect.Descriptor instead.
func (*GetShortURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *GetShortURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetShortURLResponse) Reset() {
	*x = GetShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShortURLResponse) ProtoMessage() {}

func (x *GetShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShortURLResponse.ProtoReflect.Descriptor instead.
func (*GetShortURLResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetShortURLResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GetUserHistoryShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetUserHistoryShortURLResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUserHistoryShortURLResponse) Reset() {
	*x = GetUserHistoryShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserHistoryShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserHistoryShortURLResponse) ProtoMessage() {}

func (x *GetUserHistoryShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserHistoryShortURLResponse.ProtoReflect.Descriptor instead.
func (*GetUserHistoryShortURLResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserHistoryShortURLResponse) GetItems() []*GetUserHistoryShortURLResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetStatsShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links int64 `protobuf:"varint,1,opt,name=links,proto3" json:"links,omitempty"`
	Users int64 `protobuf:"varint,2,opt,name=users,proto3" json:"users,omitempty"`
}

func (x *GetStatsShortURLResponse) Reset() {
	*x = GetStatsShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsShortURLResponse) ProtoMessage() {}

func (x *GetStatsShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsShortURLResponse.ProtoReflect.Descriptor instead.
func (*GetStatsShortURLResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetStatsShortURLResponse) GetLinks() int64 {
	if x != nil {
		return x.Links
	}
	return 0
}

func (x *GetStatsShortURLResponse) GetUsers() int64 {
	if x != nil {
		return x.Users
	}
	return 0
}

type CreateShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateShortURLRequest) Reset() {
	*x = CreateShortURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLRequest) ProtoMessage() {}

func (x *CreateShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateShortURLRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *CreateShortURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateShortURLResponse) Reset() {
	*x = CreateShortURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLResponse) ProtoMessage() {}

func (x *CreateShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLResponse.ProtoReflect.Descriptor instead.
func (*CreateShortURLResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *CreateShortURLResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CreateBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CreateBatchRequest_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateBatchRequest) Reset() {
	*x = CreateBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchRequest) ProtoMessage() {}

func (x *CreateBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchRequest.ProtoReflect.Descriptor instead.
func (*CreateBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBatchRequest) GetItems() []*CreateBatchRequest_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CreateBatchResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateBatchResponse) Reset() {
	*x = CreateBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchResponse) ProtoMessage() {}

func (x *CreateBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchResponse.ProtoReflect.Descriptor instead.
func (*CreateBatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *CreateBatchResponse) GetItems() []*CreateBatchResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DeleteBatchRequest) Reset() {
	*x = DeleteBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBatchRequest) ProtoMessage() {}

func (x *DeleteBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBatchRequest.ProtoReflect.Descriptor instead.
func (*DeleteBatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBatchRequest) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetUserHistoryShortURLResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetUserHistoryShortURLResponse_Item) Reset() {
	*x = GetUserHistoryShortURLResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserHistoryShortURLResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserHistoryShortURLResponse_Item) ProtoMessage() {}

func (x *GetUserHistoryShortURLResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserHistoryShortURLResponse_Item.ProtoReflect.Descriptor instead.
func (*GetUserHistoryShortURLResponse_Item) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetUserHistoryShortURLResponse_Item) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *GetUserHistoryShortURLResponse_Item) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateBatchRequest_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CreateBatchRequest_Item) Reset() {
	*x = CreateBatchRequest_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchRequest_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchRequest_Item) ProtoMessage() {}

func (x *CreateBatchRequest_Item) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchRequest_Item.ProtoReflect.Descriptor instead.
func (*CreateBatchRequest_Item) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{6, 0}
}

func (x *CreateBatchRequest_Item) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateBatchRequest_Item) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateBatchResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CreateBatchResponse_Item) Reset() {
	*x = CreateBatchResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchResponse_Item) ProtoMessage() {}

func (x *CreateBatchResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_server_grpc_proto_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchResponse_Item.ProtoReflect.Descriptor instead.
func (*CreateBatchResponse_Item) Descriptor() ([]byte, []int) {
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP(), []int{7, 0}
}

func (x *CreateBatchResponse_Item) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *CreateBatchResponse_Item) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

var File_internal_app_server_grpc_proto_shortener_proto protoreflect.FileDescriptor

var file_internal_app_server_grpc_proto_shortener_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x46, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x46, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x29, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x30, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x50, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x95,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x4a, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x32, 0x96, 0x03, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x46, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x73, 0x6e, 0x65, 0x72,
	0x61, 0x6e, 0x6f, 0x2f, 0x67, 0x6f, 0x2d, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_server_grpc_proto_shortener_proto_rawDescOnce sync.Once
	file_internal_app_server_grpc_proto_shortener_proto_rawDescData = file_internal_app_server_grpc_proto_shortener_proto_rawDesc
)

func file_internal_app_server_grpc_proto_shortener_proto_rawDescGZIP() []byte {
	file_internal_app_server_grpc_proto_shortener_proto_rawDescOnce.Do(func() {
		file_internal_app_server_grpc_proto_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_server_grpc_proto_shortener_proto_rawDescData)
	})
	return file_internal_app_server_grpc_proto_shortener_proto_rawDescData
}

var file_internal_app_server_grpc_proto_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_app_server_grpc_proto_shortener_proto_goTypes = []interface{}{
	(*GetShortURLRequest)(nil),                  // 0: pb.GetShortURLRequest
	(*GetShortURLResponse)(nil),                 // 1: pb.GetShortURLResponse
	(*GetUserHistoryShortURLResponse)(nil),      // 2: pb.GetUserHistoryShortURLResponse
	(*GetStatsShortURLResponse)(nil),            // 3: pb.GetStatsShortURLResponse
	(*CreateShortURLRequest)(nil),               // 4: pb.CreateShortURLRequest
	(*CreateShortURLResponse)(nil),              // 5: pb.CreateShortURLResponse
	(*CreateBatchRequest)(nil),                  // 6: pb.CreateBatchRequest
	(*CreateBatchResponse)(nil),                 // 7: pb.CreateBatchResponse
	(*DeleteBatchRequest)(nil),                  // 8: pb.DeleteBatchRequest
	(*GetUserHistoryShortURLResponse_Item)(nil), // 9: pb.GetUserHistoryShortURLResponse.Item
	(*CreateBatchRequest_Item)(nil),             // 10: pb.CreateBatchRequest.Item
	(*CreateBatchResponse_Item)(nil),            // 11: pb.CreateBatchResponse.Item
	(*emptypb.Empty)(nil),                       // 12: google.protobuf.Empty
}
var file_internal_app_server_grpc_proto_shortener_proto_depIdxs = []int32{
	9,  // 0: pb.GetUserHistoryShortURLResponse.items:type_name -> pb.GetUserHistoryShortURLResponse.Item
	10, // 1: pb.CreateBatchRequest.items:type_name -> pb.CreateBatchRequest.Item
	11, // 2: pb.CreateBatchResponse.items:type_name -> pb.CreateBatchResponse.Item
	0,  // 3: pb.Shortener.Get:input_type -> pb.GetShortURLRequest
	12, // 4: pb.Shortener.GetUserHistory:input_type -> google.protobuf.Empty
	12, // 5: pb.Shortener.GetStats:input_type -> google.protobuf.Empty
	4,  // 6: pb.Shortener.CreateURL:input_type -> pb.CreateShortURLRequest
	6,  // 7: pb.Shortener.CreateBatch:input_type -> pb.CreateBatchRequest
	8,  // 8: pb.Shortener.DeleteBatch:input_type -> pb.DeleteBatchRequest
	1,  // 9: pb.Shortener.Get:output_type -> pb.GetShortURLResponse
	2,  // 10: pb.Shortener.GetUserHistory:output_type -> pb.GetUserHistoryShortURLResponse
	3,  // 11: pb.Shortener.GetStats:output_type -> pb.GetStatsShortURLResponse
	5,  // 12: pb.Shortener.CreateURL:output_type -> pb.CreateShortURLResponse
	7,  // 13: pb.Shortener.CreateBatch:output_type -> pb.CreateBatchResponse
	12, // 14: pb.Shortener.DeleteBatch:output_type -> google.protobuf.Empty
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_internal_app_server_grpc_proto_shortener_proto_init() }
func file_internal_app_server_grpc_proto_shortener_proto_init() {
	if File_internal_app_server_grpc_proto_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortURLRequest); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShortURLResponse); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserHistoryShortURLResponse); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsShortURLResponse); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLRequest); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortURLResponse); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchRequest); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchResponse); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBatchRequest); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserHistoryShortURLResponse_Item); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchRequest_Item); i {
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
		file_internal_app_server_grpc_proto_shortener_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchResponse_Item); i {
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
			RawDescriptor: file_internal_app_server_grpc_proto_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_server_grpc_proto_shortener_proto_goTypes,
		DependencyIndexes: file_internal_app_server_grpc_proto_shortener_proto_depIdxs,
		MessageInfos:      file_internal_app_server_grpc_proto_shortener_proto_msgTypes,
	}.Build()
	File_internal_app_server_grpc_proto_shortener_proto = out.File
	file_internal_app_server_grpc_proto_shortener_proto_rawDesc = nil
	file_internal_app_server_grpc_proto_shortener_proto_goTypes = nil
	file_internal_app_server_grpc_proto_shortener_proto_depIdxs = nil
}