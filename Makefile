.PHONY: gen clean lint buildts

# Generate
gen:
	$(MAKE) clean
	buf generate

clean:
	rm -rf gen/
	rm -rf dist/
	rm -rf buildcache/ 

fmt:
	buf format -w

buildts:
	$(MAKE) gen
	pnpm run build
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