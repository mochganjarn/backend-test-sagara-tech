run:
	@go run main.go
install: 
	@go mod tidy
lint:
	@golangci-lint run