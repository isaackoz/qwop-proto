// @generated by protoc-gen-connect-query v2.0.1 with parameter "target=ts"
// @generated from file backend/account/v1/account.proto (package backend.account.v1, syntax proto3)
/* eslint-disable */

import { AccountService } from "./account_pb";

/**
 * Settings related rpcs
 *
 * @generated from rpc backend.account.v1.AccountService.UpdatePersonalSettings
 */
export const updatePersonalSettings = AccountService.method.updatePersonalSettings;

/**
 * User related rpcs
 *
 * @generated from rpc backend.account.v1.AccountService.CheckUsernameAvailable
 */
export const checkUsernameAvailable = AccountService.method.checkUsernameAvailable;
