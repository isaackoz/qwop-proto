// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: backend/healthz/healthz.proto

package healthz

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_backend_healthz_healthz_proto protoreflect.FileDescriptor

var file_backend_healthz_healthz_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x7a, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x1a, 0x1a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x7a, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb6, 0x01, 0x0a, 0x0e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_backend_healthz_healthz_proto_goTypes = []any{
	(*GetPingRequest)(nil),    // 0: backend.healthz.GetPingRequest
	(*GetStatusRequest)(nil),  // 1: backend.healthz.GetStatusRequest
	(*GetPingResponse)(nil),   // 2: backend.healthz.GetPingResponse
	(*GetStatusResponse)(nil), // 3: backend.healthz.GetStatusResponse
}
var file_backend_healthz_healthz_proto_depIdxs = []int32{
	0, // 0: backend.healthz.HealthzService.GetPing:input_type -> backend.healthz.GetPingRequest
	1, // 1: backend.healthz.HealthzService.GetStatus:input_type -> backend.healthz.GetStatusRequest
	2, // 2: backend.healthz.HealthzService.GetPing:output_type -> backend.healthz.GetPingResponse
	3, // 3: backend.healthz.HealthzService.GetStatus:output_type -> backend.healthz.GetStatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_healthz_healthz_proto_init() }
func file_backend_healthz_healthz_proto_init() {
	if File_backend_healthz_healthz_proto != nil {
		return
	}
	file_backend_healthz_ping_proto_init()
	file_backend_healthz_status_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_backend_healthz_healthz_proto_rawDesc), len(file_backend_healthz_healthz_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_healthz_healthz_proto_goTypes,
		DependencyIndexes: file_backend_healthz_healthz_proto_depIdxs,
	}.Build()
	File_backend_healthz_healthz_proto = out.File
	file_backend_healthz_healthz_proto_goTypes = nil
	file_backend_healthz_healthz_proto_depIdxs = nil
}
