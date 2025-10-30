.PHONY: lint tools gogen genbackend genadmin genqctxe fmt security test

tools:
	rm -rf ./bin
	mkdir ./bin
	# pnpm install
	GOBIN=$(shell pwd)/bin go install \
		github.com/bufbuild/buf/cmd/buf \
		github.com/golangci/golangci-lint/cmd/golangci-lint \
		github.com/google/yamlfmt/cmd/yamlfmt \
		github.com/grpc-ecosystem/grpc-health-probe \
		golang.org/x/vuln/cmd/govulncheck \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		connectrpc.com/connect/cmd/protoc-gen-connect-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

gogen:
	go generate ./...

fmt:
	buf format -w

genadmin:
	rm -rf gen/admin
	rm -rf dist/admin
	cd ./src/admin && buf generate
	cp ./src/admin/package.json ./gen/admin/ts && cp ./src/admin/tsconfig.json ./gen/admin/ts
	cd ./gen/admin/ts && pnpm install --ignore-workspace
	cd ./gen/admin/ts && pnpm run build
	cp ./src/admin/package.json ./dist/admin
	pnpm dlx tsx ./scripts/build.ts --input ./dist/admin --output ./dist/admin
	cd ./dist/admin && pnpm run pack

genqctxe:
	cd ./src/qctxe && buf generate

genbackend:
	rm -rf gen/backend
	rm -rf dist/backend
	# bump the version
	cd ./src/backend && pnpm version patch
	cd ./src/backend && buf generate
	cp ./src/backend/package.json ./gen/backend/ts && cp ./src/backend/tsconfig.json ./gen/backend/ts
	cd ./gen/backend/ts && pnpm install --ignore-workspace
	cd ./gen/backend/ts && pnpm run build
	cp ./src/backend/package.json ./dist/backend
	pnpm dlx tsx ./scripts/build.ts --input ./dist/backend --output ./dist/backend
	cd ./dist/backend && pnpm run pack

# Lint
lint:
	buf lint

# Security

security:
	govulncheck ./...

# Test
test:
	go test -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic $$(go list ./... | grep -v /cmd/)