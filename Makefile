run:
	go run main.go
lint:
	golangci-lint run
format:
	gofmt -w .