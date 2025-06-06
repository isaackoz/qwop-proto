// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chat/v1/chat.proto

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
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "chat.v1.ChatService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChatServiceChatProcedure is the fully-qualified name of the ChatService's Chat RPC.
	ChatServiceChatProcedure = "/chat.v1.ChatService/Chat"
	// ChatServiceMockChatProcedure is the fully-qualified name of the ChatService's MockChat RPC.
	ChatServiceMockChatProcedure = "/chat.v1.ChatService/MockChat"
)

// ChatServiceClient is a client for the chat.v1.ChatService service.
type ChatServiceClient interface {
	Chat(context.Context, *connect.Request[ChatRequest]) (*connect.ServerStreamForClient[ChatResponse], error)
	MockChat(context.Context, *connect.Request[MockChatRequest]) (*connect.ServerStreamForClient[MockChatResponse], error)
}

// NewChatServiceClient constructs a client for the chat.v1.ChatService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	chatServiceMethods := File_chat_v1_chat_proto.Services().ByName("ChatService").Methods()
	return &chatServiceClient{
		chat: connect.NewClient[ChatRequest, ChatResponse](
			httpClient,
			baseURL+ChatServiceChatProcedure,
			connect.WithSchema(chatServiceMethods.ByName("Chat")),
			connect.WithClientOptions(opts...),
		),
		mockChat: connect.NewClient[MockChatRequest, MockChatResponse](
			httpClient,
			baseURL+ChatServiceMockChatProcedure,
			connect.WithSchema(chatServiceMethods.ByName("MockChat")),
			connect.WithClientOptions(opts...),
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	chat     *connect.Client[ChatRequest, ChatResponse]
	mockChat *connect.Client[MockChatRequest, MockChatResponse]
}

// Chat calls chat.v1.ChatService.Chat.
func (c *chatServiceClient) Chat(ctx context.Context, req *connect.Request[ChatRequest]) (*connect.ServerStreamForClient[ChatResponse], error) {
	return c.chat.CallServerStream(ctx, req)
}

// MockChat calls chat.v1.ChatService.MockChat.
func (c *chatServiceClient) MockChat(ctx context.Context, req *connect.Request[MockChatRequest]) (*connect.ServerStreamForClient[MockChatResponse], error) {
	return c.mockChat.CallServerStream(ctx, req)
}

// ChatServiceHandler is an implementation of the chat.v1.ChatService service.
type ChatServiceHandler interface {
	Chat(context.Context, *connect.Request[ChatRequest], *connect.ServerStream[ChatResponse]) error
	MockChat(context.Context, *connect.Request[MockChatRequest], *connect.ServerStream[MockChatResponse]) error
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	chatServiceMethods := File_chat_v1_chat_proto.Services().ByName("ChatService").Methods()
	chatServiceChatHandler := connect.NewServerStreamHandler(
		ChatServiceChatProcedure,
		svc.Chat,
		connect.WithSchema(chatServiceMethods.ByName("Chat")),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceMockChatHandler := connect.NewServerStreamHandler(
		ChatServiceMockChatProcedure,
		svc.MockChat,
		connect.WithSchema(chatServiceMethods.ByName("MockChat")),
		connect.WithHandlerOptions(opts...),
	)
	return "/chat.v1.ChatService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChatServiceChatProcedure:
			chatServiceChatHandler.ServeHTTP(w, r)
		case ChatServiceMockChatProcedure:
			chatServiceMockChatHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) Chat(context.Context, *connect.Request[ChatRequest], *connect.ServerStream[ChatResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.Chat is not implemented"))
}

func (UnimplementedChatServiceHandler) MockChat(context.Context, *connect.Request[MockChatRequest], *connect.ServerStream[MockChatResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.MockChat is not implemented"))
}
