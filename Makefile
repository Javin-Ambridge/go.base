# VARIABLES
PACKAGE="github.com/Javin-Ambridge/go.base"
BINARY_NAME="go.base"

default: usage

clean: ## Trash binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@rm -rf $(GOPATH)/bin/$(BINARY_NAME) 2> /dev/null
	@echo "Clean OK"

test: ## Run all tests
	@echo "--> testing..."
	@go test -v $(PACKAGE)/...

coverage: ## Run all tests and get coverage
	@echo "--> testing and creating coverage HTML..."
	@go test ./... -coverprofile=c.out && go tool cover -html=c.out
	@echo "Coverage Built and Opened in Default Web Browser"

install: clean ## Compile sources and build binary
	@echo "--> installing..."
	@go install $(PACKAGE) || (echo "Compilation error" && exit 1)
	@echo "Install OK"

run: install ## Run your application
	@echo "--> running application..."
	@$(GOPATH)/bin/$(BINARY_NAME)

usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

mocks: ## Generates all the mocks
	@echo "--> Deleting Mocks and Recreating..."
	@rm -rf .gen/mocks/
	@mkdir .gen/mocks/
	@$(GOPATH)/bin/mockgen -destination=.gen/mocks/http/http_response_writer_mock.go -package=mocks net/http ResponseWriter
	@$(GOPATH)/bin/mockgen -destination=.gen/mocks/fx/lifecycle_mock.go -package=mocks go.uber.org/fx Lifecycle
	@$(GOPATH)/bin/mockgen -destination=.gen/mocks/go.base/handler_mocks.go -package=mocks github.com/Javin-Ambridge/go.base/handler Handler


install-golang:
	@echo "--> installing base golang packages..."
	@go get github.com/stretchr/testify/assert
	@go get github.com/golang/mock/gomock
	@go get github.com/shopspring/decimal
	@go get github.com/jinzhu/gorm
	@go get github.com/kubernetes/utils/pointer
	@go get github.com/lib/pq
	@go get github.com/pkg/errors
	@go get go.uber.org/config
	@go get go.uber.org/fx
	@go get go.uber.org/zap
	@go get github.com/mkideal/cli

