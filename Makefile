.PHONY: fmt gy help run test

.DEFAULT_GOAL := help

GOPATH := $(shell go env GOPATH)
PARSER_GO := parser/parser.go
PARSER_GO_Y := parser/parser.go.y

fmt: ## Format GO code
	@go fmt ./...

gy: ## Execute goyacc
	"$(GOPATH)/bin/goyacc" -o $(PARSER_GO) $(PARSER_GO_Y)

run: ## Execute main.go
	@go run main.go

test: ## Check if everyting going well
	@go test -v ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
