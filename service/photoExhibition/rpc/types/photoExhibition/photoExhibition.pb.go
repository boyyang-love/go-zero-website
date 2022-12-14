// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: photoExhibition.proto

package photoExhibition

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

// 创建
type CreatePhotoExhibitionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	SubTitle string `protobuf:"bytes,2,opt,name=SubTitle,proto3" json:"SubTitle,omitempty"`
	Des      string `protobuf:"bytes,3,opt,name=Des,proto3" json:"Des,omitempty"`
	Cover    string `protobuf:"bytes,4,opt,name=Cover,proto3" json:"Cover,omitempty"`
	UserId   int32  `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *CreatePhotoExhibitionReq) Reset() {
	*x = CreatePhotoExhibitionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhotoExhibitionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhotoExhibitionReq) ProtoMessage() {}

func (x *CreatePhotoExhibitionReq) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhotoExhibitionReq.ProtoReflect.Descriptor instead.
func (*CreatePhotoExhibitionReq) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePhotoExhibitionReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePhotoExhibitionReq) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *CreatePhotoExhibitionReq) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *CreatePhotoExhibitionReq) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreatePhotoExhibitionReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreatePhotoExhibitionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Id   int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *CreatePhotoExhibitionRes) Reset() {
	*x = CreatePhotoExhibitionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePhotoExhibitionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePhotoExhibitionRes) ProtoMessage() {}

func (x *CreatePhotoExhibitionRes) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePhotoExhibitionRes.ProtoReflect.Descriptor instead.
func (*CreatePhotoExhibitionRes) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePhotoExhibitionRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreatePhotoExhibitionRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreatePhotoExhibitionRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 修改
type EditPhotoExhibitionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	SubTitle string `protobuf:"bytes,3,opt,name=SubTitle,proto3" json:"SubTitle,omitempty"`
	Des      string `protobuf:"bytes,4,opt,name=Des,proto3" json:"Des,omitempty"`
}

func (x *EditPhotoExhibitionReq) Reset() {
	*x = EditPhotoExhibitionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditPhotoExhibitionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditPhotoExhibitionReq) ProtoMessage() {}

func (x *EditPhotoExhibitionReq) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditPhotoExhibitionReq.ProtoReflect.Descriptor instead.
func (*EditPhotoExhibitionReq) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{2}
}

func (x *EditPhotoExhibitionReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditPhotoExhibitionReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditPhotoExhibitionReq) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *EditPhotoExhibitionReq) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

type EditPhotoExhibitionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *EditPhotoExhibitionRes) Reset() {
	*x = EditPhotoExhibitionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditPhotoExhibitionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditPhotoExhibitionRes) ProtoMessage() {}

func (x *EditPhotoExhibitionRes) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditPhotoExhibitionRes.ProtoReflect.Descriptor instead.
func (*EditPhotoExhibitionRes) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{3}
}

func (x *EditPhotoExhibitionRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EditPhotoExhibitionRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 删除
type DelPhotoExhibitionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DelPhotoExhibitionReq) Reset() {
	*x = DelPhotoExhibitionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelPhotoExhibitionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelPhotoExhibitionReq) ProtoMessage() {}

func (x *DelPhotoExhibitionReq) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelPhotoExhibitionReq.ProtoReflect.Descriptor instead.
func (*DelPhotoExhibitionReq) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{4}
}

func (x *DelPhotoExhibitionReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DelPhotoExhibitionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *DelPhotoExhibitionRes) Reset() {
	*x = DelPhotoExhibitionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelPhotoExhibitionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelPhotoExhibitionRes) ProtoMessage() {}

func (x *DelPhotoExhibitionRes) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelPhotoExhibitionRes.ProtoReflect.Descriptor instead.
func (*DelPhotoExhibitionRes) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{5}
}

func (x *DelPhotoExhibitionRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DelPhotoExhibitionRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 修改状态
type EditStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Status    int32  `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	RejectRes string `protobuf:"bytes,3,opt,name=RejectRes,proto3" json:"RejectRes,omitempty"`
}

