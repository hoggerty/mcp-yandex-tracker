# Directory of this Makefile, works both standalone (make -C tools ...)
# and included from the root Makefile (include tools/Makefile).
TOOLS_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(abspath $(TOOLS_DIR)..)
GOBIN := $(abspath $(TOOLS_DIR)bin)
GORELEASER := $(GOBIN)/goreleaser

.PHONY: install
install:
	cd $(TOOLS_DIR) && GOBIN=$(GOBIN) go install tool

$(GORELEASER):
	$(MAKE) install

.PHONY: release-check
release-check: $(GORELEASER)
	cd $(ROOT_DIR) && $(GORELEASER) check

.PHONY: release-snapshot
release-snapshot: $(GORELEASER)
	cd $(ROOT_DIR) && $(GORELEASER) release --snapshot --clean

.PHONY: tools-clean
tools-clean:
	rm -rf $(GOBIN)
