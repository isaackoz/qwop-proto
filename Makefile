.PHONY: gen clean lint buildts tools gogen

# Generate
gen:
	$(MAKE) clean
	buf generate
	find ./src -name "*.proto" | xargs protoc --proto_path=./src --descriptor_set_out=out.protoset --include_imports

clean:
	rm -rf gen/
	rm -rf dist/
	rm -rf buildcache/

tools:
	go install -C internal/tools \
		github.com/bufbuild/buf/cmd/buf \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/google/yamlfmt/cmd/yamlfmt \
		github.com/grpc-ecosystem/grpc-health-probe \
		golang.org/x/vuln/cmd/govulncheck \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		connectrpc.com/connect/cmd/protoc-gen-connect-go

gogen:
	go generate ./...

fmt:
	buf format -w

buildts:
	$(MAKE) gen
	pnpm run build
	cp README.md dist/README.md
	pnpm run pack

# Lint
lint:
	buf lint

# Security

security:
	govulncheck ./...

# Test
test:
	go test -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic $$(go list ./... | grep -v /cmd/)