// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: pb/telephone.proto

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

type GetContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetContactRequest) Reset() {
	*x = GetContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactRequest) ProtoMessage() {}

func (x *GetContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactRequest.ProtoReflect.Descriptor instead.
func (*GetContactRequest) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{0}
}

func (x *GetContactRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type GetContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lastname string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Number   string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetContactResponse) Reset() {
	*x = GetContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactResponse) ProtoMessage() {}

func (x *GetContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactResponse.ProtoReflect.Descriptor instead.
func (*GetContactResponse) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{1}
}

func (x *GetContactResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetContactResponse) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *GetContactResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type ListContactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListContactsRequest) Reset() {
	*x = ListContactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContactsRequest) ProtoMessage() {}

func (x *ListContactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContactsRequest.ProtoReflect.Descriptor instead.
func (*ListContactsRequest) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{2}
}

type ListContactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lastname string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Number   string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ListContactsResponse) Reset() {
	*x = ListContactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContactsResponse) ProtoMessage() {}

func (x *ListContactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContactsResponse.ProtoReflect.Descriptor instead.
func (*ListContactsResponse) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{3}
}

func (x *ListContactsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListContactsResponse) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *ListContactsResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type RecordCallHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *RecordCallHistoryRequest) Reset() {
	*x = RecordCallHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordCallHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordCallHistoryRequest) ProtoMessage() {}

func (x *RecordCallHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordCallHistoryRequest.ProtoReflect.Descriptor instead.
func (*RecordCallHistoryRequest) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{4}
}

func (x *RecordCallHistoryRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type RecordCallHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallCount   int32 `protobuf:"varint,1,opt,name=call_count,json=callCount,proto3" json:"call_count,omitempty"`
	ElapsedTime int32 `protobuf:"varint,2,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *RecordCallHistoryResponse) Reset() {
	*x = RecordCallHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordCallHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordCallHistoryResponse) ProtoMessage() {}

func (x *RecordCallHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordCallHistoryResponse.ProtoReflect.Descriptor instead.
func (*RecordCallHistoryResponse) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{5}
}

func (x *RecordCallHistoryResponse) GetCallCount() int32 {
	if x != nil {
		return x.CallCount
	}
	return 0
}

func (x *RecordCallHistoryResponse) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{6}
}

func (x *SendMessageRequest) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_telephone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_telephone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_pb_telephone_proto_rawDescGZIP(), []int{7}
}

func (x *SendMessageResponse) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_pb_telephone_proto protoreflect.FileDescriptor

var file_pb_telephone_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x18, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x5d,
	0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x27, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa5,
	0x02, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x52,
	0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43,
	0x61, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x61, 0x6c,
	0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x73, 0x6f, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_telephone_proto_rawDescOnce sync.Once
	file_pb_telephone_proto_rawDescData = file_pb_telephone_proto_rawDesc
)

func file_pb_telephone_proto_rawDescGZIP() []byte {
	file_pb_telephone_proto_rawDescOnce.Do(func() {
		file_pb_telephone_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_telephone_proto_rawDescData)
	})
	return file_pb_telephone_proto_rawDescData
}

var file_pb_telephone_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_telephone_proto_goTypes = []interface{}{
	(*GetContactRequest)(nil),         // 0: pb.GetContactRequest
	(*GetContactResponse)(nil),        // 1: pb.GetContactResponse
	(*ListContactsRequest)(nil),       // 2: pb.ListContactsRequest
	(*ListContactsResponse)(nil),      // 3: pb.ListContactsResponse
	(*RecordCallHistoryRequest)(nil),  // 4: pb.RecordCallHistoryRequest
	(*RecordCallHistoryResponse)(nil), // 5: pb.RecordCallHistoryResponse
	(*SendMessageRequest)(nil),        // 6: pb.SendMessageRequest
	(*SendMessageResponse)(nil),       // 7: pb.SendMessageResponse
}
var file_pb_telephone_proto_depIdxs = []int32{
	0, // 0: pb.Telephone.GetContact:input_type -> pb.GetContactRequest
	2, // 1: pb.Telephone.ListContacts:input_type -> pb.ListContactsRequest
	4, // 2: pb.Telephone.RecordCallHistory:input_type -> pb.RecordCallHistoryRequest
	6, // 3: pb.Telephone.SendMessage:input_type -> pb.SendMessageRequest
	1, // 4: pb.Telephone.GetContact:output_type -> pb.GetContactResponse
	3, // 5: pb.Telephone.ListContacts:output_type -> pb.ListContactsResponse
	5, // 6: pb.Telephone.RecordCallHistory:output_type -> pb.RecordCallHistoryResponse
	7, // 7: pb.Telephone.SendMessage:output_type -> pb.SendMessageResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_telephone_proto_init() }
func file_pb_telephone_proto_init() {
	if File_pb_telephone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_telephone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactRequest); i {
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
		file_pb_telephone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactResponse); i {
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
		file_pb_telephone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContactsRequest); i {
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
		file_pb_telephone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContactsResponse); i {
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
		file_pb_telephone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordCallHistoryRequest); i {
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
		file_pb_telephone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordCallHistoryResponse); i {
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
		file_pb_telephone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_pb_telephone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
			RawDescriptor: file_pb_telephone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_telephone_proto_goTypes,
		DependencyIndexes: file_pb_telephone_proto_depIdxs,
		MessageInfos:      file_pb_telephone_proto_msgTypes,
	}.Build()
	File_pb_telephone_proto = out.File
	file_pb_telephone_proto_rawDesc = nil
	file_pb_telephone_proto_goTypes = nil
	file_pb_telephone_proto_depIdxs = nil
}
