// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: qctxe/healthz/v1/health.proto

package v1

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

var File_qctxe_healthz_v1_health_proto protoreflect.FileDescriptor

var file_qctxe_healthz_v1_health_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x71, 0x63, 0x74, 0x78, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb9, 0x01,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x71, 0x63, 0x74,
	0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x71,
	0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22,
	0x2e, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_qctxe_healthz_v1_health_proto_goTypes = []any{
	(*GetPingRequest)(nil),    // 0: qctxe.healthz.v1.GetPingRequest
	(*GetStatusRequest)(nil),  // 1: qctxe.healthz.v1.GetStatusRequest
	(*GetPingResponse)(nil),   // 2: qctxe.healthz.v1.GetPingResponse
	(*GetStatusResponse)(nil), // 3: qctxe.healthz.v1.GetStatusResponse
}
var file_qctxe_healthz_v1_health_proto_depIdxs = []int32{
	0, // 0: qctxe.healthz.v1.HealthService.GetPing:input_type -> qctxe.healthz.v1.GetPingRequest
	1, // 1: qctxe.healthz.v1.HealthService.GetStatus:input_type -> qctxe.healthz.v1.GetStatusRequest
	2, // 2: qctxe.healthz.v1.HealthService.GetPing:output_type -> qctxe.healthz.v1.GetPingResponse
	3, // 3: qctxe.healthz.v1.HealthService.GetStatus:output_type -> qctxe.healthz.v1.GetStatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qctxe_healthz_v1_health_proto_init() }
func file_qctxe_healthz_v1_health_proto_init() {
	if File_qctxe_healthz_v1_health_proto != nil {
		return
	}
	file_qctxe_healthz_v1_ping_proto_init()
	file_qctxe_healthz_v1_status_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qctxe_healthz_v1_health_proto_rawDesc), len(file_qctxe_healthz_v1_health_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qctxe_healthz_v1_health_proto_goTypes,
		DependencyIndexes: file_qctxe_healthz_v1_health_proto_depIdxs,
	}.Build()
	File_qctxe_healthz_v1_health_proto = out.File
	file_qctxe_healthz_v1_health_proto_goTypes = nil
	file_qctxe_healthz_v1_health_proto_depIdxs = nil
}
