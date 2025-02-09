GO ?= go
GOFMT ?= gofmt "-s"
GOFILES = $(shell find . -name "*.go")

GOLANGCI_LINT_PACKAGE ?= github.com/golangci/golangci-lint/cmd/golangci-lint@latest
GOSEC_PACKAGE ?= github.com/securego/gosec/v2/cmd/gosec@latest
GOVULNCHECK_PACKAGE ?= golang.org/x/vuln/cmd/govulncheck@latest
GOSWAG_PACKAGE ?= github.com/swaggo/swag/cmd/swag@latest
GOWIRE_PACKAGE ?= github.com/google/wire/cmd/wire@latest
EDITORCONFIG_CHECKER_PACKAGE ?= github.com/editorconfig-checker/editorconfig-checker/v3/cmd/editorconfig-checker@latest

##@ Verification
.PHONY: lint
lint: tidy golangci-lint gosec govulncheck fmt

.PHONY: golangci-lint
golangci-lint:
	@echo "##### Running golangci-lint"
	go run $(GOLANGCI_LINT_PACKAGE) run -v

.PHONY: gosec
gosec:
	@echo "##### Running gosec"
	go run $(GOSEC_PACKAGE) ./...

.PHONY: govulncheck
govulncheck:
	@echo "##### Running govulncheck"
	go run $(GOVULNCHECK_PACKAGE) ./...

.PHONY: fmt
# Ensure consistent code formatting.
fmt:
	$(GOFMT) -w $(GOFILES)

##@ Run
# Build the application
.PHONY: all
all: build test

.PHONY: build
build: clean generate
	@echo "Building..."
	@go build -o main cmd/go-template/main.go cmd/go-template/wire_gen.go

# Clean the binary
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f main

# Run the application
# ref: https://github.com/google/wire/pull/363
.PHONY: run
run:
	@echo "Running..."
	@go run cmd/go-template/main.go cmd/go-template/wire_gen.go

# Create DB container
.PHONY: docker-run
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
.PHONY: docker-down
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

##@ Test
# Test the application
.PHONY: test
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Tests for the application
.PHONY: itest
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Live Reload
.PHONY: watch
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

##@ Generate
.PHONY: generate
generate:
	@echo "Generating go files..."
	@go run $(GOWIRE_PACKAGE) gen ./cmd/go-template/wire.go

# Generate Swagger docs
.PHONY: swag
swag:
	@echo "Generating Swagger docs..."
	@go run $(GO_SWAG_PACKAGE) init -g cmd/go-template/main.go --parseDependency --parseInternal

##@ Migrations
.PHONY: migrate
migrate:
	@echo "Running migrations..."
	@go run migrations/main.go

##@ deps
.PHONY: tidy
tidy:
	@echo "Tidying up..."
	@go mod tidy

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@(GO) install $(GOLANGCI_LINT_PACKAGE)
	@(GO) install $(GOSEC_PACKAGE)
	@(GO) install $(GOVULNCHECK_PACKAGE)
	@(GO) install $(GOWIRE_PACKAGE)
	@(GO) install $(GO_SWAG_PACKAGE)
	@(GO) install $(EDITORCONFIG_CHECKER_PACKAGE)
