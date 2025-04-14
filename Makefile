# ==============================================================================
# Define global Makefile variables for easier reference

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# Project root directory
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# Directory for build artifacts and temporary files
OUTPUT_DIR := $(ROOT_DIR)/_output

# ==============================================================================
# Define Makefile 'all' phony target. When executing `make`, the 'all' target will be executed by default
.PHONY: all
all: format build

# ==============================================================================
# Define other necessary phony targets

.PHONY: build
build: tidy # Compile source code, depends on 'tidy' target to automatically add/remove dependency packages.
	@go build -v -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/cmd/blog/main.go

.PHONY: format
format: # Format Go source code.
	@gofmt -s -w ./

.PHONY: tidy
tidy: # Automatically add/remove dependency packages.
	@go mod tidy

.PHONY: clean
clean: # Clean build artifacts, temporary files, etc.
	@-rm -vrf $(OUTPUT_DIR)