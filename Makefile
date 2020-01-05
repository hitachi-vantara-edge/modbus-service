SERVER_OUT := "cmd/modbus/main"
PKG := "github.com/hitachi-vantara-edge/modbus-service"
SERVER_PKG_BUILD := "${PKG}/cmd/modbus"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
CERT_PATH := "~/certs"

.PHONY: all build

all: build

dep: ## Get the dependencies
	@glide install -v

build: dep ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(TEST_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docker:
	@docker build -f build/Dockerfile -t hitachivantaraedge/modbus-service .

clean-glide: ## Remove vendor and glide lock
	@rm -Rf vendor glide.lock

glide: clean-glide ## Get the dependencies
	@glide install -v
	@glide update -v