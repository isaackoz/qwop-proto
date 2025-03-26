// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file qctxe/healthz/health.proto (package qctxe.healthz, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { GetPingRequestSchema, GetPingResponseSchema } from "./ping_pb";
import { file_qctxe_healthz_ping } from "./ping_pb";
import type { GetStatusRequestSchema, GetStatusResponseSchema } from "./status_pb";
import { file_qctxe_healthz_status } from "./status_pb";

/**
 * Describes the file qctxe/healthz/health.proto.
 */
export const file_qctxe_healthz_health: GenFile = /*@__PURE__*/
  fileDesc("ChpxY3R4ZS9oZWFsdGh6L2hlYWx0aC5wcm90bxINcWN0eGUuaGVhbHRoejKtAQoNSGVhbHRoU2VydmljZRJKCgdHZXRQaW5nEh0ucWN0eGUuaGVhbHRoei5HZXRQaW5nUmVxdWVzdBoeLnFjdHhlLmhlYWx0aHouR2V0UGluZ1Jlc3BvbnNlIgASUAoJR2V0U3RhdHVzEh8ucWN0eGUuaGVhbHRoei5HZXRTdGF0dXNSZXF1ZXN0GiAucWN0eGUuaGVhbHRoei5HZXRTdGF0dXNSZXNwb25zZSIAQgtaCS4vaGVhbHRoemIGcHJvdG8z", [file_qctxe_healthz_ping, file_qctxe_healthz_status]);

/**
 * @generated from service qctxe.healthz.HealthService
 */
export const HealthService: GenService<{
  /**
   * @generated from rpc qctxe.healthz.HealthService.GetPing
   */
  getPing: {
    methodKind: "unary";
    input: typeof GetPingRequestSchema;
    output: typeof GetPingResponseSchema;
  },
  /**
   * @generated from rpc qctxe.healthz.HealthService.GetStatus
   */
  getStatus: {
    methodKind: "unary";
    input: typeof GetStatusRequestSchema;
    output: typeof GetStatusResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_qctxe_healthz_health, 0);

