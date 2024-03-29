.DEFAULT_GOAL := default

APP     := http-request-dump
VERSION := $(shell git describe --tags)
VERSION := $(if $(VERSION:-=),$(VERSION),unknown)

GOCMD   := $(shell which go)
GOCGO   := 0

GOTAGS     = -tags urfave_cli_no_docs
LDFLAGS    = -ldflags "-s -w -X main.version=$(VERSION)"
MAKEFLAGS += --silent

clean:
	$(GOCMD) clean -cache
	rm -rf build/$(APP)-*

fmt:
	$(GOCMD) fmt ./...

compile:
	CGO_ENABLED=$(GOCGO) GOOS=linux   GOARCH=amd64 $(GOCMD) build $(GOTAGS) $(LDFLAGS) -o build/$(APP)-linux-amd64       .
	CGO_ENABLED=$(GOCGO) GOOS=linux   GOARCH=arm64 $(GOCMD) build $(GOTAGS) $(LDFLAGS) -o build/$(APP)-linux-arm64       .
	CGO_ENABLED=$(GOCGO) GOOS=windows GOARCH=amd64 $(GOCMD) build $(GOTAGS) $(LDFLAGS) -o build/$(APP)-windows-amd64.exe .
	CGO_ENABLED=$(GOCGO) GOOS=windows GOARCH=arm64 $(GOCMD) build $(GOTAGS) $(LDFLAGS) -o build/$(APP)-windows-arm64.exe .
	CGO_ENABLED=$(GOCGO) GOOS=darwin  GOARCH=amd64 $(GOCMD) build $(GOTAGS) $(LDFLAGS) -o build/$(APP)-darwin-amd64      .
	CGO_ENABLED=$(GOCGO) GOOS=darwin  GOARCH=arm64 $(GOCMD) build $(GOTAGS) $(LDFLAGS) -o build/$(APP)-darwin-arm64      .

container:
	docker build \
	--file dockerfile \
	--tag $(shell echo $(APP) | tr A-Z a-z):$(shell echo $(VERSION)) \
	.

default: clean fmt compile;