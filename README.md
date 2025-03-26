# QWOP Proto

## Setup
### Creating/Modifying new protos
To create a new .proto, follow the scope/service/version format. For versioning, any changes in tag numbers/ordering requires a version bump for the service. v1 is considered unstable right now and subject to rapidly change under development. v2 is considered stable and unlikely to change.


### Building
1.  #### Prerequisites
  - Make
  - [Buf CLI](https://buf.build/docs/cli/installation/)
  - Node
  - Go
2. #### Install TypeScript dependencies
```bash
pnpm install
```
3. 

## Client Generations
Client generations are available under `./gen` for both Go and TypeScript. 
### Go
Go client generations are available at `./gen/go`.

To use them,
```bash
TODO
```

### TypeScript
TypeScript client generations are available at `./gen/ts` or from npm.

1. Install from npm
```bash
pnpm add @ikoz/qwop-proto
```

## Development

### grpcurl

This package doesn't contain an API server. Once you have one setup, you can use these commands to interact/test each method if the server does not have a reflection api (if it does, you don't need to use the .protoset file).

1. List all services
```bash
grpcurl -protoset out.protoset list
```

2. Generate the compiled protoset
```bash
# gen outputs to out.protoset
make gen
# OR manually create it
find ./src -name "*.proto" | xargs protoc --proto_path=./src --descriptor_set_out=out.protoset --include_imports
```

3. Call a service/method
```bash
# Example: call the GetPing method on the PingService
grpcurl -plaintext -protoset out.protoset localhost:8080 qctxe.ping.v1.PingService/GetPing
```