.PHONY: dcb dcr fmt gy help run test

.DEFAULT_GOAL := help

GOPATH := $(shell go env GOPATH)
PARSER_GO := parser/parser.go
PARSER_GO_Y := parser/parser.go.y

dcb: ## Build from Dockerfile
	@docker build . -t go_type_inference

dcr: ## Run from Docker Image  (Execute program. Entry file is main.go)
	@echo "Execute program...."
	@docker run --rm --name go_type_inference go_type_inference

fmt: ## Format .go code
	@go fmt ./...

gy: ## Execute goyacc
	"$(GOPATH)/bin/goyacc" -o $(PARSER_GO) $(PARSER_GO_Y)

run: ## Execute main.go
	@go run main.go

test: ## Executes test functions in test files
	@go test -v ./...

help: ## Show options and brief explanations
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
