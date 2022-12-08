.PHONY: build lint fmt test

BINDIR:=bin
BINARY:=$(BINDIR)/sgit
DEFAULT_GOAL:=$(BINARY)

$(BINARY): fmt lint test
	@echo "Building..."
	go build -o $(BINARY) ./cmd/sgit

build-min:
	@echo "Building..."
	go build -o $(BINARY) ./cmd/sgit

build: $(BINARY)

lint:
	@echo "Linting..."
	go vet ./...

fmt:
	@echo "Formatting..."
	go fmt ./...

test:
	@echo "Testing..."
	go test -v ./...
