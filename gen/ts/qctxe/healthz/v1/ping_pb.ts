// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file qctxe/healthz/v1/ping.proto (package qctxe.healthz.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file qctxe/healthz/v1/ping.proto.
 */
export const file_qctxe_healthz_v1_ping: GenFile = /*@__PURE__*/
  fileDesc("ChtxY3R4ZS9oZWFsdGh6L3YxL3BpbmcucHJvdG8SEHFjdHhlLmhlYWx0aHoudjEiEAoOR2V0UGluZ1JlcXVlc3QibAoPR2V0UGluZ1Jlc3BvbnNlEg8KB21lc3NhZ2UYASABKAkSEQoJdGltZXN0YW1wGAIgASgDEjUKBnN0YXR1cxgDIAEoDjIlLnFjdHhlLmhlYWx0aHoudjEuUGluZ1Jlc3BvbnNlT3B0aW9ucyqAAQoTUGluZ1Jlc3BvbnNlT3B0aW9ucxIlCiFQSU5HX1JFU1BPTlNFX09QVElPTlNfVU5TUEVDSUZJRUQQABIhCh1QSU5HX1JFU1BPTlNFX09QVElPTlNfU1VDQ0VTUxABEh8KG1BJTkdfUkVTUE9OU0VfT1BUSU9OU19FUlJPUhACQgZaBC4vdjFiBnByb3RvMw");

/**
 * @generated from message qctxe.healthz.v1.GetPingRequest
 */
export type GetPingRequest = Message<"qctxe.healthz.v1.GetPingRequest"> & {
};

/**
 * Describes the message qctxe.healthz.v1.GetPingRequest.
 * Use `create(GetPingRequestSchema)` to create a new message.
 */
export const GetPingRequestSchema: GenMessage<GetPingRequest> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_v1_ping, 0);

/**
 * @generated from message qctxe.healthz.v1.GetPingResponse
 */
export type GetPingResponse = Message<"qctxe.healthz.v1.GetPingResponse"> & {
  /**
   * @generated from field: string message = 1;
   */
  message: string;

  /**
   * @generated from field: int64 timestamp = 2;
   */
  timestamp: bigint;

  /**
   * @generated from field: qctxe.healthz.v1.PingResponseOptions status = 3;
   */
  status: PingResponseOptions;
};

/**
 * Describes the message qctxe.healthz.v1.GetPingResponse.
 * Use `create(GetPingResponseSchema)` to create a new message.
 */
export const GetPingResponseSchema: GenMessage<GetPingResponse> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_v1_ping, 1);

/**
 * @generated from enum qctxe.healthz.v1.PingResponseOptions
 */
export enum PingResponseOptions {
  /**
   * @generated from enum value: PING_RESPONSE_OPTIONS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PING_RESPONSE_OPTIONS_SUCCESS = 1;
   */
  SUCCESS = 1,

  /**
   * @generated from enum value: PING_RESPONSE_OPTIONS_ERROR = 2;
   */
  ERROR = 2,
}

/**
 * Describes the enum qctxe.healthz.v1.PingResponseOptions.
 */
export const PingResponseOptionsSchema: GenEnum<PingResponseOptions> = /*@__PURE__*/
  enumDesc(file_qctxe_healthz_v1_ping, 0);

