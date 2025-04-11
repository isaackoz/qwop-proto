// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file backend/account/v1/account.proto (package backend.account.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { UpdatePersonalSettingsRequestSchema, UpdatePersonalSettingsResponseSchema } from "./settings_pb";
import { file_backend_account_v1_settings } from "./settings_pb";

/**
 * Describes the file backend/account/v1/account.proto.
 */
export const file_backend_account_v1_account: GenFile = /*@__PURE__*/
  fileDesc("CiBiYWNrZW5kL2FjY291bnQvdjEvYWNjb3VudC5wcm90bxISYmFja2VuZC5hY2NvdW50LnYxMpQBCg5BY2NvdW50U2VydmljZRKBAQoWVXBkYXRlUGVyc29uYWxTZXR0aW5ncxIxLmJhY2tlbmQuYWNjb3VudC52MS5VcGRhdGVQZXJzb25hbFNldHRpbmdzUmVxdWVzdBoyLmJhY2tlbmQuYWNjb3VudC52MS5VcGRhdGVQZXJzb25hbFNldHRpbmdzUmVzcG9uc2UiAEIGWgQuL3YxYgZwcm90bzM", [file_backend_account_v1_settings]);

/**
 * @generated from service backend.account.v1.AccountService
 */
export const AccountService: GenService<{
  /**
   * Settings related rpcs
   *
   * @generated from rpc backend.account.v1.AccountService.UpdatePersonalSettings
   */
  updatePersonalSettings: {
    methodKind: "unary";
    input: typeof UpdatePersonalSettingsRequestSchema;
    output: typeof UpdatePersonalSettingsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_backend_account_v1_account, 0);

