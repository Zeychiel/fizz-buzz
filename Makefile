GOBASE := $(shell pwd)
GOPATH := $(shell go env GOPATH)


install: go-get -v all /cmd/main.go

build: go-build

.PHONY: run
run:
	@echo "  >  Running project..."
	@-GOPATH=$(GOPATH) go run $(GOBASE)/cmd/server/*.go

test:
	@echo "  >  Running project..."
	@-GOPATH=$(GOPATH) go test ./... -v
