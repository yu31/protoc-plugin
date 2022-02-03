// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: gosql.proto

package pbgosql

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Serialize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Format:
	//	*Serialize_Json
	//	*Serialize_Protojson
	//	*Serialize_Proto
	//	*Serialize_Gogoproto
	Format isSerialize_Format `protobuf_oneof:"Format"`
}

func (x *Serialize) Reset() {
	*x = Serialize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Serialize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Serialize) ProtoMessage() {}

func (x *Serialize) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Serialize.ProtoReflect.Descriptor instead.
func (*Serialize) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{0}
}

func (m *Serialize) GetFormat() isSerialize_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (x *Serialize) GetJson() *JSON {
	if x, ok := x.GetFormat().(*Serialize_Json); ok {
		return x.Json
	}
	return nil
}

func (x *Serialize) GetProtojson() *ProtoJSON {
	if x, ok := x.GetFormat().(*Serialize_Protojson); ok {
		return x.Protojson
	}
	return nil
}

func (x *Serialize) GetProto() *Proto {
	if x, ok := x.GetFormat().(*Serialize_Proto); ok {
		return x.Proto
	}
	return nil
}

func (x *Serialize) GetGogoproto() *GoGoProto {
	if x, ok := x.GetFormat().(*Serialize_Gogoproto); ok {
		return x.Gogoproto
	}
	return nil
}

type isSerialize_Format interface {
	isSerialize_Format()
}

type Serialize_Json struct {
	Json *JSON `protobuf:"bytes,1,opt,name=json,proto3,oneof"`
}

type Serialize_Protojson struct {
	Protojson *ProtoJSON `protobuf:"bytes,2,opt,name=protojson,proto3,oneof"`
}

type Serialize_Proto struct {
	Proto *Proto `protobuf:"bytes,3,opt,name=proto,proto3,oneof"`
}

type Serialize_Gogoproto struct {
	Gogoproto *GoGoProto `protobuf:"bytes,4,opt,name=gogoproto,proto3,oneof"`
}

func (*Serialize_Json) isSerialize_Format() {}

func (*Serialize_Protojson) isSerialize_Format() {}

func (*Serialize_Proto) isSerialize_Format() {}

func (*Serialize_Gogoproto) isSerialize_Format() {}

// Use "encoding/json"
type JSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JSON) Reset() {
	*x = JSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSON) ProtoMessage() {}

func (x *JSON) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSON.ProtoReflect.Descriptor instead.
func (*JSON) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{1}
}

// Use "google.golang.org/protobuf/encoding/protojson"
type ProtoJSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Marshal_   *ProtoJSON_MarshalOptions   `protobuf:"bytes,1,opt,name=marshal,proto3" json:"marshal,omitempty"`
	Unmarshal_ *ProtoJSON_UnmarshalOptions `protobuf:"bytes,2,opt,name=unmarshal,proto3" json:"unmarshal,omitempty"`
}

func (x *ProtoJSON) Reset() {
	*x = ProtoJSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoJSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoJSON) ProtoMessage() {}

func (x *ProtoJSON) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoJSON.ProtoReflect.Descriptor instead.
func (*ProtoJSON) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoJSON) GetMarshal_() *ProtoJSON_MarshalOptions {
	if x != nil {
		return x.Marshal_
	}
	return nil
}

func (x *ProtoJSON) GetUnmarshal_() *ProtoJSON_UnmarshalOptions {
	if x != nil {
		return x.Unmarshal_
	}
	return nil
}

// Use google.golang.org/protobuf/proto.
type Proto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Marshal_   *Proto_MarshalOptions   `protobuf:"bytes,1,opt,name=marshal,proto3" json:"marshal,omitempty"`
	Unmarshal_ *Proto_UnmarshalOptions `protobuf:"bytes,2,opt,name=unmarshal,proto3" json:"unmarshal,omitempty"`
}

func (x *Proto) Reset() {
	*x = Proto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto) ProtoMessage() {}

func (x *Proto) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto.ProtoReflect.Descriptor instead.
func (*Proto) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{3}
}

