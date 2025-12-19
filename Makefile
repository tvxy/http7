BINARY_NAME=http7
DIST_DIR=dist
LDFLAGS=-ldflags "-s -w"

.DEFAULT_GOAL := help

.PHONY: all build build-all clean run help

all: clean build-all

build:
	go build $(LDFLAGS) -o $(BINARY_NAME) main.go

build-all:
	mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 main.go

clean:
	rm -f $(BINARY_NAME)
	rm -rf $(DIST_DIR)

run:
	go run main.go

help:
	@echo "Available commands:"
	@echo "  make all       - Clean and build all platforms"
	@echo "  make build     - Build for local system"
	@echo "  make build-all - Build all platforms"
	@echo "  make clean     - Remove built files"
	@echo "  make run       - Run locally"
