// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: backend/auth/v1/auth.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type RegisterAuthType int32

const (
	RegisterAuthType_REGISTER_AUTH_TYPE_UNSPECIFIED RegisterAuthType = 0
	RegisterAuthType_REGISTER_AUTH_TYPE_EMAIL       RegisterAuthType = 1
)

// Enum value maps for RegisterAuthType.
var (
	RegisterAuthType_name = map[int32]string{
		0: "REGISTER_AUTH_TYPE_UNSPECIFIED",
		1: "REGISTER_AUTH_TYPE_EMAIL",
	}
	RegisterAuthType_value = map[string]int32{
		"REGISTER_AUTH_TYPE_UNSPECIFIED": 0,
		"REGISTER_AUTH_TYPE_EMAIL":       1,
	}
)

func (x RegisterAuthType) Enum() *RegisterAuthType {
	p := new(RegisterAuthType)
	*p = x
	return p
}

func (x RegisterAuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterAuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_backend_auth_v1_auth_proto_enumTypes[0].Descriptor()
}

func (RegisterAuthType) Type() protoreflect.EnumType {
	return &file_backend_auth_v1_auth_proto_enumTypes[0]
}

func (x RegisterAuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterAuthType.Descriptor instead.
func (RegisterAuthType) EnumDescriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

type RegisterUserInfoRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Email           string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username        string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	VerificationKey string                 `protobuf:"bytes,3,opt,name=verification_key,json=verificationKey,proto3" json:"verification_key,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RegisterUserInfoRequest) Reset() {
	*x = RegisterUserInfoRequest{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserInfoRequest) ProtoMessage() {}

func (x *RegisterUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserInfoRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserInfoRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterUserInfoRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterUserInfoRequest) GetVerificationKey() string {
	if x != nil {
		return x.VerificationKey
	}
	return ""
}

type RegisterUserInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserInfoResponse) Reset() {
	*x = RegisterUserInfoResponse{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserInfoResponse) ProtoMessage() {}

func (x *RegisterUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserInfoResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

type VerifyEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyEmailRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VerifyEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type VerifyEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailResponse) Reset() {
	*x = VerifyEmailResponse{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailResponse) ProtoMessage() {}

func (x *VerifyEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailResponse.ProtoReflect.Descriptor instead.
func (*VerifyEmailResponse) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

type CompleteRegistrationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          RegisterAuthType       `protobuf:"varint,1,opt,name=type,proto3,enum=backend.auth.v1.RegisterAuthType" json:"type,omitempty"`
	Password      *string                `protobuf:"bytes,2,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Code          string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteRegistrationRequest) Reset() {
	*x = CompleteRegistrationRequest{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRegistrationRequest) ProtoMessage() {}

func (x *CompleteRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRegistrationRequest.ProtoReflect.Descriptor instead.
func (*CompleteRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *CompleteRegistrationRequest) GetType() RegisterAuthType {
	if x != nil {
		return x.Type
	}
	return RegisterAuthType_REGISTER_AUTH_TYPE_UNSPECIFIED
}

func (x *CompleteRegistrationRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *CompleteRegistrationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CompleteRegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CompleteRegistrationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteRegistrationResponse) Reset() {
	*x = CompleteRegistrationResponse{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRegistrationResponse) ProtoMessage() {}

func (x *CompleteRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRegistrationResponse.ProtoReflect.Descriptor instead.
func (*CompleteRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

type CheckUsernameAvailableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckUsernameAvailableRequest) Reset() {
	*x = CheckUsernameAvailableRequest{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckUsernameAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUsernameAvailableRequest) ProtoMessage() {}

func (x *CheckUsernameAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUsernameAvailableRequest.ProtoReflect.Descriptor instead.
func (*CheckUsernameAvailableRequest) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *CheckUsernameAvailableRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CheckUsernameAvailableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Available     bool                   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckUsernameAvailableResponse) Reset() {
	*x = CheckUsernameAvailableResponse{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckUsernameAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUsernameAvailableResponse) ProtoMessage() {}

func (x *CheckUsernameAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUsernameAvailableResponse.ProtoReflect.Descriptor instead.
func (*CheckUsernameAvailableResponse) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *CheckUsernameAvailableResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *CheckUsernameAvailableResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type PasswordLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordLoginRequest) Reset() {
	*x = PasswordLoginRequest{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordLoginRequest) ProtoMessage() {}

func (x *PasswordLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordLoginRequest.ProtoReflect.Descriptor instead.
func (*PasswordLoginRequest) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *PasswordLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PasswordLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PasswordLoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordLoginResponse) Reset() {
	*x = PasswordLoginResponse{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordLoginResponse) ProtoMessage() {}

func (x *PasswordLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordLoginResponse.ProtoReflect.Descriptor instead.
func (*PasswordLoginResponse) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{9}
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{10}
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_backend_auth_v1_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_auth_v1_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_backend_auth_v1_auth_proto_rawDescGZIP(), []int{11}
}

var File_backend_auth_v1_auth_proto protoreflect.FileDescriptor

var file_backend_auth_v1_auth_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x17, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e,
	0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x15,
	0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x1d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x1e, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x1e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x14, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x54, 0x0a,
	0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49,
	0x4c, 0x10, 0x01, 0x32, 0xf7, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x14, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7b, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2e, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60,
	0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x25, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_backend_auth_v1_auth_proto_rawDescOnce sync.Once
	file_backend_auth_v1_auth_proto_rawDescData []byte
)

func file_backend_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_backend_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_backend_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_backend_auth_v1_auth_proto_rawDesc), len(file_backend_auth_v1_auth_proto_rawDesc)))
	})
	return file_backend_auth_v1_auth_proto_rawDescData
}

var file_backend_auth_v1_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backend_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_backend_auth_v1_auth_proto_goTypes = []any{
	(RegisterAuthType)(0),                  // 0: backend.auth.v1.RegisterAuthType
	(*RegisterUserInfoRequest)(nil),        // 1: backend.auth.v1.RegisterUserInfoRequest
	(*RegisterUserInfoResponse)(nil),       // 2: backend.auth.v1.RegisterUserInfoResponse
	(*VerifyEmailRequest)(nil),             // 3: backend.auth.v1.VerifyEmailRequest
	(*VerifyEmailResponse)(nil),            // 4: backend.auth.v1.VerifyEmailResponse
	(*CompleteRegistrationRequest)(nil),    // 5: backend.auth.v1.CompleteRegistrationRequest
	(*CompleteRegistrationResponse)(nil),   // 6: backend.auth.v1.CompleteRegistrationResponse
	(*CheckUsernameAvailableRequest)(nil),  // 7: backend.auth.v1.CheckUsernameAvailableRequest
	(*CheckUsernameAvailableResponse)(nil), // 8: backend.auth.v1.CheckUsernameAvailableResponse
	(*PasswordLoginRequest)(nil),           // 9: backend.auth.v1.PasswordLoginRequest
	(*PasswordLoginResponse)(nil),          // 10: backend.auth.v1.PasswordLoginResponse
	(*LogoutRequest)(nil),                  // 11: backend.auth.v1.LogoutRequest
	(*LogoutResponse)(nil),                 // 12: backend.auth.v1.LogoutResponse
}
var file_backend_auth_v1_auth_proto_depIdxs = []int32{
	0,  // 0: backend.auth.v1.CompleteRegistrationRequest.type:type_name -> backend.auth.v1.RegisterAuthType
	1,  // 1: backend.auth.v1.AuthService.RegisterUserInfo:input_type -> backend.auth.v1.RegisterUserInfoRequest
	3,  // 2: backend.auth.v1.AuthService.VerifyEmail:input_type -> backend.auth.v1.VerifyEmailRequest
	5,  // 3: backend.auth.v1.AuthService.CompleteRegistration:input_type -> backend.auth.v1.CompleteRegistrationRequest
	7,  // 4: backend.auth.v1.AuthService.CheckUsernameAvailable:input_type -> backend.auth.v1.CheckUsernameAvailableRequest
	9,  // 5: backend.auth.v1.AuthService.PasswordLogin:input_type -> backend.auth.v1.PasswordLoginRequest
	11, // 6: backend.auth.v1.AuthService.Logout:input_type -> backend.auth.v1.LogoutRequest
	2,  // 7: backend.auth.v1.AuthService.RegisterUserInfo:output_type -> backend.auth.v1.RegisterUserInfoResponse
	4,  // 8: backend.auth.v1.AuthService.VerifyEmail:output_type -> backend.auth.v1.VerifyEmailResponse
	6,  // 9: backend.auth.v1.AuthService.CompleteRegistration:output_type -> backend.auth.v1.CompleteRegistrationResponse
	8,  // 10: backend.auth.v1.AuthService.CheckUsernameAvailable:output_type -> backend.auth.v1.CheckUsernameAvailableResponse
	10, // 11: backend.auth.v1.AuthService.PasswordLogin:output_type -> backend.auth.v1.PasswordLoginResponse
	12, // 12: backend.auth.v1.AuthService.Logout:output_type -> backend.auth.v1.LogoutResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_backend_auth_v1_auth_proto_init() }
func file_backend_auth_v1_auth_proto_init() {
	if File_backend_auth_v1_auth_proto != nil {
		return
	}
	file_backend_auth_v1_auth_proto_msgTypes[4].OneofWrappers = []any{}
	file_backend_auth_v1_auth_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_backend_auth_v1_auth_proto_rawDesc), len(file_backend_auth_v1_auth_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_backend_auth_v1_auth_proto_depIdxs,
		EnumInfos:         file_backend_auth_v1_auth_proto_enumTypes,
		MessageInfos:      file_backend_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_backend_auth_v1_auth_proto = out.File
	file_backend_auth_v1_auth_proto_goTypes = nil
	file_backend_auth_v1_auth_proto_depIdxs = nil
}
