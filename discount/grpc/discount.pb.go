// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.26.1
// source: discount.proto

package grpc

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

type GetDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
}

func (x *GetDiscountRequest) Reset() {
	*x = GetDiscountRequest{}
	mi := &file_discount_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiscountRequest) ProtoMessage() {}

func (x *GetDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiscountRequest.ProtoReflect.Descriptor instead.
func (*GetDiscountRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{0}
}

func (x *GetDiscountRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type CouponModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      int32  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CouponModel) Reset() {
	*x = CouponModel{}
	mi := &file_discount_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CouponModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponModel) ProtoMessage() {}

func (x *CouponModel) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponModel.ProtoReflect.Descriptor instead.
func (*CouponModel) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{1}
}

func (x *CouponModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CouponModel) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *CouponModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CouponModel) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coupon *CouponModel `protobuf:"bytes,1,opt,name=coupon,proto3" json:"coupon,omitempty"`
}

func (x *CreateDiscountRequest) Reset() {
	*x = CreateDiscountRequest{}
	mi := &file_discount_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDiscountRequest) ProtoMessage() {}

func (x *CreateDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDiscountRequest.ProtoReflect.Descriptor instead.
func (*CreateDiscountRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDiscountRequest) GetCoupon() *CouponModel {
	if x != nil {
		return x.Coupon
	}
	return nil
}

type UpdateDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coupon *CouponModel `protobuf:"bytes,1,opt,name=coupon,proto3" json:"coupon,omitempty"`
}

func (x *UpdateDiscountRequest) Reset() {
	*x = UpdateDiscountRequest{}
	mi := &file_discount_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiscountRequest) ProtoMessage() {}

func (x *UpdateDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiscountRequest.ProtoReflect.Descriptor instead.
func (*UpdateDiscountRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDiscountRequest) GetCoupon() *CouponModel {
	if x != nil {
		return x.Coupon
	}
	return nil
}

type DeleteDiscountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string `protobuf:"bytes,1,opt,name=productName,proto3" json:"productName,omitempty"`
}

func (x *DeleteDiscountRequest) Reset() {
	*x = DeleteDiscountRequest{}
	mi := &file_discount_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDiscountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDiscountRequest) ProtoMessage() {}

func (x *DeleteDiscountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDiscountRequest.ProtoReflect.Descriptor instead.
func (*DeleteDiscountRequest) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDiscountRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type DeleteDiscountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteDiscountResponse) Reset() {
	*x = DeleteDiscountResponse{}
	mi := &file_discount_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDiscountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDiscountResponse) ProtoMessage() {}

func (x *DeleteDiscountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDiscountResponse.ProtoReflect.Descriptor instead.
func (*DeleteDiscountResponse) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDiscountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_discount_proto protoreflect.FileDescriptor

var file_discount_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x36, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x79,
	0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0x42, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x22, 0x39, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xa3, 0x02, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x40, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_proto_rawDescOnce sync.Once
	file_discount_proto_rawDescData = file_discount_proto_rawDesc
)

func file_discount_proto_rawDescGZIP() []byte {
	file_discount_proto_rawDescOnce.Do(func() {
		file_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_proto_rawDescData)
	})
	return file_discount_proto_rawDescData
}

var file_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_discount_proto_goTypes = []any{
	(*GetDiscountRequest)(nil),     // 0: grpc.GetDiscountRequest
	(*CouponModel)(nil),            // 1: grpc.CouponModel
	(*CreateDiscountRequest)(nil),  // 2: grpc.CreateDiscountRequest
	(*UpdateDiscountRequest)(nil),  // 3: grpc.UpdateDiscountRequest
	(*DeleteDiscountRequest)(nil),  // 4: grpc.DeleteDiscountRequest
	(*DeleteDiscountResponse)(nil), // 5: grpc.DeleteDiscountResponse
}
var file_discount_proto_depIdxs = []int32{
	1, // 0: grpc.CreateDiscountRequest.coupon:type_name -> grpc.CouponModel
	1, // 1: grpc.UpdateDiscountRequest.coupon:type_name -> grpc.CouponModel
	0, // 2: grpc.DiscountProtoService.GetDiscount:input_type -> grpc.GetDiscountRequest
	2, // 3: grpc.DiscountProtoService.CreateDiscount:input_type -> grpc.CreateDiscountRequest
	3, // 4: grpc.DiscountProtoService.UpdateDiscount:input_type -> grpc.UpdateDiscountRequest
	4, // 5: grpc.DiscountProtoService.DeleteDiscount:input_type -> grpc.DeleteDiscountRequest
	1, // 6: grpc.DiscountProtoService.GetDiscount:output_type -> grpc.CouponModel
	1, // 7: grpc.DiscountProtoService.CreateDiscount:output_type -> grpc.CouponModel
	1, // 8: grpc.DiscountProtoService.UpdateDiscount:output_type -> grpc.CouponModel
	5, // 9: grpc.DiscountProtoService.DeleteDiscount:output_type -> grpc.DeleteDiscountResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_discount_proto_init() }
func file_discount_proto_init() {
	if File_discount_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discount_proto_goTypes,
		DependencyIndexes: file_discount_proto_depIdxs,
		MessageInfos:      file_discount_proto_msgTypes,
	}.Build()
	File_discount_proto = out.File
	file_discount_proto_rawDesc = nil
	file_discount_proto_goTypes = nil
	file_discount_proto_depIdxs = nil
}