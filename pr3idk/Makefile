# Makefile for Go Project

# The Go binary name
BINARY_NAME=app

# The Go source files directory
SRC_DIR=./cmd/app

# Go build flags
GO_FLAGS=-v

# Go test flags
TEST_FLAGS=-v

# Default target: build the binary
all: build

# Build the Go binary
build:
	@echo "Building the Go binary..."
	go build $(GO_FLAGS) -o $(BINARY_NAME) $(SRC_DIR)

# Run the Go application
run:
	@echo "Running the Go application..."
	go build $(GO_FLAGS) -o .bin/$(BINARY_NAME) $(SRC_DIR)
	.bin/$(BINARY_NAME)

# Clean the project (remove compiled binary)
clean:
	@echo "Cleaning up..."
	rm -f .bin/$(BINARY_NAME)

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy
