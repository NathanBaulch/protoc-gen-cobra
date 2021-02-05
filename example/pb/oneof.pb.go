// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: oneof.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Choose:
	//	*FetchRequest_Option1
	//	*FetchRequest_Option2
	//	*FetchRequest_Option3
	Choose isFetchRequest_Choose `protobuf_oneof:"choose"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{0}
}

func (m *FetchRequest) GetChoose() isFetchRequest_Choose {
	if m != nil {
		return m.Choose
	}
	return nil
}

func (x *FetchRequest) GetOption1() string {
	if x, ok := x.GetChoose().(*FetchRequest_Option1); ok {
		return x.Option1
	}
	return ""
}

func (x *FetchRequest) GetOption2() string {
	if x, ok := x.GetChoose().(*FetchRequest_Option2); ok {
		return x.Option2
	}
	return ""
}

func (x *FetchRequest) GetOption3() string {
	if x, ok := x.GetChoose().(*FetchRequest_Option3); ok {
		return x.Option3
	}
	return ""
}

type isFetchRequest_Choose interface {
	isFetchRequest_Choose()
}

type FetchRequest_Option1 struct {
	Option1 string `protobuf:"bytes,1,opt,name=option1,proto3,oneof"`
}

type FetchRequest_Option2 struct {
	Option2 string `protobuf:"bytes,2,opt,name=option2,proto3,oneof"`
}

type FetchRequest_Option3 struct {
	Option3 string `protobuf:"bytes,3,opt,name=option3,proto3,oneof"`
}

func (*FetchRequest_Option1) isFetchRequest_Choose() {}

func (*FetchRequest_Option2) isFetchRequest_Choose() {}

func (*FetchRequest_Option3) isFetchRequest_Choose() {}

type FetchNestedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L0 *FetchNestedRequest_Outer `protobuf:"bytes,1,opt,name=l0,proto3" json:"l0,omitempty"`
}

func (x *FetchNestedRequest) Reset() {
	*x = FetchNestedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchNestedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchNestedRequest) ProtoMessage() {}

func (x *FetchNestedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchNestedRequest.ProtoReflect.Descriptor instead.
func (*FetchNestedRequest) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1}
}

func (x *FetchNestedRequest) GetL0() *FetchNestedRequest_Outer {
	if x != nil {
		return x.L0
	}
	return nil
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{2}
}

func (x *FetchResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type FetchNestedRequest_Outer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L1 *FetchNestedRequest_Outer_Middle `protobuf:"bytes,1,opt,name=l1,proto3" json:"l1,omitempty"`
}

func (x *FetchNestedRequest_Outer) Reset() {
	*x = FetchNestedRequest_Outer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchNestedRequest_Outer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchNestedRequest_Outer) ProtoMessage() {}

func (x *FetchNestedRequest_Outer) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchNestedRequest_Outer.ProtoReflect.Descriptor instead.
func (*FetchNestedRequest_Outer) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FetchNestedRequest_Outer) GetL1() *FetchNestedRequest_Outer_Middle {
	if x != nil {
		return x.L1
	}
	return nil
}

type FetchNestedRequest_Outer_Middle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L2 *FetchNestedRequest_Outer_Middle_Inner `protobuf:"bytes,1,opt,name=l2,proto3" json:"l2,omitempty"`
}

func (x *FetchNestedRequest_Outer_Middle) Reset() {
	*x = FetchNestedRequest_Outer_Middle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchNestedRequest_Outer_Middle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchNestedRequest_Outer_Middle) ProtoMessage() {}

func (x *FetchNestedRequest_Outer_Middle) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchNestedRequest_Outer_Middle.ProtoReflect.Descriptor instead.
func (*FetchNestedRequest_Outer_Middle) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *FetchNestedRequest_Outer_Middle) GetL2() *FetchNestedRequest_Outer_Middle_Inner {
	if x != nil {
		return x.L2
	}
	return nil
}

type FetchNestedRequest_Outer_Middle_Inner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Choose:
	//	*FetchNestedRequest_Outer_Middle_Inner_Option1
	//	*FetchNestedRequest_Outer_Middle_Inner_Option2
	//	*FetchNestedRequest_Outer_Middle_Inner_Option3
	Choose isFetchNestedRequest_Outer_Middle_Inner_Choose `protobuf_oneof:"choose"`
}

func (x *FetchNestedRequest_Outer_Middle_Inner) Reset() {
	*x = FetchNestedRequest_Outer_Middle_Inner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oneof_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchNestedRequest_Outer_Middle_Inner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchNestedRequest_Outer_Middle_Inner) ProtoMessage() {}

func (x *FetchNestedRequest_Outer_Middle_Inner) ProtoReflect() protoreflect.Message {
	mi := &file_oneof_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchNestedRequest_Outer_Middle_Inner.ProtoReflect.Descriptor instead.
func (*FetchNestedRequest_Outer_Middle_Inner) Descriptor() ([]byte, []int) {
	return file_oneof_proto_rawDescGZIP(), []int{1, 0, 0, 0}
}

func (m *FetchNestedRequest_Outer_Middle_Inner) GetChoose() isFetchNestedRequest_Outer_Middle_Inner_Choose {
	if m != nil {
		return m.Choose
	}
	return nil
}

func (x *FetchNestedRequest_Outer_Middle_Inner) GetOption1() string {
	if x, ok := x.GetChoose().(*FetchNestedRequest_Outer_Middle_Inner_Option1); ok {
		return x.Option1
	}
	return ""
}

func (x *FetchNestedRequest_Outer_Middle_Inner) GetOption2() string {
	if x, ok := x.GetChoose().(*FetchNestedRequest_Outer_Middle_Inner_Option2); ok {
		return x.Option2
	}
	return ""
}

func (x *FetchNestedRequest_Outer_Middle_Inner) GetOption3() string {
	if x, ok := x.GetChoose().(*FetchNestedRequest_Outer_Middle_Inner_Option3); ok {
		return x.Option3
	}
	return ""
}

type isFetchNestedRequest_Outer_Middle_Inner_Choose interface {
	isFetchNestedRequest_Outer_Middle_Inner_Choose()
}

type FetchNestedRequest_Outer_Middle_Inner_Option1 struct {
	Option1 string `protobuf:"bytes,1,opt,name=option1,proto3,oneof"`
}

type FetchNestedRequest_Outer_Middle_Inner_Option2 struct {
	Option2 string `protobuf:"bytes,2,opt,name=option2,proto3,oneof"`
}

type FetchNestedRequest_Outer_Middle_Inner_Option3 struct {
	Option3 string `protobuf:"bytes,3,opt,name=option3,proto3,oneof"`
}

func (*FetchNestedRequest_Outer_Middle_Inner_Option1) isFetchNestedRequest_Outer_Middle_Inner_Choose() {
}

func (*FetchNestedRequest_Outer_Middle_Inner_Option2) isFetchNestedRequest_Outer_Middle_Inner_Choose() {
}

func (*FetchNestedRequest_Outer_Middle_Inner_Option3) isFetchNestedRequest_Outer_Middle_Inner_Choose() {
}

var File_oneof_proto protoreflect.FileDescriptor

var file_oneof_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x6c, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x31, 0x12, 0x1a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x12, 0x1a,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x68,
	0x6f, 0x6f, 0x73, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x02, 0x6c,
	0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x02, 0x6c, 0x30, 0x1a, 0xf3,
	0x01, 0x0a, 0x05, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x02, 0x6c, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x52, 0x02,
	0x6c, 0x31, 0x1a, 0xaf, 0x01, 0x0a, 0x06, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x12, 0x3e, 0x0a,
	0x02, 0x6c, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x02, 0x6c, 0x32, 0x1a, 0x65, 0x0a,
	0x05, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x31, 0x12, 0x1a, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x12, 0x1a,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x68,
	0x6f, 0x6f, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x83, 0x01, 0x0a, 0x05,
	0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x36, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x15,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_oneof_proto_rawDescOnce sync.Once
	file_oneof_proto_rawDescData = file_oneof_proto_rawDesc
)

func file_oneof_proto_rawDescGZIP() []byte {
	file_oneof_proto_rawDescOnce.Do(func() {
		file_oneof_proto_rawDescData = protoimpl.X.CompressGZIP(file_oneof_proto_rawDescData)
	})
	return file_oneof_proto_rawDescData
}

var file_oneof_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_oneof_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),                          // 0: example.FetchRequest
	(*FetchNestedRequest)(nil),                    // 1: example.FetchNestedRequest
	(*FetchResponse)(nil),                         // 2: example.FetchResponse
	(*FetchNestedRequest_Outer)(nil),              // 3: example.FetchNestedRequest.Outer
	(*FetchNestedRequest_Outer_Middle)(nil),       // 4: example.FetchNestedRequest.Outer.Middle
	(*FetchNestedRequest_Outer_Middle_Inner)(nil), // 5: example.FetchNestedRequest.Outer.Middle.Inner
}
var file_oneof_proto_depIdxs = []int32{
	3, // 0: example.FetchNestedRequest.l0:type_name -> example.FetchNestedRequest.Outer
	4, // 1: example.FetchNestedRequest.Outer.l1:type_name -> example.FetchNestedRequest.Outer.Middle
	5, // 2: example.FetchNestedRequest.Outer.Middle.l2:type_name -> example.FetchNestedRequest.Outer.Middle.Inner
	0, // 3: example.Oneof.Fetch:input_type -> example.FetchRequest
	1, // 4: example.Oneof.FetchNested:input_type -> example.FetchNestedRequest
	2, // 5: example.Oneof.Fetch:output_type -> example.FetchResponse
	2, // 6: example.Oneof.FetchNested:output_type -> example.FetchResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_oneof_proto_init() }
func file_oneof_proto_init() {
	if File_oneof_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oneof_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_oneof_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchNestedRequest); i {
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
		file_oneof_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_oneof_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchNestedRequest_Outer); i {
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
		file_oneof_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchNestedRequest_Outer_Middle); i {
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
		file_oneof_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchNestedRequest_Outer_Middle_Inner); i {
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
	file_oneof_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FetchRequest_Option1)(nil),
		(*FetchRequest_Option2)(nil),
		(*FetchRequest_Option3)(nil),
	}
	file_oneof_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*FetchNestedRequest_Outer_Middle_Inner_Option1)(nil),
		(*FetchNestedRequest_Outer_Middle_Inner_Option2)(nil),
		(*FetchNestedRequest_Outer_Middle_Inner_Option3)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oneof_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oneof_proto_goTypes,
		DependencyIndexes: file_oneof_proto_depIdxs,
		MessageInfos:      file_oneof_proto_msgTypes,
	}.Build()
	File_oneof_proto = out.File
	file_oneof_proto_rawDesc = nil
	file_oneof_proto_goTypes = nil
	file_oneof_proto_depIdxs = nil
}
