// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: defaults.proto

package pbdefaults

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

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set default value for field that types includes:
	// string, enum, bool,
	// int32, int64, sint32, sint64, sfixed32, sfixed64,
	// uint32, uint64,fixed32, fixed64.
	Basic *string `protobuf:"bytes,101,opt,name=basic,proto3,oneof" json:"basic,omitempty"`
	// Set default value for field of types repeated.
	Array []string `protobuf:"bytes,102,rep,name=array,proto3" json:"array,omitempty"`
	// Set default value for field of types map.
	Map map[string]string `protobuf:"bytes,103,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defaults_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_defaults_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_defaults_proto_rawDescGZIP(), []int{0}
}

func (x *FieldOptions) GetBasic() string {
	if x != nil && x.Basic != nil {
		return *x.Basic
	}
	return ""
}

func (x *FieldOptions) GetArray() []string {
	if x != nil {
		return x.Array
	}
	return nil
}

func (x *FieldOptions) GetMap() map[string]string {
	if x != nil {
		return x.Map
	}
	return nil
}

type OneOfOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field name that in oneof part.
	Field *string `protobuf:"bytes,1,opt,name=field,proto3,oneof" json:"field,omitempty"`
}

func (x *OneOfOptions) Reset() {
	*x = OneOfOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_defaults_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneOfOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneOfOptions) ProtoMessage() {}

func (x *OneOfOptions) ProtoReflect() protoreflect.Message {
	mi := &file_defaults_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneOfOptions.ProtoReflect.Descriptor instead.
func (*OneOfOptions) Descriptor() ([]byte, []int) {
	return file_defaults_proto_rawDescGZIP(), []int{1}
}

func (x *OneOfOptions) GetField() string {
	if x != nil && x.Field != nil {
		return *x.Field
	}
	return ""
}

var file_defaults_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         64020,
		Name:          "defaults.field",
		Tag:           "bytes,64020,opt,name=field",
		Filename:      "defaults.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*OneOfOptions)(nil),
		Field:         64021,
		Name:          "defaults.oneof",
		Tag:           "bytes,64021,opt,name=oneof",
		Filename:      "defaults.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional defaults.FieldOptions field = 64020;
	E_Field = &file_defaults_proto_extTypes[0]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional defaults.OneOfOptions oneof = 64021;
	E_Oneof = &file_defaults_proto_extTypes[1]
)

var File_defaults_proto protoreflect.FileDescriptor

var file_defaults_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a,
	0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a,
	0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x18, 0x66, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x31,
	0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x67, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61,
	0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x22, 0x33, 0x0a, 0x0c, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4d, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x94, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4d, 0x0a, 0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x95, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x05, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x64, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x79, 0x75, 0x33, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x70, 0x62, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x0a, 0x50,
	0x42, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x50, 0x00, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x62, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_defaults_proto_rawDescOnce sync.Once
	file_defaults_proto_rawDescData = file_defaults_proto_rawDesc
)

func file_defaults_proto_rawDescGZIP() []byte {
	file_defaults_proto_rawDescOnce.Do(func() {
		file_defaults_proto_rawDescData = protoimpl.X.CompressGZIP(file_defaults_proto_rawDescData)
	})
	return file_defaults_proto_rawDescData
}

var file_defaults_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_defaults_proto_goTypes = []interface{}{
	(*FieldOptions)(nil),              // 0: defaults.FieldOptions
	(*OneOfOptions)(nil),              // 1: defaults.OneOfOptions
	nil,                               // 2: defaults.FieldOptions.MapEntry
	(*descriptorpb.FieldOptions)(nil), // 3: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil), // 4: google.protobuf.OneofOptions
}
var file_defaults_proto_depIdxs = []int32{
	2, // 0: defaults.FieldOptions.map:type_name -> defaults.FieldOptions.MapEntry
	3, // 1: defaults.field:extendee -> google.protobuf.FieldOptions
	4, // 2: defaults.oneof:extendee -> google.protobuf.OneofOptions
	0, // 3: defaults.field:type_name -> defaults.FieldOptions
	1, // 4: defaults.oneof:type_name -> defaults.OneOfOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_defaults_proto_init() }
func file_defaults_proto_init() {
	if File_defaults_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_defaults_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
		file_defaults_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneOfOptions); i {
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
	file_defaults_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_defaults_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_defaults_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_defaults_proto_goTypes,
		DependencyIndexes: file_defaults_proto_depIdxs,
		MessageInfos:      file_defaults_proto_msgTypes,
		ExtensionInfos:    file_defaults_proto_extTypes,
	}.Build()
	File_defaults_proto = out.File
	file_defaults_proto_rawDesc = nil
	file_defaults_proto_goTypes = nil
	file_defaults_proto_depIdxs = nil
}
