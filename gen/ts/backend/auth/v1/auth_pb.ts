// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file backend/auth/v1/auth.proto (package backend.auth.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_buf_validate_validate } from "../../../buf/validate/validate_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file backend/auth/v1/auth.proto.
 */
export const file_backend_auth_v1_auth: GenFile = /*@__PURE__*/
  fileDesc("ChpiYWNrZW5kL2F1dGgvdjEvYXV0aC5wcm90bxIPYmFja2VuZC5hdXRoLnYxIlQKF1JlZ2lzdGVyVXNlckluZm9SZXF1ZXN0Eg0KBWVtYWlsGAEgASgJEhAKCHVzZXJuYW1lGAIgASgJEhgKEHZlcmlmaWNhdGlvbl9rZXkYAyABKAkiGgoYUmVnaXN0ZXJVc2VySW5mb1Jlc3BvbnNlIjEKElZlcmlmeUVtYWlsUmVxdWVzdBIMCgRjb2RlGAEgASgJEg0KBWVtYWlsGAIgASgJIhUKE1ZlcmlmeUVtYWlsUmVzcG9uc2UijwEKG0NvbXBsZXRlUmVnaXN0cmF0aW9uUmVxdWVzdBIvCgR0eXBlGAEgASgOMiEuYmFja2VuZC5hdXRoLnYxLlJlZ2lzdGVyQXV0aFR5cGUSFQoIcGFzc3dvcmQYAiABKAlIAIgBARIMCgRjb2RlGAMgASgJEg0KBWVtYWlsGAQgASgJQgsKCV9wYXNzd29yZCIeChxDb21wbGV0ZVJlZ2lzdHJhdGlvblJlc3BvbnNlIjwKHUNoZWNrVXNlcm5hbWVBdmFpbGFibGVSZXF1ZXN0EhsKCHVzZXJuYW1lGAEgASgJQgm6SAZyBBABGB4iVQoeQ2hlY2tVc2VybmFtZUF2YWlsYWJsZVJlc3BvbnNlEhEKCWF2YWlsYWJsZRgBIAEoCBIUCgdtZXNzYWdlGAIgASgJSACIAQFCCgoIX21lc3NhZ2UiNwoUUGFzc3dvcmRMb2dpblJlcXVlc3QSDQoFZW1haWwYASABKAkSEAoIcGFzc3dvcmQYAiABKAkiFwoVUGFzc3dvcmRMb2dpblJlc3BvbnNlIg8KDUxvZ291dFJlcXVlc3QiEAoOTG9nb3V0UmVzcG9uc2UqVAoQUmVnaXN0ZXJBdXRoVHlwZRIiCh5SRUdJU1RFUl9BVVRIX1RZUEVfVU5TUEVDSUZJRUQQABIcChhSRUdJU1RFUl9BVVRIX1RZUEVfRU1BSUwQATL3BAoLQXV0aFNlcnZpY2USaQoQUmVnaXN0ZXJVc2VySW5mbxIoLmJhY2tlbmQuYXV0aC52MS5SZWdpc3RlclVzZXJJbmZvUmVxdWVzdBopLmJhY2tlbmQuYXV0aC52MS5SZWdpc3RlclVzZXJJbmZvUmVzcG9uc2UiABJaCgtWZXJpZnlFbWFpbBIjLmJhY2tlbmQuYXV0aC52MS5WZXJpZnlFbWFpbFJlcXVlc3QaJC5iYWNrZW5kLmF1dGgudjEuVmVyaWZ5RW1haWxSZXNwb25zZSIAEnUKFENvbXBsZXRlUmVnaXN0cmF0aW9uEiwuYmFja2VuZC5hdXRoLnYxLkNvbXBsZXRlUmVnaXN0cmF0aW9uUmVxdWVzdBotLmJhY2tlbmQuYXV0aC52MS5Db21wbGV0ZVJlZ2lzdHJhdGlvblJlc3BvbnNlIgASewoWQ2hlY2tVc2VybmFtZUF2YWlsYWJsZRIuLmJhY2tlbmQuYXV0aC52MS5DaGVja1VzZXJuYW1lQXZhaWxhYmxlUmVxdWVzdBovLmJhY2tlbmQuYXV0aC52MS5DaGVja1VzZXJuYW1lQXZhaWxhYmxlUmVzcG9uc2UiABJgCg1QYXNzd29yZExvZ2luEiUuYmFja2VuZC5hdXRoLnYxLlBhc3N3b3JkTG9naW5SZXF1ZXN0GiYuYmFja2VuZC5hdXRoLnYxLlBhc3N3b3JkTG9naW5SZXNwb25zZSIAEksKBkxvZ291dBIeLmJhY2tlbmQuYXV0aC52MS5Mb2dvdXRSZXF1ZXN0Gh8uYmFja2VuZC5hdXRoLnYxLkxvZ291dFJlc3BvbnNlIgBCBloELi92MWIGcHJvdG8z", [file_buf_validate_validate]);

