// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: common/v1/options.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_common_v1_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         69420,
		Name:          "common.v1.public_route",
		Tag:           "varint,69420,opt,name=public_route",
		Filename:      "common/v1/options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional bool public_route = 69420;
	E_PublicRoute = &file_common_v1_options_proto_extTypes[0]
)

var File_common_v1_options_proto protoreflect.FileDescriptor

const file_common_v1_options_proto_rawDesc = "" +
	"\n" +
	"\x17common/v1/options.proto\x12\tcommon.v1\x1a google/protobuf/descriptor.proto:C\n" +
	"\fpublic_route\x12\x1e.google.protobuf.MethodOptions\x18\xac\x9e\x04 \x01(\bR\vpublicRouteB9Z7github.com/isaackoz/qwop-proto/gen/backend/go/common/v1b\x06proto3"

var file_common_v1_options_proto_goTypes = []any{
	(*descriptorpb.MethodOptions)(nil), // 0: google.protobuf.MethodOptions
}
var file_common_v1_options_proto_depIdxs = []int32{
	0, // 0: common.v1.public_route:extendee -> google.protobuf.MethodOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_options_proto_init() }
func file_common_v1_options_proto_init() {
	if File_common_v1_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_v1_options_proto_rawDesc), len(file_common_v1_options_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_options_proto_goTypes,
		DependencyIndexes: file_common_v1_options_proto_depIdxs,
		ExtensionInfos:    file_common_v1_options_proto_extTypes,
	}.Build()
	File_common_v1_options_proto = out.File
	file_common_v1_options_proto_goTypes = nil
	file_common_v1_options_proto_depIdxs = nil
}
