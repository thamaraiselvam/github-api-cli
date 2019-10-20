.PHONY: all
all: help

VERSION?=1.0
BUILD?=$(shell git describe --tags --always)
GOLINT:=$(shell command -v golint 2> /dev/null)
GOPATH?=$(shell go env GOPATH)

APP_EXECUTABLE="./out/"

RICHGO=$(shell command -v richgo 2> /dev/null)
GO111MODULE=on
GO_ARGS=-mod=vendor
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

ifdef VERBOSE
	TESTARGS="-v"
endif

ifeq ($(RICHGO),)
	GOBINARY=go
else
	GOBINARY=richgo
endif

build: fmt vet lint test compile ## Build the application in dev machine

compile: ensure-build-dir # compile cli
	$(GOBINARY) build -o $(APP_EXECUTABLE)

ensure-build-dir :
	mkdir -p out

fmt:
	$(GOBINARY) fmt $(ALL_PACKAGES)

lint:
	./scripts/lint $(ALL_PACKAGES)

test:
	$(GOBINARY) test $(ALL_PACKAGES) -parallel 4 $(TESTARGS) -coverprofile ./out/coverage

vet:
	$(GOBINARY) vet $(ALL_PACKAGES)

setup: # Install dependencies
ifeq ($(GOLINT),)
	GO111MODULE=off $(GOBINARY) get -u golang.org/x/lint/golint
endif
ifeq ($(RICHGO),)
	GO111MODULE=off  \$(GOBINARY) get -u github.com/kyoh86/richgo
endif
	mkdir -p $(PWD)/out




