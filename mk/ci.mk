.PHONY: lint
lint:
	@golangci-lint version
	@golangci-lint run
