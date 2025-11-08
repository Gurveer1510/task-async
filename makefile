# Binary name
BINARY_NAME=task-scheduler

# Build output directory
BIN_DIR=bin

# Go command
GO=go

# Default target
all: build

# Build the binary
build:
	@echo "🚀 Building $(BINARY_NAME)..."
	$(GO) build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/server
	@echo "✅ Build complete!"

# Run the app
run:
	@echo "🏃 Running $(BINARY_NAME)..."
	$(BIN_DIR)/$(BINARY_NAME)

# Run with fresh build
dev: build run

# Clean up build artifacts
clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BIN_DIR)
	@echo "✨ Done!"

# Run all tests
test:
	@echo "🧪 Running tests..."
	$(GO) test ./... -v
