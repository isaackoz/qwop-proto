// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: account/v1/account.proto

package v1

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AccountServiceName is the fully-qualified name of the AccountService service.
	AccountServiceName = "account.v1.AccountService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccountServiceUpdatePersonalSettingsProcedure is the fully-qualified name of the AccountService's
	// UpdatePersonalSettings RPC.
	AccountServiceUpdatePersonalSettingsProcedure = "/account.v1.AccountService/UpdatePersonalSettings"
	// AccountServiceGetAccountSettingsProcedure is the fully-qualified name of the AccountService's
	// GetAccountSettings RPC.
	AccountServiceGetAccountSettingsProcedure = "/account.v1.AccountService/GetAccountSettings"
	// AccountServiceGetGeneralSettingsProcedure is the fully-qualified name of the AccountService's
	// GetGeneralSettings RPC.
	AccountServiceGetGeneralSettingsProcedure = "/account.v1.AccountService/GetGeneralSettings"
	// AccountServiceUpdateGeneralSettingsProcedure is the fully-qualified name of the AccountService's
	// UpdateGeneralSettings RPC.
	AccountServiceUpdateGeneralSettingsProcedure = "/account.v1.AccountService/UpdateGeneralSettings"
)

// AccountServiceClient is a client for the account.v1.AccountService service.
type AccountServiceClient interface {
	// Settings related rpcs
	UpdatePersonalSettings(context.Context, *connect.Request[UpdatePersonalSettingsRequest]) (*connect.Response[UpdatePersonalSettingsResponse], error)
	GetAccountSettings(context.Context, *connect.Request[GetAccountSettingsRequest]) (*connect.Response[GetAccountSettingsResponse], error)
	GetGeneralSettings(context.Context, *connect.Request[GetGeneralSettingsRequest]) (*connect.Response[GetGeneralSettingsResponse], error)
	UpdateGeneralSettings(context.Context, *connect.Request[UpdateGeneralSettingsRequest]) (*connect.Response[UpdateGeneralSettingsResponse], error)
}

// NewAccountServiceClient constructs a client for the account.v1.AccountService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccountServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccountServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	accountServiceMethods := File_account_v1_account_proto.Services().ByName("AccountService").Methods()
	return &accountServiceClient{
		updatePersonalSettings: connect.NewClient[UpdatePersonalSettingsRequest, UpdatePersonalSettingsResponse](
			httpClient,
			baseURL+AccountServiceUpdatePersonalSettingsProcedure,
			connect.WithSchema(accountServiceMethods.ByName("UpdatePersonalSettings")),
			connect.WithClientOptions(opts...),
		),
		getAccountSettings: connect.NewClient[GetAccountSettingsRequest, GetAccountSettingsResponse](
			httpClient,
			baseURL+AccountServiceGetAccountSettingsProcedure,
			connect.WithSchema(accountServiceMethods.ByName("GetAccountSettings")),
			connect.WithClientOptions(opts...),
		),
		getGeneralSettings: connect.NewClient[GetGeneralSettingsRequest, GetGeneralSettingsResponse](
			httpClient,
			baseURL+AccountServiceGetGeneralSettingsProcedure,
			connect.WithSchema(accountServiceMethods.ByName("GetGeneralSettings")),
			connect.WithClientOptions(opts...),
		),
		updateGeneralSettings: connect.NewClient[UpdateGeneralSettingsRequest, UpdateGeneralSettingsResponse](
			httpClient,
			baseURL+AccountServiceUpdateGeneralSettingsProcedure,
			connect.WithSchema(accountServiceMethods.ByName("UpdateGeneralSettings")),
			connect.WithClientOptions(opts...),
		),
	}
}

// accountServiceClient implements AccountServiceClient.
type accountServiceClient struct {
	updatePersonalSettings *connect.Client[UpdatePersonalSettingsRequest, UpdatePersonalSettingsResponse]
	getAccountSettings     *connect.Client[GetAccountSettingsRequest, GetAccountSettingsResponse]
	getGeneralSettings     *connect.Client[GetGeneralSettingsRequest, GetGeneralSettingsResponse]
	updateGeneralSettings  *connect.Client[UpdateGeneralSettingsRequest, UpdateGeneralSettingsResponse]
}

// UpdatePersonalSettings calls account.v1.AccountService.UpdatePersonalSettings.
func (c *accountServiceClient) UpdatePersonalSettings(ctx context.Context, req *connect.Request[UpdatePersonalSettingsRequest]) (*connect.Response[UpdatePersonalSettingsResponse], error) {
	return c.updatePersonalSettings.CallUnary(ctx, req)
}

