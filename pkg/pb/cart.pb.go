// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/cart.proto

package pb

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

// Add to cart
type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddToCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddToCartResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddToCartResponce) Reset() {
	*x = AddToCartResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponce) ProtoMessage() {}

func (x *AddToCartResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponce.ProtoReflect.Descriptor instead.
func (*AddToCartResponce) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartResponce) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddToCartResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveFromCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *RemoveFromCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveFromcartResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveFromcartResponce) Reset() {
	*x = RemoveFromcartResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromcartResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromcartResponce) ProtoMessage() {}

func (x *RemoveFromcartResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromcartResponce.ProtoReflect.Descriptor instead.
func (*RemoveFromcartResponce) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFromcartResponce) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RemoveFromcartResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FindCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindCartRequest) Reset() {
	*x = FindCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCartRequest) ProtoMessage() {}

func (x *FindCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCartRequest.ProtoReflect.Descriptor instead.
func (*FindCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{4}
}

func (x *FindCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindCartData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	Total     int64 `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	Qty       int64 `protobuf:"varint,3,opt,name=Qty,proto3" json:"Qty,omitempty"`
}

func (x *FindCartData) Reset() {
	*x = FindCartData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCartData) ProtoMessage() {}

func (x *FindCartData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCartData.ProtoReflect.Descriptor instead.
func (*FindCartData) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{5}
}

func (x *FindCartData) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *FindCartData) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FindCartData) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type FindCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64           `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Total  int64           `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Data   []*FindCartData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *FindCartResponse) Reset() {
	*x = FindCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCartResponse) ProtoMessage() {}

func (x *FindCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCartResponse.ProtoReflect.Descriptor instead.
func (*FindCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{6}
}

func (x *FindCartResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindCartResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FindCartResponse) GetData() []*FindCartData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DeleteCartRequest) Reset() {
	*x = DeleteCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartRequest) ProtoMessage() {}

func (x *DeleteCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteCartResponse) Reset() {
	*x = DeleteCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartResponse) ProtoMessage() {}

func (x *DeleteCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteCartResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_cart_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCartResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteCartResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_pb_cart_proto protoreflect.FileDescriptor

var file_pkg_pb_cart_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x48, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4d, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x63, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x29, 0x0a,
	0x0f, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x51, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x51, 0x74, 0x79, 0x22, 0x7e,
	0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0x94, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x63, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_cart_proto_rawDescOnce sync.Once
	file_pkg_pb_cart_proto_rawDescData = file_pkg_pb_cart_proto_rawDesc
)

func file_pkg_pb_cart_proto_rawDescGZIP() []byte {
	file_pkg_pb_cart_proto_rawDescOnce.Do(func() {
		file_pkg_pb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_cart_proto_rawDescData)
	})
	return file_pkg_pb_cart_proto_rawDescData
}

var file_pkg_pb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_pb_cart_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),       // 0: cart.AddToCartRequest
	(*AddToCartResponce)(nil),      // 1: cart.AddToCartResponce
	(*RemoveFromCartRequest)(nil),  // 2: cart.RemoveFromCartRequest
	(*RemoveFromcartResponce)(nil), // 3: cart.RemoveFromcartResponce
	(*FindCartRequest)(nil),        // 4: cart.FindCartRequest
	(*FindCartData)(nil),           // 5: cart.FindCartData
	(*FindCartResponse)(nil),       // 6: cart.FindCartResponse
	(*DeleteCartRequest)(nil),      // 7: cart.DeleteCartRequest
	(*DeleteCartResponse)(nil),     // 8: cart.DeleteCartResponse
}
var file_pkg_pb_cart_proto_depIdxs = []int32{
	5, // 0: cart.FindCartResponse.data:type_name -> cart.FindCartData
	0, // 1: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	2, // 2: cart.CartService.RemoveFromCart:input_type -> cart.RemoveFromCartRequest
	4, // 3: cart.CartService.FindCart:input_type -> cart.FindCartRequest
	7, // 4: cart.CartService.DeleteCart:input_type -> cart.DeleteCartRequest
	1, // 5: cart.CartService.AddToCart:output_type -> cart.AddToCartResponce
	3, // 6: cart.CartService.RemoveFromCart:output_type -> cart.RemoveFromcartResponce
	6, // 7: cart.CartService.FindCart:output_type -> cart.FindCartResponse
	8, // 8: cart.CartService.DeleteCart:output_type -> cart.DeleteCartResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_pb_cart_proto_init() }
func file_pkg_pb_cart_proto_init() {
	if File_pkg_pb_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
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
		file_pkg_pb_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponce); i {
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
		file_pkg_pb_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFromCartRequest); i {
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
		file_pkg_pb_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFromcartResponce); i {
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
		file_pkg_pb_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCartRequest); i {
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
		file_pkg_pb_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCartData); i {
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
		file_pkg_pb_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCartResponse); i {
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
		file_pkg_pb_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartRequest); i {
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
		file_pkg_pb_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartResponse); i {
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
			RawDescriptor: file_pkg_pb_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_cart_proto_goTypes,
		DependencyIndexes: file_pkg_pb_cart_proto_depIdxs,
		MessageInfos:      file_pkg_pb_cart_proto_msgTypes,
	}.Build()
	File_pkg_pb_cart_proto = out.File
	file_pkg_pb_cart_proto_rawDesc = nil
	file_pkg_pb_cart_proto_goTypes = nil
	file_pkg_pb_cart_proto_depIdxs = nil
}
