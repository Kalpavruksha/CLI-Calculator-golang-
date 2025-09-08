# Makefile for CLI Calculator

# Build the application
build:
	go build -o calculator.exe main.go

# Run the application with sample arguments
run:
	go run main.go

# Run with specific operation
add:
	go run main.go add 5 10

subtract:
	go run main.go subtract 10 5

multiply:
	go run main.go multiply 3 4

divide:
	go run main.go divide 10 2

# Run all tests
test:
	go test -v

# Clean build artifacts
clean:
	del calculator.exe 2>nul || true

# Install dependencies (if any)
install:
	go mod tidy

# Help
help:
	@echo "CLI Calculator Makefile"
	@echo "======================"
	@echo "build    - Build the application"
	@echo "run      - Run the application (requires arguments)"
	@echo "add      - Run addition example"
	@echo "subtract - Run subtraction example"
	@echo "multiply - Run multiplication example"
	@echo "divide   - Run division example"
	@echo "test     - Run tests"
	@echo "clean    - Clean build artifacts"
	@echo "install  - Install dependencies"