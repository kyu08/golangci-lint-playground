run:
	@go run ./...

lint:
	GOTOOLCHAIN=go1.24.0 go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2 run  ./... --config ./.golangci.yml

.PHONY: run lint