// GetAccountSettings calls account.v1.AccountService.GetAccountSettings.
func (c *accountServiceClient) GetAccountSettings(ctx context.Context, req *connect.Request[GetAccountSettingsRequest]) (*connect.Response[GetAccountSettingsResponse], error) {
	return c.getAccountSettings.CallUnary(ctx, req)
}

// GetGeneralSettings calls account.v1.AccountService.GetGeneralSettings.
func (c *accountServiceClient) GetGeneralSettings(ctx context.Context, req *connect.Request[GetGeneralSettingsRequest]) (*connect.Response[GetGeneralSettingsResponse], error) {
	return c.getGeneralSettings.CallUnary(ctx, req)
}

// UpdateGeneralSettings calls account.v1.AccountService.UpdateGeneralSettings.
func (c *accountServiceClient) UpdateGeneralSettings(ctx context.Context, req *connect.Request[UpdateGeneralSettingsRequest]) (*connect.Response[UpdateGeneralSettingsResponse], error) {
	return c.updateGeneralSettings.CallUnary(ctx, req)
}

// AccountServiceHandler is an implementation of the account.v1.AccountService service.
type AccountServiceHandler interface {
	// Settings related rpcs
	UpdatePersonalSettings(context.Context, *connect.Request[UpdatePersonalSettingsRequest]) (*connect.Response[UpdatePersonalSettingsResponse], error)
	GetAccountSettings(context.Context, *connect.Request[GetAccountSettingsRequest]) (*connect.Response[GetAccountSettingsResponse], error)
	GetGeneralSettings(context.Context, *connect.Request[GetGeneralSettingsRequest]) (*connect.Response[GetGeneralSettingsResponse], error)
	UpdateGeneralSettings(context.Context, *connect.Request[UpdateGeneralSettingsRequest]) (*connect.Response[UpdateGeneralSettingsResponse], error)
}

// NewAccountServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccountServiceHandler(svc AccountServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accountServiceMethods := File_account_v1_account_proto.Services().ByName("AccountService").Methods()
	accountServiceUpdatePersonalSettingsHandler := connect.NewUnaryHandler(
		AccountServiceUpdatePersonalSettingsProcedure,
		svc.UpdatePersonalSettings,
		connect.WithSchema(accountServiceMethods.ByName("UpdatePersonalSettings")),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceGetAccountSettingsHandler := connect.NewUnaryHandler(
		AccountServiceGetAccountSettingsProcedure,
		svc.GetAccountSettings,
		connect.WithSchema(accountServiceMethods.ByName("GetAccountSettings")),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceGetGeneralSettingsHandler := connect.NewUnaryHandler(
		AccountServiceGetGeneralSettingsProcedure,
		svc.GetGeneralSettings,
		connect.WithSchema(accountServiceMethods.ByName("GetGeneralSettings")),
		connect.WithHandlerOptions(opts...),
	)
	accountServiceUpdateGeneralSettingsHandler := connect.NewUnaryHandler(
		AccountServiceUpdateGeneralSettingsProcedure,
		svc.UpdateGeneralSettings,
		connect.WithSchema(accountServiceMethods.ByName("UpdateGeneralSettings")),
		connect.WithHandlerOptions(opts...),
	)
	return "/account.v1.AccountService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccountServiceUpdatePersonalSettingsProcedure:
			accountServiceUpdatePersonalSettingsHandler.ServeHTTP(w, r)
		case AccountServiceGetAccountSettingsProcedure:
			accountServiceGetAccountSettingsHandler.ServeHTTP(w, r)
		case AccountServiceGetGeneralSettingsProcedure:
			accountServiceGetGeneralSettingsHandler.ServeHTTP(w, r)
		case AccountServiceUpdateGeneralSettingsProcedure:
			accountServiceUpdateGeneralSettingsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccountServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccountServiceHandler struct{}

func (UnimplementedAccountServiceHandler) UpdatePersonalSettings(context.Context, *connect.Request[UpdatePersonalSettingsRequest]) (*connect.Response[UpdatePersonalSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.UpdatePersonalSettings is not implemented"))
}

func (UnimplementedAccountServiceHandler) GetAccountSettings(context.Context, *connect.Request[GetAccountSettingsRequest]) (*connect.Response[GetAccountSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.GetAccountSettings is not implemented"))
}

func (UnimplementedAccountServiceHandler) GetGeneralSettings(context.Context, *connect.Request[GetGeneralSettingsRequest]) (*connect.Response[GetGeneralSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.GetGeneralSettings is not implemented"))
}

func (UnimplementedAccountServiceHandler) UpdateGeneralSettings(context.Context, *connect.Request[UpdateGeneralSettingsRequest]) (*connect.Response[UpdateGeneralSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("account.v1.AccountService.UpdateGeneralSettings is not implemented"))
}
