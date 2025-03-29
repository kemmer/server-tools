.PHONY: build
name = servertools

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(name)_linux server-tools/cmd/server-tools
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/$(name)_macos server-tools/cmd/server-tools

run:
	@go run server-tools/cmd/server-tools

clean:
	@go clean
	@go clean -modcache
	@rm -rf ./bin

lint:
	@golangci-lint run -v

fix:
	@golangci-lint run -v --fix

test:
	@go test ./...

update:
	@go get -u ./...
	@go mod tidy