PROJECT_NAME	  := gomock-extra-matcher
PKG				  := github.com/oxyno-zeta/$(PROJECT_NAME)

# go option
GO        ?= go
# Uncomment to enable vendor
GO_VENDOR := # -mod=vendor

# Required for globs to work correctly
SHELL=/usr/bin/env bash

#  Version

HAS_GORELEASER := $(shell command -v goreleaser;)
HAS_GIT := $(shell command -v git;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_CURL:=$(shell command -v curl;)

.DEFAULT_GOAL := code/lint

#############
#   Build   #
#############

.PHONY: code/lint
code/lint: setup/dep/install
	golangci-lint run ./...

#############
#   Tests   #
#############

.PHONY: test/unit
test/unit: setup/dep/install
	$(GO) test $(GO_VENDOR) -v -coverpkg=./... -coverprofile=c.out.tmp ./...

.PHONY: test/coverage
test/coverage:
	cat c.out.tmp | grep -v "mock_" > c.out
	$(GO) tool cover -html=c.out -o coverage.html
	$(GO) tool cover -func c.out

#############
#   Setup   #
#############

.PHONY: setup/generate
setup/generate:
	$(GO) $(GO_VENDOR) generate ./...

.PHONY: setup/dep/install
setup/dep/install:
ifndef HAS_GOLANGCI_LINT
	@echo "=> Installing golangci-lint tool"
ifndef HAS_CURL
	$(error You must install curl)
endif
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.43.0
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
	go mod download
	go mod tidy

.PHONY: setup/dep/update
setup/dep/update:
	$(GO) get -u ./...

.PHONY: setup/dep/vendor
setup/dep/vendor:
	$(GO) mod vendor