func (x *EditStatusReq) Reset() {
	*x = EditStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditStatusReq) ProtoMessage() {}

func (x *EditStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditStatusReq.ProtoReflect.Descriptor instead.
func (*EditStatusReq) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{6}
}

func (x *EditStatusReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditStatusReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *EditStatusReq) GetRejectRes() string {
	if x != nil {
		return x.RejectRes
	}
	return ""
}

type EditStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *EditStatusRes) Reset() {
	*x = EditStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditStatusRes) ProtoMessage() {}

func (x *EditStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditStatusRes.ProtoReflect.Descriptor instead.
func (*EditStatusRes) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{7}
}

func (x *EditStatusRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EditStatusRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 获取详细信息
type PhotoExhibitionInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *PhotoExhibitionInfoReq) Reset() {
	*x = PhotoExhibitionInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoExhibitionInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoExhibitionInfoReq) ProtoMessage() {}

func (x *PhotoExhibitionInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoExhibitionInfoReq.ProtoReflect.Descriptor instead.
func (*PhotoExhibitionInfoReq) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{8}
}

func (x *PhotoExhibitionInfoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PhotoExhibitionInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Id        int32  `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
	Title     string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	SubTitle  string `protobuf:"bytes,5,opt,name=SubTitle,proto3" json:"SubTitle,omitempty"`
	Des       string `protobuf:"bytes,6,opt,name=Des,proto3" json:"Des,omitempty"`
	Cover     string `protobuf:"bytes,7,opt,name=Cover,proto3" json:"Cover,omitempty"`
	UserId    int32  `protobuf:"varint,8,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Status    int32  `protobuf:"varint,9,opt,name=Status,proto3" json:"Status,omitempty"`
	RejectRes string `protobuf:"bytes,10,opt,name=RejectRes,proto3" json:"RejectRes,omitempty"`
}

func (x *PhotoExhibitionInfoRes) Reset() {
	*x = PhotoExhibitionInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_photoExhibition_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoExhibitionInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoExhibitionInfoRes) ProtoMessage() {}

func (x *PhotoExhibitionInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_photoExhibition_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoExhibitionInfoRes.ProtoReflect.Descriptor instead.
func (*PhotoExhibitionInfoRes) Descriptor() ([]byte, []int) {
	return file_photoExhibition_proto_rawDescGZIP(), []int{9}
}

func (x *PhotoExhibitionInfoRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PhotoExhibitionInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PhotoExhibitionInfoRes) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PhotoExhibitionInfoRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PhotoExhibitionInfoRes) GetSubTitle() string {
	if x != nil {
		return x.SubTitle
	}
	return ""
}

func (x *PhotoExhibitionInfoRes) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *PhotoExhibitionInfoRes) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *PhotoExhibitionInfoRes) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PhotoExhibitionInfoRes) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PhotoExhibitionInfoRes) GetRejectRes() string {
	if x != nil {
		return x.RejectRes
	}
	return ""
}

var File_photoExhibition_proto protoreflect.FileDescriptor

var file_photoExhibition_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78,
	0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53,
	0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x16, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75, 0x62,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x44, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x16, 0x45, 0x64, 0x69, 0x74, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69,
	0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22,
	0x55, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x28, 0x0a,
	0x16, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0xf6, 0x01, 0x0a, 0x16, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x75, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x32, 0x86, 0x04, 0x0a, 0x0f, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45,
	0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68,
	0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x12, 0x4c, 0x0a, 0x0a, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x1e, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x67, 0x0a, 0x13, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45,
	0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45,
	0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x1a, 0x27, 0x2e, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_photoExhibition_proto_rawDescOnce sync.Once
	file_photoExhibition_proto_rawDescData = file_photoExhibition_proto_rawDesc
)

func file_photoExhibition_proto_rawDescGZIP() []byte {
	file_photoExhibition_proto_rawDescOnce.Do(func() {
		file_photoExhibition_proto_rawDescData = protoimpl.X.CompressGZIP(file_photoExhibition_proto_rawDescData)
	})
	return file_photoExhibition_proto_rawDescData
}

