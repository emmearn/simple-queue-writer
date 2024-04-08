# Golang commands

vet:
	@go vet -tags=integration ./...

format:
	@go fmt ./...

tidy:
	@GOPRIVATE=bitbucket.org/m_arnone/gogeco go mod tidy

download:
	@go mod download

vendor:
	@go mod vendor

mocks:
	@go generate ./...

test:
	@go test -cover -v ./...

integration-test:
	@go test -cover -v -tags=integration ./...

build:
	@go build -o bin/main cmd/*.go

unbuild:
	@rm -rf deployments/.serverless
	@rm -rf bin
	@rm -rf node_modules
	@rm -rf vendor
	@rm -rf package-lock.json
	@make tidy