/**
 * @generated from message backend.auth.v1.RegisterUserInfoRequest
 */
export type RegisterUserInfoRequest = Message<"backend.auth.v1.RegisterUserInfoRequest"> & {
  /**
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * @generated from field: string username = 2;
   */
  username: string;

  /**
   * @generated from field: string verification_key = 3;
   */
  verificationKey: string;
};

/**
 * Describes the message backend.auth.v1.RegisterUserInfoRequest.
 * Use `create(RegisterUserInfoRequestSchema)` to create a new message.
 */
export const RegisterUserInfoRequestSchema: GenMessage<RegisterUserInfoRequest> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 0);

/**
 * @generated from message backend.auth.v1.RegisterUserInfoResponse
 */
export type RegisterUserInfoResponse = Message<"backend.auth.v1.RegisterUserInfoResponse"> & {
};

/**
 * Describes the message backend.auth.v1.RegisterUserInfoResponse.
 * Use `create(RegisterUserInfoResponseSchema)` to create a new message.
 */
export const RegisterUserInfoResponseSchema: GenMessage<RegisterUserInfoResponse> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 1);

/**
 * @generated from message backend.auth.v1.VerifyEmailRequest
 */
export type VerifyEmailRequest = Message<"backend.auth.v1.VerifyEmailRequest"> & {
  /**
   * @generated from field: string code = 1;
   */
  code: string;

  /**
   * @generated from field: string email = 2;
   */
  email: string;
};

/**
 * Describes the message backend.auth.v1.VerifyEmailRequest.
 * Use `create(VerifyEmailRequestSchema)` to create a new message.
 */
export const VerifyEmailRequestSchema: GenMessage<VerifyEmailRequest> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 2);

/**
 * @generated from message backend.auth.v1.VerifyEmailResponse
 */
export type VerifyEmailResponse = Message<"backend.auth.v1.VerifyEmailResponse"> & {
};

/**
 * Describes the message backend.auth.v1.VerifyEmailResponse.
 * Use `create(VerifyEmailResponseSchema)` to create a new message.
 */
export const VerifyEmailResponseSchema: GenMessage<VerifyEmailResponse> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 3);

/**
 * @generated from message backend.auth.v1.CompleteRegistrationRequest
 */
export type CompleteRegistrationRequest = Message<"backend.auth.v1.CompleteRegistrationRequest"> & {
  /**
   * @generated from field: backend.auth.v1.RegisterAuthType type = 1;
   */
  type: RegisterAuthType;

  /**
   * @generated from field: optional string password = 2;
   */
  password?: string;

  /**
   * @generated from field: string code = 3;
   */
  code: string;

  /**
   * @generated from field: string email = 4;
   */
  email: string;
};

/**
 * Describes the message backend.auth.v1.CompleteRegistrationRequest.
 * Use `create(CompleteRegistrationRequestSchema)` to create a new message.
 */
export const CompleteRegistrationRequestSchema: GenMessage<CompleteRegistrationRequest> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 4);

/**
 * @generated from message backend.auth.v1.CompleteRegistrationResponse
 */
export type CompleteRegistrationResponse = Message<"backend.auth.v1.CompleteRegistrationResponse"> & {
};

/**
 * Describes the message backend.auth.v1.CompleteRegistrationResponse.
 * Use `create(CompleteRegistrationResponseSchema)` to create a new message.
 */
export const CompleteRegistrationResponseSchema: GenMessage<CompleteRegistrationResponse> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 5);

/**
 * @generated from message backend.auth.v1.CheckUsernameAvailableRequest
 */
export type CheckUsernameAvailableRequest = Message<"backend.auth.v1.CheckUsernameAvailableRequest"> & {
  /**
   * @generated from field: string username = 1;
   */
  username: string;
};

/**
 * Describes the message backend.auth.v1.CheckUsernameAvailableRequest.
 * Use `create(CheckUsernameAvailableRequestSchema)` to create a new message.
 */
export const CheckUsernameAvailableRequestSchema: GenMessage<CheckUsernameAvailableRequest> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 6);

/**
 * @generated from message backend.auth.v1.CheckUsernameAvailableResponse
 */
