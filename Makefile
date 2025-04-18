# ==============================================================================
# Define global Makefile variables for easier reference

COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# Project root directory
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))
# Directory for build artifacts and temporary files
OUTPUT_DIR := $(ROOT_DIR)/_output


# ==============================================================================
# Define version-related variables

## Specify the version package used by the application, values will be injected into variables via `-ldflags -X`
VERSION_PACKAGE=blog/internal/pkg/version

## Define VERSION semantic version
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

## Check if the code repository is dirty (dirty by default)
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

# ==============================================================================
# Define Makefile 'all' phony target. When executing `make`, the 'all' target will be executed by default
.PHONY: all
	all: format build

# ==============================================================================
# Define other necessary phony targets

.PHONY: build
build: tidy # Compile source code, depends on 'tidy' target to automatically add/remove dependency packages.
	@go build -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/blog $(ROOT_DIR)/cmd/blog/main.go

.PHONY: format
format: # Format Go source code.
	@gofmt -s -w ./

.PHONY: tidy
tidy: # Automatically add/remove dependency packages.
	@go mod tidy

.PHONY: clean
clean: # Clean build artifacts, temporary files, etc.
	@-rm -vrf $(OUTPUT_DIR)