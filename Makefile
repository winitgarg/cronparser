# Makefile for Cron Parser Project

# Binary name
BINARY_NAME=cronparser

# Test coverage profile
COVER_PROFILE=coverage.out

# List of targets
.PHONY: all install-go build run test test-cover test-coverage test-coverage-with-file clean help

help:
	@echo "Makefile Usage:"
	@echo "  make all                     - Install Go (macOS) and build the Go binary"
	@echo "  make install-go              - Check and install Go if it's not already installed and the OS is macOS"
	@echo "  make build                   - Build the Go binary"
	@echo "  make run                     - Display the usage of the binary"
	@echo "  make test                    - Run the test cases"
	@echo "  make test-cover              - Run tests with coverage"
	@echo "  make test-coverage           - Run tests with coverage and generate an HTML report"
	@echo "  make test-coverage-with-file - Output coverage profile information for each function"
	@echo "  make clean                   - Clean up generated files"

# Check if Go is installed and install it if it's not and the OS is macOS
install-go:
	@if ! which go > /dev/null 2>&1; then \
		UNAME_S := $(shell uname -s); \
		if [ "$$UNAME_S" = "Darwin" ]; then \
			echo "Go is not installed. Installing Go on macOS..."; \
			brew install go; \
			if [ $$? -ne 0 ]; then \
				echo "Error: Homebrew is required to install Go. Please install Homebrew first."; \
				exit 1; \
			fi \
		else \
			echo "Go is not installed. Please install Go manually for your operating system."; \
			exit 1; \
		fi \
	else \
		echo "Go is already installed."; \
	fi

# Default target
all: install-go build

# Build the Go binary
build: install-go
	@echo "Building the Go binary..."
	go build -o $(BINARY_NAME) cmd/main.go

# Prints the usage
run:
	@echo "Usage: ./cronparser \"*/15 0 1,15 * 1-5 /usr/bin/find\""

# Run test cases
test:
	@echo "Running test cases..."
	go test -v ./...

# Run tests and collect coverage data
test-cover:
	@echo "Running test cases with coverage..."
	go test -cover ./...

test-coverage:
	@echo "Running test cases with coverage..."
	go test -coverprofile=$(COVER_PROFILE) ./...
	@echo "Generating coverage report..."
	go tool cover -html=$(COVER_PROFILE)

test-coverage-with-file:
	@echo "Running test cases to show coverage for each file..."
	go tool cover -func=coverage.out

# Clean up generated files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME) $(COVER_PROFILE)