export type CheckUsernameAvailableResponse = Message<"backend.auth.v1.CheckUsernameAvailableResponse"> & {
  /**
   * @generated from field: bool available = 1;
   */
  available: boolean;

  /**
   * @generated from field: optional string message = 2;
   */
  message?: string;
};

/**
 * Describes the message backend.auth.v1.CheckUsernameAvailableResponse.
 * Use `create(CheckUsernameAvailableResponseSchema)` to create a new message.
 */
export const CheckUsernameAvailableResponseSchema: GenMessage<CheckUsernameAvailableResponse> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 7);

/**
 * @generated from message backend.auth.v1.PasswordLoginRequest
 */
export type PasswordLoginRequest = Message<"backend.auth.v1.PasswordLoginRequest"> & {
  /**
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;
};

/**
 * Describes the message backend.auth.v1.PasswordLoginRequest.
 * Use `create(PasswordLoginRequestSchema)` to create a new message.
 */
export const PasswordLoginRequestSchema: GenMessage<PasswordLoginRequest> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 8);

/**
 * @generated from message backend.auth.v1.PasswordLoginResponse
 */
export type PasswordLoginResponse = Message<"backend.auth.v1.PasswordLoginResponse"> & {
};

/**
 * Describes the message backend.auth.v1.PasswordLoginResponse.
 * Use `create(PasswordLoginResponseSchema)` to create a new message.
 */
export const PasswordLoginResponseSchema: GenMessage<PasswordLoginResponse> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 9);

/**
 * @generated from message backend.auth.v1.LogoutRequest
 */
export type LogoutRequest = Message<"backend.auth.v1.LogoutRequest"> & {
};

/**
 * Describes the message backend.auth.v1.LogoutRequest.
 * Use `create(LogoutRequestSchema)` to create a new message.
 */
export const LogoutRequestSchema: GenMessage<LogoutRequest> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 10);

/**
 * @generated from message backend.auth.v1.LogoutResponse
 */
export type LogoutResponse = Message<"backend.auth.v1.LogoutResponse"> & {
};

/**
 * Describes the message backend.auth.v1.LogoutResponse.
 * Use `create(LogoutResponseSchema)` to create a new message.
 */
export const LogoutResponseSchema: GenMessage<LogoutResponse> = /*@__PURE__*/
  messageDesc(file_backend_auth_v1_auth, 11);

/**
 * @generated from enum backend.auth.v1.RegisterAuthType
 */
export enum RegisterAuthType {
  /**
   * @generated from enum value: REGISTER_AUTH_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: REGISTER_AUTH_TYPE_EMAIL = 1;
   */
  EMAIL = 1,
}

/**
 * Describes the enum backend.auth.v1.RegisterAuthType.
 */
export const RegisterAuthTypeSchema: GenEnum<RegisterAuthType> = /*@__PURE__*/
  enumDesc(file_backend_auth_v1_auth, 0);

/**
 * @generated from service backend.auth.v1.AuthService
 */
export const AuthService: GenService<{
  /**
   * @generated from rpc backend.auth.v1.AuthService.RegisterUserInfo
   */
  registerUserInfo: {
    methodKind: "unary";
    input: typeof RegisterUserInfoRequestSchema;
    output: typeof RegisterUserInfoResponseSchema;
  },
  /**
   * @generated from rpc backend.auth.v1.AuthService.VerifyEmail
   */
  verifyEmail: {
    methodKind: "unary";
    input: typeof VerifyEmailRequestSchema;
    output: typeof VerifyEmailResponseSchema;
  },
  /**
   * @generated from rpc backend.auth.v1.AuthService.CompleteRegistration
   */
  completeRegistration: {
    methodKind: "unary";
    input: typeof CompleteRegistrationRequestSchema;
    output: typeof CompleteRegistrationResponseSchema;
  },
  /**
   * @generated from rpc backend.auth.v1.AuthService.CheckUsernameAvailable
   */
  checkUsernameAvailable: {
    methodKind: "unary";
    input: typeof CheckUsernameAvailableRequestSchema;
    output: typeof CheckUsernameAvailableResponseSchema;
  },
  /**
   * @generated from rpc backend.auth.v1.AuthService.PasswordLogin
   */
  passwordLogin: {
    methodKind: "unary";
    input: typeof PasswordLoginRequestSchema;
    output: typeof PasswordLoginResponseSchema;
  },
  /**
   * @generated from rpc backend.auth.v1.AuthService.Logout
   */
  logout: {
    methodKind: "unary";
    input: typeof LogoutRequestSchema;
    output: typeof LogoutResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_backend_auth_v1_auth, 0);

