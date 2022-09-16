// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.13.0
// source: api/openapi-spec/sealet/configuration.proto

package sealet

import (
	_ "github.com/golang/protobuf/ptypes/duration"
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

// config file yaml
type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Debug *Configuration_Debug `protobuf:"bytes,20,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_sealet_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_sealet_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_sealet_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetDebug() *Configuration_Debug {
	if x != nil {
		return x.Debug
	}
	return nil
}

type Configuration_Debug struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dump *Configuration_Debug_Dump `protobuf:"bytes,1,opt,name=dump,proto3" json:"dump,omitempty"`
}

func (x *Configuration_Debug) Reset() {
	*x = Configuration_Debug{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_sealet_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration_Debug) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration_Debug) ProtoMessage() {}

func (x *Configuration_Debug) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_sealet_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration_Debug.ProtoReflect.Descriptor instead.
func (*Configuration_Debug) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_sealet_configuration_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Configuration_Debug) GetDump() *Configuration_Debug_Dump {
	if x != nil {
		return x.Dump
	}
	return nil
}

type Configuration_Debug_Dump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DumpDir string `protobuf:"bytes,2,opt,name=dump_dir,json=dumpDir,proto3" json:"dump_dir,omitempty"`
}

func (x *Configuration_Debug_Dump) Reset() {
	*x = Configuration_Debug_Dump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_openapi_spec_sealet_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration_Debug_Dump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration_Debug_Dump) ProtoMessage() {}

func (x *Configuration_Debug_Dump) ProtoReflect() protoreflect.Message {
	mi := &file_api_openapi_spec_sealet_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration_Debug_Dump.ProtoReflect.Descriptor instead.
func (*Configuration_Debug_Dump) Descriptor() ([]byte, []int) {
	return file_api_openapi_spec_sealet_configuration_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Configuration_Debug_Dump) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Configuration_Debug_Dump) GetDumpDir() string {
	if x != nil {
		return x.DumpDir
	}
	return ""
}

var File_api_openapi_spec_sealet_configuration_proto protoreflect.FileDescriptor

var file_api_openapi_spec_sealet_configuration_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70,
	0x65, 0x63, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73,
	0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x6c,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x1a, 0x85, 0x01, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x75,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x2e, 0x44, 0x75, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x75, 0x6d, 0x70, 0x1a, 0x3b, 0x0a, 0x04, 0x44,
	0x75, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x75, 0x6d, 0x70, 0x44, 0x69, 0x72, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x79, 0x64, 0x78, 0x68, 0x2f, 0x73, 0x65,
	0x61, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x73,
	0x65, 0x61, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x61, 0x6c, 0x65, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_openapi_spec_sealet_configuration_proto_rawDescOnce sync.Once
	file_api_openapi_spec_sealet_configuration_proto_rawDescData = file_api_openapi_spec_sealet_configuration_proto_rawDesc
)

func file_api_openapi_spec_sealet_configuration_proto_rawDescGZIP() []byte {
	file_api_openapi_spec_sealet_configuration_proto_rawDescOnce.Do(func() {
		file_api_openapi_spec_sealet_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_openapi_spec_sealet_configuration_proto_rawDescData)
	})
	return file_api_openapi_spec_sealet_configuration_proto_rawDescData
}

var file_api_openapi_spec_sealet_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_openapi_spec_sealet_configuration_proto_goTypes = []interface{}{
	(*Configuration)(nil),            // 0: sea.api.sealet.v1.Configuration
	(*Configuration_Debug)(nil),      // 1: sea.api.sealet.v1.Configuration.Debug
	(*Configuration_Debug_Dump)(nil), // 2: sea.api.sealet.v1.Configuration.Debug.Dump
}
var file_api_openapi_spec_sealet_configuration_proto_depIdxs = []int32{
	1, // 0: sea.api.sealet.v1.Configuration.debug:type_name -> sea.api.sealet.v1.Configuration.Debug
	2, // 1: sea.api.sealet.v1.Configuration.Debug.dump:type_name -> sea.api.sealet.v1.Configuration.Debug.Dump
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_openapi_spec_sealet_configuration_proto_init() }
func file_api_openapi_spec_sealet_configuration_proto_init() {
	if File_api_openapi_spec_sealet_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_openapi_spec_sealet_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_api_openapi_spec_sealet_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration_Debug); i {
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
		file_api_openapi_spec_sealet_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration_Debug_Dump); i {
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
			RawDescriptor: file_api_openapi_spec_sealet_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_openapi_spec_sealet_configuration_proto_goTypes,
		DependencyIndexes: file_api_openapi_spec_sealet_configuration_proto_depIdxs,
		MessageInfos:      file_api_openapi_spec_sealet_configuration_proto_msgTypes,
	}.Build()
	File_api_openapi_spec_sealet_configuration_proto = out.File
	file_api_openapi_spec_sealet_configuration_proto_rawDesc = nil
	file_api_openapi_spec_sealet_configuration_proto_goTypes = nil
	file_api_openapi_spec_sealet_configuration_proto_depIdxs = nil
}