func (x *Proto) GetMarshal_() *Proto_MarshalOptions {
	if x != nil {
		return x.Marshal_
	}
	return nil
}

func (x *Proto) GetUnmarshal_() *Proto_UnmarshalOptions {
	if x != nil {
		return x.Unmarshal_
	}
	return nil
}

// Use github.com/gogo/protobuf/proto.
type GoGoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoGoProto) Reset() {
	*x = GoGoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoGoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoGoProto) ProtoMessage() {}

func (x *GoGoProto) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoGoProto.ProtoReflect.Descriptor instead.
func (*GoGoProto) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{4}
}

// protojson.MarshalOptions
type ProtoJSON_MarshalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Multiline       bool   `protobuf:"varint,1,opt,name=Multiline,proto3" json:"Multiline,omitempty"`
	Indent          string `protobuf:"bytes,2,opt,name=Indent,proto3" json:"Indent,omitempty"`
	AllowPartial    bool   `protobuf:"varint,3,opt,name=AllowPartial,proto3" json:"AllowPartial,omitempty"`
	UseProtoNames   bool   `protobuf:"varint,4,opt,name=UseProtoNames,proto3" json:"UseProtoNames,omitempty"`
	UseEnumNumbers  bool   `protobuf:"varint,5,opt,name=UseEnumNumbers,proto3" json:"UseEnumNumbers,omitempty"`
	EmitUnpopulated bool   `protobuf:"varint,6,opt,name=EmitUnpopulated,proto3" json:"EmitUnpopulated,omitempty"`
}

func (x *ProtoJSON_MarshalOptions) Reset() {
	*x = ProtoJSON_MarshalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoJSON_MarshalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoJSON_MarshalOptions) ProtoMessage() {}

func (x *ProtoJSON_MarshalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoJSON_MarshalOptions.ProtoReflect.Descriptor instead.
func (*ProtoJSON_MarshalOptions) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ProtoJSON_MarshalOptions) GetMultiline() bool {
	if x != nil {
		return x.Multiline
	}
	return false
}

func (x *ProtoJSON_MarshalOptions) GetIndent() string {
	if x != nil {
		return x.Indent
	}
	return ""
}

func (x *ProtoJSON_MarshalOptions) GetAllowPartial() bool {
	if x != nil {
		return x.AllowPartial
	}
	return false
}

func (x *ProtoJSON_MarshalOptions) GetUseProtoNames() bool {
	if x != nil {
		return x.UseProtoNames
	}
	return false
}

func (x *ProtoJSON_MarshalOptions) GetUseEnumNumbers() bool {
	if x != nil {
		return x.UseEnumNumbers
	}
	return false
}

func (x *ProtoJSON_MarshalOptions) GetEmitUnpopulated() bool {
	if x != nil {
		return x.EmitUnpopulated
	}
	return false
}

// protojson.UnmarshalOptions
type ProtoJSON_UnmarshalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowPartial   bool `protobuf:"varint,1,opt,name=AllowPartial,proto3" json:"AllowPartial,omitempty"`
	DiscardUnknown bool `protobuf:"varint,2,opt,name=DiscardUnknown,proto3" json:"DiscardUnknown,omitempty"`
}

func (x *ProtoJSON_UnmarshalOptions) Reset() {
	*x = ProtoJSON_UnmarshalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoJSON_UnmarshalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoJSON_UnmarshalOptions) ProtoMessage() {}

func (x *ProtoJSON_UnmarshalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoJSON_UnmarshalOptions.ProtoReflect.Descriptor instead.
func (*ProtoJSON_UnmarshalOptions) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ProtoJSON_UnmarshalOptions) GetAllowPartial() bool {
	if x != nil {
		return x.AllowPartial
	}
	return false
}

func (x *ProtoJSON_UnmarshalOptions) GetDiscardUnknown() bool {
	if x != nil {
		return x.DiscardUnknown
	}
	return false
}

