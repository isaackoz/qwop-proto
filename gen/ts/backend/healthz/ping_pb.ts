// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file backend/healthz/ping.proto (package backend.healthz, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file backend/healthz/ping.proto.
 */
export const file_backend_healthz_ping: GenFile = /*@__PURE__*/
  fileDesc("ChpiYWNrZW5kL2hlYWx0aHovcGluZy5wcm90bxIPYmFja2VuZC5oZWFsdGh6IhAKDkdldFBpbmdSZXF1ZXN0ImsKD0dldFBpbmdSZXNwb25zZRIPCgdtZXNzYWdlGAEgASgJEhEKCXRpbWVzdGFtcBgCIAEoAxI0CgZzdGF0dXMYAyABKA4yJC5iYWNrZW5kLmhlYWx0aHouUGluZ1Jlc3BvbnNlT3B0aW9ucypoChNQaW5nUmVzcG9uc2VPcHRpb25zEh0KGVBJTkdfUkVTUE9OU0VfVU5TUEVDSUZJRUQQABIZChVQSU5HX1JFU1BPTlNFX1NVQ0NFU1MQARIXChNQSU5HX1JFU1BPTlNFX0VSUk9SEAJCC1oJLi9oZWFsdGh6YgZwcm90bzM");

/**
 * @generated from message backend.healthz.GetPingRequest
 */
export type GetPingRequest = Message<"backend.healthz.GetPingRequest"> & {
};

/**
 * Describes the message backend.healthz.GetPingRequest.
 * Use `create(GetPingRequestSchema)` to create a new message.
 */
export const GetPingRequestSchema: GenMessage<GetPingRequest> = /*@__PURE__*/
  messageDesc(file_backend_healthz_ping, 0);

/**
 * @generated from message backend.healthz.GetPingResponse
 */
export type GetPingResponse = Message<"backend.healthz.GetPingResponse"> & {
  /**
   * @generated from field: string message = 1;
   */
  message: string;

  /**
   * @generated from field: int64 timestamp = 2;
   */
  timestamp: bigint;

  /**
   * @generated from field: backend.healthz.PingResponseOptions status = 3;
   */
  status: PingResponseOptions;
};

/**
 * Describes the message backend.healthz.GetPingResponse.
 * Use `create(GetPingResponseSchema)` to create a new message.
 */
export const GetPingResponseSchema: GenMessage<GetPingResponse> = /*@__PURE__*/
  messageDesc(file_backend_healthz_ping, 1);

/**
 * @generated from enum backend.healthz.PingResponseOptions
 */
export enum PingResponseOptions {
  /**
   * @generated from enum value: PING_RESPONSE_UNSPECIFIED = 0;
   */
  PING_RESPONSE_UNSPECIFIED = 0,

  /**
   * @generated from enum value: PING_RESPONSE_SUCCESS = 1;
   */
  PING_RESPONSE_SUCCESS = 1,

  /**
   * @generated from enum value: PING_RESPONSE_ERROR = 2;
   */
  PING_RESPONSE_ERROR = 2,
}

/**
 * Describes the enum backend.healthz.PingResponseOptions.
 */
export const PingResponseOptionsSchema: GenEnum<PingResponseOptions> = /*@__PURE__*/
  enumDesc(file_backend_healthz_ping, 0);

