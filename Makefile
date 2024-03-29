# Copyright 2022 VMware, Inc.
# SPDX-License-Identifier: BSD-2-Clause

SHELL = /bin/bash

default: build

# #### GO Binary Management ####
.PHONY: deps-go-binary deps-counterfeiter deps-ginkgo deps-golangci-lint

GO_VERSION := $(shell go version)
GO_VERSION_REQUIRED = go1.18
GO_VERSION_MATCHED := $(shell go version | grep $(GO_VERSION_REQUIRED))

deps-go-binary:
ifndef GO_VERSION
	$(error Go not installed)
endif
ifndef GO_VERSION_MATCHED
	$(error Required Go version is $(GO_VERSION_REQUIRED), but was $(GO_VERSION))
endif
	@:

HAS_COUNTERFEITER := $(shell command -v counterfeiter;)
HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
PLATFORM := $(shell uname -s)

deps-counterfeiter: deps-go-binary
ifndef HAS_COUNTERFEITER
	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
endif

deps-ginkgo: deps-go-binary
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/ginkgo@latest
endif

deps-golangci-lint: deps-go-binary
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
endif
endif

# #### CLEAN ####
.PHONY: clean

clean: deps-go-binary
	rm -rf build/*
	go clean --modcache


# #### DEPS ####
.PHONY: deps deps-counterfeiter deps-ginkgo deps-modules

deps-modules: deps-go-binary
	go mod download

deps: deps-modules deps-counterfeiter deps-ginkgo


# #### BUILD ####
.PHONY: build

SRC = $(shell find . -name "*.go" | grep -v "_test\." )
VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/vmware-labs/marketplace-cli/v2/cmd.version=$(VERSION)"

build/check: $(SRC)
	go build -o build/check -ldflags "-X vmware-samples/concourse-resource-for-marketplace/m/v2/cmd.cmdToExecute=check" ./main.go

build/in: $(SRC)
	go build -o build/in -ldflags "-X vmware-samples/concourse-resource-for-marketplace/m/v2/cmd.cmdToExecute=in" ./main.go

build/out: $(SRC)
	go build -o build/out -ldflags "-X vmware-samples/concourse-resource-for-marketplace/m/v2/cmd.cmdToExecute=out" ./main.go

build: deps build/check build/in build/out

build-image: Dockerfile
	docker build . --tag harbor-repo.vmware.com/tanzu_isv_engineering/mkpcli_concourse_resource:$(VERSION)

# #### TESTS ####
.PHONY: lint test test-features test-units

test-units: deps
	ginkgo -r .

test: deps lint test-units

lint: deps-golangci-lint
	golangci-lint run

# #### DEVOPS ####
.PHONY: set-pipeline
set-pipeline: ci/pipeline.yaml
	fly -t tie set-pipeline --config ci/pipeline.yaml --pipeline marketplace-concourse-resource
