SOURCEDIR=.
SOURCES = $(shell find $(SOURCEDIR) -name '*.go')
VERSION=$(shell git describe --always --tags)

.PHONY: test
test:
	go test -count=1 -cover -race ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: imports
imports:
	goimports -w -local "github.com/reef-pi" ./controller

.PHONY: build
build: clean imports vet test

clean:
	-find . -name '*.db' -exec rm {} \;
