// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file backend/healthz/v1/status.proto (package backend.healthz.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file backend/healthz/v1/status.proto.
 */
export const file_backend_healthz_v1_status: GenFile = /*@__PURE__*/
  fileDesc("Ch9iYWNrZW5kL2hlYWx0aHovdjEvc3RhdHVzLnByb3RvEhJiYWNrZW5kLmhlYWx0aHoudjEiEgoQR2V0U3RhdHVzUmVxdWVzdCJFCg1TZXJ2aWNlSGVhbHRoEgoKAnVwGAEgASgIEhQKDGxhc3RfY2hlY2tlZBgCIAEoAxISCgpsYXN0X2Vycm9yGAMgASgJIkMKDUNsaWVudHNIZWFsdGgSMgoHcG9seWdvbhgBIAEoCzIhLmJhY2tlbmQuaGVhbHRoei52MS5TZXJ2aWNlSGVhbHRoItUBChFHZXRTdGF0dXNSZXNwb25zZRITCgtpc19kZWdyYWRlZBgBIAEoCBIQCghpc19mYXRhbBgCIAEoCBIzCghkYXRhYmFzZRgDIAEoCzIhLmJhY2tlbmQuaGVhbHRoei52MS5TZXJ2aWNlSGVhbHRoEjAKBXJlZGlzGAQgASgLMiEuYmFja2VuZC5oZWFsdGh6LnYxLlNlcnZpY2VIZWFsdGgSMgoHY2xpZW50cxgFIAEoCzIhLmJhY2tlbmQuaGVhbHRoei52MS5DbGllbnRzSGVhbHRoQgZaBC4vdjFiBnByb3RvMw");

/**
 * @generated from message backend.healthz.v1.GetStatusRequest
 */
export type GetStatusRequest = Message<"backend.healthz.v1.GetStatusRequest"> & {
};

/**
 * Describes the message backend.healthz.v1.GetStatusRequest.
 * Use `create(GetStatusRequestSchema)` to create a new message.
 */
export const GetStatusRequestSchema: GenMessage<GetStatusRequest> = /*@__PURE__*/
  messageDesc(file_backend_healthz_v1_status, 0);

/**
 * @generated from message backend.healthz.v1.ServiceHealth
 */
export type ServiceHealth = Message<"backend.healthz.v1.ServiceHealth"> & {
  /**
   * @generated from field: bool up = 1;
   */
  up: boolean;

  /**
   * @generated from field: int64 last_checked = 2;
   */
  lastChecked: bigint;

  /**
   * @generated from field: string last_error = 3;
   */
  lastError: string;
};

/**
 * Describes the message backend.healthz.v1.ServiceHealth.
 * Use `create(ServiceHealthSchema)` to create a new message.
 */
export const ServiceHealthSchema: GenMessage<ServiceHealth> = /*@__PURE__*/
  messageDesc(file_backend_healthz_v1_status, 1);

/**
 * @generated from message backend.healthz.v1.ClientsHealth
 */
export type ClientsHealth = Message<"backend.healthz.v1.ClientsHealth"> & {
  /**
   * @generated from field: backend.healthz.v1.ServiceHealth polygon = 1;
   */
  polygon?: ServiceHealth;
};

/**
 * Describes the message backend.healthz.v1.ClientsHealth.
 * Use `create(ClientsHealthSchema)` to create a new message.
 */
export const ClientsHealthSchema: GenMessage<ClientsHealth> = /*@__PURE__*/
  messageDesc(file_backend_healthz_v1_status, 2);

/**
 * @generated from message backend.healthz.v1.GetStatusResponse
 */
export type GetStatusResponse = Message<"backend.healthz.v1.GetStatusResponse"> & {
  /**
   * @generated from field: bool is_degraded = 1;
   */
  isDegraded: boolean;

  /**
   * @generated from field: bool is_fatal = 2;
   */
  isFatal: boolean;

  /**
   * @generated from field: backend.healthz.v1.ServiceHealth database = 3;
   */
  database?: ServiceHealth;

  /**
   * @generated from field: backend.healthz.v1.ServiceHealth redis = 4;
   */
  redis?: ServiceHealth;

  /**
   * @generated from field: backend.healthz.v1.ClientsHealth clients = 5;
   */
  clients?: ClientsHealth;
};

/**
 * Describes the message backend.healthz.v1.GetStatusResponse.
 * Use `create(GetStatusResponseSchema)` to create a new message.
 */
export const GetStatusResponseSchema: GenMessage<GetStatusResponse> = /*@__PURE__*/
  messageDesc(file_backend_healthz_v1_status, 3);