var file_photoExhibition_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_photoExhibition_proto_goTypes = []interface{}{
	(*CreatePhotoExhibitionReq)(nil), // 0: photoExhibition.CreatePhotoExhibitionReq
	(*CreatePhotoExhibitionRes)(nil), // 1: photoExhibition.CreatePhotoExhibitionRes
	(*EditPhotoExhibitionReq)(nil),   // 2: photoExhibition.EditPhotoExhibitionReq
	(*EditPhotoExhibitionRes)(nil),   // 3: photoExhibition.EditPhotoExhibitionRes
	(*DelPhotoExhibitionReq)(nil),    // 4: photoExhibition.DelPhotoExhibitionReq
	(*DelPhotoExhibitionRes)(nil),    // 5: photoExhibition.DelPhotoExhibitionRes
	(*EditStatusReq)(nil),            // 6: photoExhibition.EditStatusReq
	(*EditStatusRes)(nil),            // 7: photoExhibition.EditStatusRes
	(*PhotoExhibitionInfoReq)(nil),   // 8: photoExhibition.PhotoExhibitionInfoReq
	(*PhotoExhibitionInfoRes)(nil),   // 9: photoExhibition.PhotoExhibitionInfoRes
}
var file_photoExhibition_proto_depIdxs = []int32{
	0, // 0: photoExhibition.PhotoExhibition.CreatePhotoExhibition:input_type -> photoExhibition.CreatePhotoExhibitionReq
	2, // 1: photoExhibition.PhotoExhibition.EditPhotoExhibition:input_type -> photoExhibition.EditPhotoExhibitionReq
	4, // 2: photoExhibition.PhotoExhibition.DelPhotoExhibition:input_type -> photoExhibition.DelPhotoExhibitionReq
	6, // 3: photoExhibition.PhotoExhibition.EditStatus:input_type -> photoExhibition.EditStatusReq
	8, // 4: photoExhibition.PhotoExhibition.PhotoExhibitionInfo:input_type -> photoExhibition.PhotoExhibitionInfoReq
	1, // 5: photoExhibition.PhotoExhibition.CreatePhotoExhibition:output_type -> photoExhibition.CreatePhotoExhibitionRes
	3, // 6: photoExhibition.PhotoExhibition.EditPhotoExhibition:output_type -> photoExhibition.EditPhotoExhibitionRes
	5, // 7: photoExhibition.PhotoExhibition.DelPhotoExhibition:output_type -> photoExhibition.DelPhotoExhibitionRes
	7, // 8: photoExhibition.PhotoExhibition.EditStatus:output_type -> photoExhibition.EditStatusRes
	9, // 9: photoExhibition.PhotoExhibition.PhotoExhibitionInfo:output_type -> photoExhibition.PhotoExhibitionInfoRes
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_photoExhibition_proto_init() }
func file_photoExhibition_proto_init() {
	if File_photoExhibition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_photoExhibition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhotoExhibitionReq); i {
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
		file_photoExhibition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePhotoExhibitionRes); i {
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
		file_photoExhibition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditPhotoExhibitionReq); i {
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
		file_photoExhibition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditPhotoExhibitionRes); i {
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
		file_photoExhibition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelPhotoExhibitionReq); i {
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
		file_photoExhibition_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelPhotoExhibitionRes); i {
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
		file_photoExhibition_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditStatusReq); i {
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
		file_photoExhibition_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditStatusRes); i {
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
		file_photoExhibition_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoExhibitionInfoReq); i {
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
		file_photoExhibition_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoExhibitionInfoRes); i {
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
			RawDescriptor: file_photoExhibition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_photoExhibition_proto_goTypes,
		DependencyIndexes: file_photoExhibition_proto_depIdxs,
		MessageInfos:      file_photoExhibition_proto_msgTypes,
	}.Build()
	File_photoExhibition_proto = out.File
	file_photoExhibition_proto_rawDesc = nil
	file_photoExhibition_proto_goTypes = nil
	file_photoExhibition_proto_depIdxs = nil
}
