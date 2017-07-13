.DEFAULT_GOAL        := help
PKG                  := dynolocker
VERSION              := $(shell cat VERSION)
GITCOMMIT            := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES  := $(shell git status --porcelain --untracked-files=no)

ifneq ($(GITUNTRACKEDCHANGES),)
GITCOMMIT := $(GITCOMMIT)-dirty
endif

CTIMEVAR          :=-X $(PKG)/version.GitCommit=$(GITCOMMIT) -X $(PKG)/version.DynolockerVersion=$(VERSION)
GO_LDFLAGS        :=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC :=-ldflags "-w $(CTIMEVAR) -extldflags -static"
GOOSES            := darwin linux
GOARCHS           := amd64

GOFMT_CMD     := $$(gofmt -w `find . -name '*.go' | grep -v vendor`)
TEST_DIRS     := $(shell find . -type f -name '*_test.go' -maxdepth 8 -exec dirname {} \; | grep -v vendor | sort -u)


help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

define test
cd $(1) && go test -v -parallel 128
endef

define cross_build
mkdir -p bin/$(1)/$(2);
GOOS=$(1) GOARCH=$(2) CGO_ENABLED=0 go build -o bin/$(1)/$(2)/$(PKG) -a -tags "static_build $(PKG)" -installsuffix $(PKG) ${GO_LDFLAGS_STATIC};
endef

.PHONY: install-deps
install-deps: ## Install deps with Glide
	@echo "Installing deps..."
	glide install

.PHONY: fmt
fmt: ## Run gofmt over all *.go files
	@echo "Running source files through gofmt..."
	$(GOFMT_CMD)

.PHONY: build
build: install-deps ## build Go binary for all GOARCH
	@echo "Building dynolocker for all GOARCH/GOOS"
	$(foreach GOARCH,$(GOARCHS),$(foreach GOOS,$(GOOSES),$(call cross_build,$(GOOS),$(GOARCH))))

.PHONY: test
test: install-deps ## Run tests
	@echo "Running tests..."
	$(foreach TEST_DIR,$(TEST_DIRS),$(call test,$(TEST_DIR)))
