GO ?= go
EXECUTABLE := gitea-mcp
VERSION ?= $(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')
LDFLAGS := -X "main.Version=$(VERSION)"

.PHONY: help
help: ## Print this help message.
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: build ## Install the application.
	@echo "Installing $(EXECUTABLE)..."
	@mkdir -p $(GOPATH)/bin
	@cp $(EXECUTABLE) $(GOPATH)/bin/$(EXECUTABLE)
	@echo "Installed $(EXECUTABLE) to $(GOPATH)/bin/$(EXECUTABLE)"
	@echo "Please add $(GOPATH)/bin to your PATH if it is not already there."

.PHONY: uninstall
uninstall: ## Uninstall the application.
	@echo "Uninstalling $(EXECUTABLE)..."
	@rm -f $(GOPATH)/bin/$(EXECUTABLE)
	@echo "Uninstalled $(EXECUTABLE) from $(GOPATH)/bin/$(EXECUTABLE)"

.PHONY: clean
clean: ## Clean the build artifacts.
	@echo "Cleaning up build artifacts..."
	@rm -f $(EXECUTABLE)
	@echo "Cleaned up $(EXECUTABLE)"

.PHONY: build
build: ## Build the application.
	$(GO) build -v -ldflags '-s -w $(LDFLAGS)' -o $(EXECUTABLE)

.PHONY: air
air: ## Install air for hot reload.
	@hash air > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/air-verse/air@latest; \
	fi

.PHONY: dev
dev: air ## run the application with hot reload
	air --build.cmd "make build" --build.bin ./gitea-mcp

.PHONY: vendor
vendor: ## tidy and verify module dependencies
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
