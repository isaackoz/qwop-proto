// health/status.proto

// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file qctxe/healthz/status.proto (package qctxe.healthz, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file qctxe/healthz/status.proto.
 */
export const file_qctxe_healthz_status: GenFile = /*@__PURE__*/
  fileDesc("ChpxY3R4ZS9oZWFsdGh6L3N0YXR1cy5wcm90bxINcWN0eGUuaGVhbHRoeiISChBHZXRTdGF0dXNSZXF1ZXN0IkUKDVNlcnZpY2VIZWFsdGgSCgoCdXAYASABKAgSFAoMbGFzdF9jaGVja2VkGAIgASgDEhIKCmxhc3RfZXJyb3IYAyABKAkiPgoNQ2xpZW50c0hlYWx0aBItCgdwb2x5Z29uGAEgASgLMhwucWN0eGUuaGVhbHRoei5TZXJ2aWNlSGVhbHRoIsYBChFHZXRTdGF0dXNSZXNwb25zZRITCgtpc19kZWdyYWRlZBgBIAEoCBIQCghpc19mYXRhbBgCIAEoCBIuCghkYXRhYmFzZRgDIAEoCzIcLnFjdHhlLmhlYWx0aHouU2VydmljZUhlYWx0aBIrCgVyZWRpcxgEIAEoCzIcLnFjdHhlLmhlYWx0aHouU2VydmljZUhlYWx0aBItCgdjbGllbnRzGAUgASgLMhwucWN0eGUuaGVhbHRoei5DbGllbnRzSGVhbHRoQgtaCS4vaGVhbHRoemIGcHJvdG8z");

/**
 * @generated from message qctxe.healthz.GetStatusRequest
 */
export type GetStatusRequest = Message<"qctxe.healthz.GetStatusRequest"> & {
};

/**
 * Describes the message qctxe.healthz.GetStatusRequest.
 * Use `create(GetStatusRequestSchema)` to create a new message.
 */
export const GetStatusRequestSchema: GenMessage<GetStatusRequest> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_status, 0);

/**
 * @generated from message qctxe.healthz.ServiceHealth
 */
export type ServiceHealth = Message<"qctxe.healthz.ServiceHealth"> & {
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
 * Describes the message qctxe.healthz.ServiceHealth.
 * Use `create(ServiceHealthSchema)` to create a new message.
 */
export const ServiceHealthSchema: GenMessage<ServiceHealth> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_status, 1);

/**
 * @generated from message qctxe.healthz.ClientsHealth
 */
export type ClientsHealth = Message<"qctxe.healthz.ClientsHealth"> & {
  /**
   * @generated from field: qctxe.healthz.ServiceHealth polygon = 1;
   */
  polygon?: ServiceHealth;
};

/**
 * Describes the message qctxe.healthz.ClientsHealth.
 * Use `create(ClientsHealthSchema)` to create a new message.
 */
export const ClientsHealthSchema: GenMessage<ClientsHealth> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_status, 2);

/**
 * @generated from message qctxe.healthz.GetStatusResponse
 */
export type GetStatusResponse = Message<"qctxe.healthz.GetStatusResponse"> & {
  /**
   * @generated from field: bool is_degraded = 1;
   */
  isDegraded: boolean;

  /**
   * @generated from field: bool is_fatal = 2;
   */
  isFatal: boolean;

  /**
   * @generated from field: qctxe.healthz.ServiceHealth database = 3;
   */
  database?: ServiceHealth;

  /**
   * @generated from field: qctxe.healthz.ServiceHealth redis = 4;
   */
  redis?: ServiceHealth;

  /**
   * @generated from field: qctxe.healthz.ClientsHealth clients = 5;
   */
  clients?: ClientsHealth;
};

/**
 * Describes the message qctxe.healthz.GetStatusResponse.
 * Use `create(GetStatusResponseSchema)` to create a new message.
 */
export const GetStatusResponseSchema: GenMessage<GetStatusResponse> = /*@__PURE__*/
  messageDesc(file_qctxe_healthz_status, 3);

