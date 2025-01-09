BIN=aioz-ads

swagger:
	@swag init -g cmd/http/main.go
	@swag fmt

build:
	@go build -o bin/$(BIN) cmd/http/*.go

run: build
	@ENV=debug ./bin/$(BIN)
