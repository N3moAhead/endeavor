BINARY_NAME=endeavor
SEED_BINARY=seed
CSS_INPUT=./web/input.css
CSS_OUTPUT=./web/static/css/style.css
CMD_PATH=./cmd/endeavor/main.go
SEED_PATH=./cmd/seed/main.go

.DEFAULT_GOAL := dev

.PHONY: all install install-tools fmt vet build build-css build-prod run clean run-seed help dev watch-go watch-css check-air

help:
	@echo "Available commands:"
	@echo "  make  setup       - Installs all dependencies and dev tools"
	@echo "  make dev          - Start the dev environment (Air + Tailwind Watch)"
	@echo "  make install      - Install NPM packages and Go dependencies"
	@echo "  make install-tools- Install 'air' for live reloading"
	@echo "  make build-prod   - Build optimized CSS and minified Go binary"
	@echo "  make clean        - Remove binaries and generated files"
	@echo "  make seed         - Build and run the seed script"

install:
	npm install
	go mod tidy

install-tools:
	go install github.com/air-verse/air@latest

setup: install install-tools
	@echo "Echo installing dependencies and dev tooling..."

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

watch-css:
	npx @tailwindcss/cli -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch

watch-go:
	air

# COMBINED: Starts both concurrently
# 'make -j2' executes two targets in parallel
dev: check-air
	@echo "Starting development environment..."
	make -j2 watch-css watch-go

# Check if Air is actually installed
check-air:
	@command -v air >/dev/null 2>&1 || { echo >&2 "Air is not installed. Please run 'make install-tools' first."; exit 1; }

build-css:
	npx @tailwindcss/cli -i $(CSS_INPUT) -o $(CSS_OUTPUT)

build: vet
	go build -o $(BINARY_NAME) $(CMD_PATH)

run: build-css build
	@echo "Starting $(BINARY_NAME)..."
	./$(BINARY_NAME)

# --- Production Builds ---
build-css-prod:
	npx @tailwindcss/cli -i $(CSS_INPUT) -o $(CSS_OUTPUT) --minify

# Builds Go binary with linker flags to strip debug info (smaller binary)
build-prod: vet build-css-prod
	go build -ldflags="-s -w" -o $(BINARY_NAME) $(CMD_PATH)

# --- Seeding ---
build-seed: vet
	go build -o $(SEED_BINARY) $(SEED_PATH)

seed: build-seed
	@echo "Seeding database..."
	./$(SEED_BINARY)

# --- Cleanup ---
clean:
	rm -f $(BINARY_NAME)
	rm -f $(SEED_BINARY)
	rm -f $(CSS_OUTPUT)
	rm -rf tmp
	rm -rf node_modules
