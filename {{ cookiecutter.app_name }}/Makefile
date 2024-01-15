ROOT_DIR?=$(shell git rev-parse --show-toplevel)
APP_NAME = $(shell basename $(CURDIR))
BUILD_DIR = build

.PHONY: linux
linux:
	@env GOOS="linux" CGO_ENABLED="0" GOARCH="amd64" go build -o  $(ROOT_DIR)/$(BUILD_DIR)/$(APP_NAME)-linux64 -ldflags="-s -w" $(ROOT_DIR)/main.go
.PHONY: windows
windows:
	@env GOOS="windows" CGO_ENABLED="0" GOARCH="amd64" go build -o  $(ROOT_DIR)/$(BUILD_DIR)/$(APP_NAME).exe -ldflags="-s -w" $(ROOT_DIR)/main.go
.PHONY: arm
arm:
	@env GOOS="linux" CGO_ENABLED="0" GOARCH="arm64" go build -o  $(ROOT_DIR)/$(BUILD_DIR)/$(APP_NAME)-arm64 -ldflags="-s -w" $(ROOT_DIR)/main.go
.PHONY: build
build:
	@make clean
	@make linux
	@make windows
	@make arm

.PHONY: clean
clean:
	@rm -rf $(ROOT_DIR)/$(BUILD_DIR)

.PHONY: all
all:
	make clean
	make build

help:
	@echo "Usage: make <command>"
	@echo "Commands:"
	@echo "  help        Show this help message"
	@echo "  linux       Build the application linux"
	@echo "  windows     Build the application windows"
	@echo "  arm         Build the application arm"
	@echo "  build       Build the application all"
	@echo "  clean       Clean the build artifacts"
	@echo "  all         Clean the build artifacts dir and rebuild"