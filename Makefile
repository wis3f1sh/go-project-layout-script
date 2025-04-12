GO ?= go
GOBUILD ?= $(GO) build
GOCLEAN ?= $(GO) clean
GOTEST ?= $(GO) test
GOGET ?= $(GO) get
GOFMT ?= $(GO) fmt

BUILDDIR ?= build

BUILDFLAGS ?= -ldflags "-s -w"

all: build

build:
	$(GOBUILD) $(BUILDFLAGS) -o $(BUILDDIR)/main main.go

test:
	$(GOTEST) -v ./...

fmt:
	$(GOFMT) -w .

clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

deps:
	$(GOGET) -u

run:
	$(GOBUILD) $(BUILDFLAGS) -o $(BUILDDIR)/main main.go
	./$(BUILDDIR)/main

.PHONY: all build test fmt clean deps run
