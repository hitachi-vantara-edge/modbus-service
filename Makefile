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

HV_REGISTRY=lumadaedge-docker-dev-sc.repo.sc.eng.hitachivantara.com
HV_GATEWAY_REPO=repository/pandora
HV_THIRD_PARTY_REPO=repository/pandora/third-parties
MODBUS_LIB_VERSION=3.1.4
IMG_TAG=${MODBUS_LIB_VERSION}
BRANCH=master
TAG_XT=

AMD64_BUILD=${HV_REGISTRY}/${HV_GATEWAY_REPO}/amd64/hiota-modbus-lib:${IMG_TAG}${TAG_XT}
ARM64_BUILD=${HV_REGISTRY}/${HV_GATEWAY_REPO}/arm64/hiota-modbus-lib:${IMG_TAG}${TAG_XT}
AMD64_BUILDER_IMAGE=${HV_REGISTRY}/${HV_THIRD_PARTY_REPO}/amd64/golang:1.10-stretch
ARM64_BUILDER_IMAGE=${HV_REGISTRY}/${HV_THIRD_PARTY_REPO}/arm64/golang:1.10-stretch
AMD64_BASE_OS=${HV_REGISTRY}/${HV_THIRD_PARTY_REPO}/amd64/debian:stretch
ARM64_BASE_OS=${HV_REGISTRY}/${HV_THIRD_PARTY_REPO}/arm64/debian:stretch
MANIFEST=${HV_REGISTRY}/${HV_GATEWAY_REPO}/hiota-modbus-lib:${IMG_TAG}${TAG_XT}

export GIT_PREFIX
all: buildx
build-amd64: 
	docker build --build-arg BUILDER_IMAGE=${AMD64_BUILDER_IMAGE} \
             --build-arg BASE_OS=${AMD64_BASE_OS} \
             --file build/Dockerfile \
             --tag ${AMD64_BUILD} .

build-arm64: 
	docker build --build-arg BUILDER_IMAGE=${ARM64_BUILDER_IMAGE} \
             --build-arg BASE_OS=${ARM64_BASE_OS} \
             --file build/Dockerfile \
             --tag ${ARM64_BUILD} .

setup-multarch-env:
	@docker run --rm --privileged \
		lumadaedge-docker-dev-sc.repo.sc.eng.hitachivantara.com/repository/pandora/third-parties/amd64/qemu-user-static  \
		--reset -p yes >/dev/null 2>&1

# Jenkins arm64 build stage should make this target
build-arm64-by-jenkins: setup-multarch-env build-arm64 build-arm64
	docker tag ${AMD64_BUILD} amd64/hiota-modbus-lib:${IMG_TAG}${TAG_XT}
	docker tag ${ARM64_BUILD} arm64/hiota-modbus-lib:${IMG_TAG}${TAG_XT}

buildx: build-amd64 build-arm64
	docker manifest rm ${MANIFEST} 2>/dev/null || echo
	docker manifest create ${MANIFEST} ${AMD64_BUILD} ${ARM64_BUILD}
	docker manifest push ${MANIFEST}
