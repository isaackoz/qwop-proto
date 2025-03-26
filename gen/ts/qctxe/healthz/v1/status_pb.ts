// health/status.proto

// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file qctxe/healthz/v1/status.proto (package qctxe.healthz.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file qctxe/healthz/v1/status.proto.
 */
export const file_qctxe_healthz_v1_status: GenFile = /*@__PURE__*/
  fileDesc("Ch1xY3R4ZS9oZWFsdGh6L3YxL3N0YXR1cy5wcm90bxIQcWN0eGUuaGVhbHRoei52MSISChBHZXRTdGF0dXNSZXF1ZXN0IkUKDVNlcnZpY2VIZWFsdGgSCgoCdXAYASABKAgSFAoMbGFzdF9jaGVja2VkGAIgASgDEhIKCmxhc3RfZXJyb3IYAyABKAkiQQoNQ2xpZW50c0hlYWx0aBIwCgdwb2x5Z29uGAEgASgLMh8ucWN0eGUuaGVhbHRoei52MS5TZXJ2aWNlSGVhbHRoIs8BChFHZXRTdGF0dXNSZXNwb25zZRITCgtpc19kZWdyYWRlZBgBIAEoCBIQCghpc19mYXRhbBgCIAEoCBIxCghkYXRhYmFzZRgDIAEoCzIfLnFjdHhlLmhlYWx0aHoudjEuU2VydmljZUhlYWx0aBIuCgVyZWRpcxgEIAEoCzIfLnFjdHhlLmhlYWx0aHoudjEuU2VydmljZUhlYWx0aBIwCgdjbGllbnRzGAUgASgLMh8ucWN0eGUuaGVhbHRoei52MS5DbGllbnRzSGVhbHRoQgZaBC4vdjFiBnByb3RvMw");

/**
 * @generated from message qctxe.healthz.v1.GetStatusRequest
 */
export type GetStatusRequest = Message<"qctxe.healthz.v1.GetStatusRequest"> & {
};

/**
 * Describes the message qctxe.healthz.v1.GetStatusRequest.
 * Use `create(GetStatusRequestSchema)` to create a new message.
 */
export const GetStatusRequestSchema: GenMessage<GetStatusRequest> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_v1_status, 0);

/**
 * @generated from message qctxe.healthz.v1.ServiceHealth
 */
export type ServiceHealth = Message<"qctxe.healthz.v1.ServiceHealth"> & {
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
 * Describes the message qctxe.healthz.v1.ServiceHealth.
 * Use `create(ServiceHealthSchema)` to create a new message.
 */
export const ServiceHealthSchema: GenMessage<ServiceHealth> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_v1_status, 1);

/**
 * @generated from message qctxe.healthz.v1.ClientsHealth
 */
export type ClientsHealth = Message<"qctxe.healthz.v1.ClientsHealth"> & {
  /**
   * @generated from field: qctxe.healthz.v1.ServiceHealth polygon = 1;
   */
  polygon?: ServiceHealth;
};

/**
 * Describes the message qctxe.healthz.v1.ClientsHealth.
 * Use `create(ClientsHealthSchema)` to create a new message.
 */
export const ClientsHealthSchema: GenMessage<ClientsHealth> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_v1_status, 2);

/**
 * @generated from message qctxe.healthz.v1.GetStatusResponse
 */
export type GetStatusResponse = Message<"qctxe.healthz.v1.GetStatusResponse"> & {
  /**
   * @generated from field: bool is_degraded = 1;
   */
  isDegraded: boolean;

  /**
   * @generated from field: bool is_fatal = 2;
   */
  isFatal: boolean;

  /**
   * @generated from field: qctxe.healthz.v1.ServiceHealth database = 3;
   */
  database?: ServiceHealth;

  /**
   * @generated from field: qctxe.healthz.v1.ServiceHealth redis = 4;
   */
  redis?: ServiceHealth;

  /**
   * @generated from field: qctxe.healthz.v1.ClientsHealth clients = 5;
   */
  clients?: ClientsHealth;
};

/**
 * Describes the message qctxe.healthz.v1.GetStatusResponse.
 * Use `create(GetStatusResponseSchema)` to create a new message.
 */
export const GetStatusResponseSchema: GenMessage<GetStatusResponse> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_v1_status, 3);

