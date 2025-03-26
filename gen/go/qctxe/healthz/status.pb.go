// health/status.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: qctxe/healthz/status.proto

package healthz

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	mi := &file_qctxe_healthz_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qctxe_healthz_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_qctxe_healthz_status_proto_rawDescGZIP(), []int{0}
}

type ServiceHealth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Up            bool                   `protobuf:"varint,1,opt,name=up,proto3" json:"up,omitempty"`
	LastChecked   int64                  `protobuf:"varint,2,opt,name=last_checked,json=lastChecked,proto3" json:"last_checked,omitempty"`
	LastError     string                 `protobuf:"bytes,3,opt,name=last_error,json=lastError,proto3" json:"last_error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceHealth) Reset() {
	*x = ServiceHealth{}
	mi := &file_qctxe_healthz_status_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceHealth) ProtoMessage() {}

func (x *ServiceHealth) ProtoReflect() protoreflect.Message {
	mi := &file_qctxe_healthz_status_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceHealth.ProtoReflect.Descriptor instead.
func (*ServiceHealth) Descriptor() ([]byte, []int) {
	return file_qctxe_healthz_status_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceHealth) GetUp() bool {
	if x != nil {
		return x.Up
	}
	return false
}

func (x *ServiceHealth) GetLastChecked() int64 {
	if x != nil {
		return x.LastChecked
	}
	return 0
}

func (x *ServiceHealth) GetLastError() string {
	if x != nil {
		return x.LastError
	}
	return ""
}

type ClientsHealth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Polygon       *ServiceHealth         `protobuf:"bytes,1,opt,name=polygon,proto3" json:"polygon,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClientsHealth) Reset() {
	*x = ClientsHealth{}
	mi := &file_qctxe_healthz_status_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientsHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientsHealth) ProtoMessage() {}

func (x *ClientsHealth) ProtoReflect() protoreflect.Message {
	mi := &file_qctxe_healthz_status_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientsHealth.ProtoReflect.Descriptor instead.
func (*ClientsHealth) Descriptor() ([]byte, []int) {
	return file_qctxe_healthz_status_proto_rawDescGZIP(), []int{2}
}

func (x *ClientsHealth) GetPolygon() *ServiceHealth {
	if x != nil {
		return x.Polygon
	}
	return nil
}

type GetStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsDegraded    bool                   `protobuf:"varint,1,opt,name=is_degraded,json=isDegraded,proto3" json:"is_degraded,omitempty"`
	IsFatal       bool                   `protobuf:"varint,2,opt,name=is_fatal,json=isFatal,proto3" json:"is_fatal,omitempty"`
	Database      *ServiceHealth         `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *ServiceHealth         `protobuf:"bytes,4,opt,name=redis,proto3" json:"redis,omitempty"`
	Clients       *ClientsHealth         `protobuf:"bytes,5,opt,name=clients,proto3" json:"clients,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	mi := &file_qctxe_healthz_status_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qctxe_healthz_status_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_qctxe_healthz_status_proto_rawDescGZIP(), []int{3}
}

func (x *GetStatusResponse) GetIsDegraded() bool {
	if x != nil {
		return x.IsDegraded
	}
	return false
}

func (x *GetStatusResponse) GetIsFatal() bool {
	if x != nil {
		return x.IsFatal
	}
	return false
}

func (x *GetStatusResponse) GetDatabase() *ServiceHealth {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *GetStatusResponse) GetRedis() *ServiceHealth {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *GetStatusResponse) GetClients() *ClientsHealth {
	if x != nil {
		return x.Clients
	}
	return nil
}

var File_qctxe_healthz_status_proto protoreflect.FileDescriptor

var file_qctxe_healthz_status_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x71, 0x63,
	0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x22, 0x12, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x61, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x0e, 0x0a, 0x02, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x75, 0x70,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x47, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x22, 0xf5, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x61, 0x74, 0x61, 0x6c, 0x12, 0x38, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x63, 0x74, 0x78, 0x65, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71,
	0x63, 0x74, 0x78, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_qctxe_healthz_status_proto_rawDescOnce sync.Once
	file_qctxe_healthz_status_proto_rawDescData []byte
)

func file_qctxe_healthz_status_proto_rawDescGZIP() []byte {
	file_qctxe_healthz_status_proto_rawDescOnce.Do(func() {
		file_qctxe_healthz_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qctxe_healthz_status_proto_rawDesc), len(file_qctxe_healthz_status_proto_rawDesc)))
	})
	return file_qctxe_healthz_status_proto_rawDescData
}

var file_qctxe_healthz_status_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_qctxe_healthz_status_proto_goTypes = []any{
	(*GetStatusRequest)(nil),  // 0: qctxe.healthz.GetStatusRequest
	(*ServiceHealth)(nil),     // 1: qctxe.healthz.ServiceHealth
	(*ClientsHealth)(nil),     // 2: qctxe.healthz.ClientsHealth
	(*GetStatusResponse)(nil), // 3: qctxe.healthz.GetStatusResponse
}
var file_qctxe_healthz_status_proto_depIdxs = []int32{
	1, // 0: qctxe.healthz.ClientsHealth.polygon:type_name -> qctxe.healthz.ServiceHealth
	1, // 1: qctxe.healthz.GetStatusResponse.database:type_name -> qctxe.healthz.ServiceHealth
	1, // 2: qctxe.healthz.GetStatusResponse.redis:type_name -> qctxe.healthz.ServiceHealth
	2, // 3: qctxe.healthz.GetStatusResponse.clients:type_name -> qctxe.healthz.ClientsHealth
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_qctxe_healthz_status_proto_init() }
func file_qctxe_healthz_status_proto_init() {
	if File_qctxe_healthz_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qctxe_healthz_status_proto_rawDesc), len(file_qctxe_healthz_status_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qctxe_healthz_status_proto_goTypes,
		DependencyIndexes: file_qctxe_healthz_status_proto_depIdxs,
		MessageInfos:      file_qctxe_healthz_status_proto_msgTypes,
	}.Build()
	File_qctxe_healthz_status_proto = out.File
	file_qctxe_healthz_status_proto_goTypes = nil
	file_qctxe_healthz_status_proto_depIdxs = nil
}
