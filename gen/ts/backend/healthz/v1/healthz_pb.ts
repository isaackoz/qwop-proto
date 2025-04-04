// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file backend/healthz/v1/healthz.proto (package backend.healthz.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { GetPingRequestSchema, GetPingResponseSchema } from "./ping_pb";
import { file_backend_healthz_v1_ping } from "./ping_pb";
import type { GetStatusRequestSchema, GetStatusResponseSchema } from "./status_pb";
import { file_backend_healthz_v1_status } from "./status_pb";

/**
 * Describes the file backend/healthz/v1/healthz.proto.
 */
export const file_backend_healthz_v1_healthz: GenFile = /*@__PURE__*/
  fileDesc("CiBiYWNrZW5kL2hlYWx0aHovdjEvaGVhbHRoei5wcm90bxISYmFja2VuZC5oZWFsdGh6LnYxMsIBCg5IZWFsdGh6U2VydmljZRJUCgdHZXRQaW5nEiIuYmFja2VuZC5oZWFsdGh6LnYxLkdldFBpbmdSZXF1ZXN0GiMuYmFja2VuZC5oZWFsdGh6LnYxLkdldFBpbmdSZXNwb25zZSIAEloKCUdldFN0YXR1cxIkLmJhY2tlbmQuaGVhbHRoei52MS5HZXRTdGF0dXNSZXF1ZXN0GiUuYmFja2VuZC5oZWFsdGh6LnYxLkdldFN0YXR1c1Jlc3BvbnNlIgBCBloELi92MWIGcHJvdG8z", [file_backend_healthz_v1_ping, file_backend_healthz_v1_status]);

/**
 * @generated from service backend.healthz.v1.HealthzService
 */
export const HealthzService: GenService<{
  /**
   * @generated from rpc backend.healthz.v1.HealthzService.GetPing
   */
  getPing: {
    methodKind: "unary";
    input: typeof GetPingRequestSchema;
    output: typeof GetPingResponseSchema;
  },
  /**
   * @generated from rpc backend.healthz.v1.HealthzService.GetStatus
   */
  getStatus: {
    methodKind: "unary";
    input: typeof GetStatusRequestSchema;
    output: typeof GetStatusResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_backend_healthz_v1_healthz, 0);

