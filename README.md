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