// proto.MarshalOptions
type Proto_MarshalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowPartial  bool `protobuf:"varint,1,opt,name=AllowPartial,proto3" json:"AllowPartial,omitempty"`
	Deterministic bool `protobuf:"varint,2,opt,name=Deterministic,proto3" json:"Deterministic,omitempty"`
	UseCachedSize bool `protobuf:"varint,3,opt,name=UseCachedSize,proto3" json:"UseCachedSize,omitempty"`
}

func (x *Proto_MarshalOptions) Reset() {
	*x = Proto_MarshalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto_MarshalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto_MarshalOptions) ProtoMessage() {}

func (x *Proto_MarshalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto_MarshalOptions.ProtoReflect.Descriptor instead.
func (*Proto_MarshalOptions) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Proto_MarshalOptions) GetAllowPartial() bool {
	if x != nil {
		return x.AllowPartial
	}
	return false
}

func (x *Proto_MarshalOptions) GetDeterministic() bool {
	if x != nil {
		return x.Deterministic
	}
	return false
}

func (x *Proto_MarshalOptions) GetUseCachedSize() bool {
	if x != nil {
		return x.UseCachedSize
	}
	return false
}

// proto.UnmarshalOptions
type Proto_UnmarshalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Merge          bool `protobuf:"varint,1,opt,name=Merge,proto3" json:"Merge,omitempty"`
	AllowPartial   bool `protobuf:"varint,2,opt,name=AllowPartial,proto3" json:"AllowPartial,omitempty"`
	DiscardUnknown bool `protobuf:"varint,3,opt,name=DiscardUnknown,proto3" json:"DiscardUnknown,omitempty"`
}

func (x *Proto_UnmarshalOptions) Reset() {
	*x = Proto_UnmarshalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gosql_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto_UnmarshalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto_UnmarshalOptions) ProtoMessage() {}

func (x *Proto_UnmarshalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_gosql_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto_UnmarshalOptions.ProtoReflect.Descriptor instead.
func (*Proto_UnmarshalOptions) Descriptor() ([]byte, []int) {
	return file_gosql_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Proto_UnmarshalOptions) GetMerge() bool {
	if x != nil {
		return x.Merge
	}
	return false
}

func (x *Proto_UnmarshalOptions) GetAllowPartial() bool {
	if x != nil {
		return x.AllowPartial
	}
	return false
}

func (x *Proto_UnmarshalOptions) GetDiscardUnknown() bool {
	if x != nil {
		return x.DiscardUnknown
	}
	return false
}

var file_gosql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Serialize)(nil),
		Field:         9001,
		Name:          "gosql.serialize",
		Tag:           "bytes,9001,opt,name=serialize",
		Filename:      "gosql.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional gosql.Serialize serialize = 9001;
	E_Serialize = &file_gosql_proto_extTypes[0]
)

var File_gosql_proto protoreflect.FileDescriptor

var file_gosql_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x6f, 0x73, 0x71, 0x6c, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x48,
	0x00, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x73,
	0x71, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4a, 0x53, 0x4f, 0x4e, 0x48, 0x00, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x73, 0x71, 0x6c,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x30, 0x0a, 0x09, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x47, 0x6f, 0x47, 0x6f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x08, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x06, 0x0a, 0x04, 0x4a,
	0x53, 0x4f, 0x4e, 0x22, 0xcc, 0x03, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x39, 0x0a, 0x07, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x09,
	0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4a, 0x53, 0x4f,
	0x4e, 0x2e, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x1a, 0xe2, 0x01,
	0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x49, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x49, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x55, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x55, 0x73, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x45, 0x6d, 0x69, 0x74,
	0x55, 0x6e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x45, 0x6d, 0x69, 0x74, 0x55, 0x6e, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x5e, 0x0a, 0x10, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69,
	0x73, 0x63, 0x61, 0x72, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x22, 0xf4, 0x02, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x07,
	0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x72, 0x73,
	0x68, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x61, 0x72, 0x73,
	0x68, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c,
	0x1a, 0x80, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x24, 0x0a,
	0x0d, 0x55, 0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x55, 0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x53,
	0x69, 0x7a, 0x65, 0x1a, 0x74, 0x0a, 0x10, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x61,
	0x72, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x22, 0x0b, 0x0a, 0x09, 0x47, 0x6f, 0x47,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x50, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa9, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f,
	0x73, 0x71, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x09, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x42, 0x59, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x79, 0x75, 0x33, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x42, 0x07, 0x50, 0x42, 0x47, 0x6f,
	0x53, 0x51, 0x4c, 0x50, 0x00, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x67, 0x6f,
	0x73, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gosql_proto_rawDescOnce sync.Once
	file_gosql_proto_rawDescData = file_gosql_proto_rawDesc
)

