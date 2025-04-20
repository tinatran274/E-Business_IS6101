BIN=aioz-ads

swagger:
	@swag init -g cmd/app/main.go
	@swag fmt

build:
	@go build -o bin/$(BIN) cmd/app/main.go

run: build
	@ENV=debug ./bin/$(BIN)