func file_gosql_proto_rawDescGZIP() []byte {
	file_gosql_proto_rawDescOnce.Do(func() {
		file_gosql_proto_rawDescData = protoimpl.X.CompressGZIP(file_gosql_proto_rawDescData)
	})
	return file_gosql_proto_rawDescData
}

var file_gosql_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gosql_proto_goTypes = []interface{}{
	(*Serialize)(nil),                   // 0: gosql.Serialize
	(*JSON)(nil),                        // 1: gosql.JSON
	(*ProtoJSON)(nil),                   // 2: gosql.ProtoJSON
	(*Proto)(nil),                       // 3: gosql.Proto
	(*GoGoProto)(nil),                   // 4: gosql.GoGoProto
	(*ProtoJSON_MarshalOptions)(nil),    // 5: gosql.ProtoJSON.MarshalOptions
	(*ProtoJSON_UnmarshalOptions)(nil),  // 6: gosql.ProtoJSON.UnmarshalOptions
	(*Proto_MarshalOptions)(nil),        // 7: gosql.Proto.MarshalOptions
	(*Proto_UnmarshalOptions)(nil),      // 8: gosql.Proto.UnmarshalOptions
	(*descriptorpb.MessageOptions)(nil), // 9: google.protobuf.MessageOptions
}
var file_gosql_proto_depIdxs = []int32{
	1,  // 0: gosql.Serialize.json:type_name -> gosql.JSON
	2,  // 1: gosql.Serialize.protojson:type_name -> gosql.ProtoJSON
	3,  // 2: gosql.Serialize.proto:type_name -> gosql.Proto
	4,  // 3: gosql.Serialize.gogoproto:type_name -> gosql.GoGoProto
	5,  // 4: gosql.ProtoJSON.marshal:type_name -> gosql.ProtoJSON.MarshalOptions
	6,  // 5: gosql.ProtoJSON.unmarshal:type_name -> gosql.ProtoJSON.UnmarshalOptions
	7,  // 6: gosql.Proto.marshal:type_name -> gosql.Proto.MarshalOptions
	8,  // 7: gosql.Proto.unmarshal:type_name -> gosql.Proto.UnmarshalOptions
	9,  // 8: gosql.serialize:extendee -> google.protobuf.MessageOptions
	0,  // 9: gosql.serialize:type_name -> gosql.Serialize
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	9,  // [9:10] is the sub-list for extension type_name
	8,  // [8:9] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_gosql_proto_init() }
func file_gosql_proto_init() {
	if File_gosql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gosql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Serialize); i {
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
		file_gosql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSON); i {
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
		file_gosql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoJSON); i {
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
		file_gosql_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto); i {
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
		file_gosql_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoGoProto); i {
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
		file_gosql_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoJSON_MarshalOptions); i {
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
		file_gosql_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoJSON_UnmarshalOptions); i {
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
		file_gosql_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto_MarshalOptions); i {
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
		file_gosql_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto_UnmarshalOptions); i {
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
	file_gosql_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Serialize_Json)(nil),
		(*Serialize_Protojson)(nil),
		(*Serialize_Proto)(nil),
		(*Serialize_Gogoproto)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gosql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_gosql_proto_goTypes,
		DependencyIndexes: file_gosql_proto_depIdxs,
		MessageInfos:      file_gosql_proto_msgTypes,
		ExtensionInfos:    file_gosql_proto_extTypes,
	}.Build()
	File_gosql_proto = out.File
	file_gosql_proto_rawDesc = nil
	file_gosql_proto_goTypes = nil
	file_gosql_proto_depIdxs = nil
}
