# CLI Calculator in Go

A simple command-line calculator written in Go to help learn Go basics including functions, CLI arguments, and input/output handling.

## Prerequisites

- Go programming language (version 1.16 or higher)

## Installation

1. Install Go from [https://golang.org/dl/](https://golang.org/dl/)
2. Make sure Go is added to your PATH during installation

## Usage

You can run the calculator directly with `go run`:

```bash
go run main.go <operation> <num1> <num2>
```

### Operations

- `add` - Addition
- `subtract` - Subtraction
- `multiply` - Multiplication
- `divide` - Division

### Examples

```bash
go run main.go add 5 10        # Output: 15
go run main.go subtract 20 8   # Output: 12
go run main.go multiply 3 7    # Output: 21
go run main.go divide 15 3     # Output: 5
```

## Building the Application

To build an executable:

```bash
go build -o calculator.exe main.go
```

Then you can run it directly:

```bash
calculator.exe add 5 10
```

### Using the Batch File

On Windows, you can also use the provided batch file to test the application:

```bash
run_calculator.bat
```

### Using Makefile

If you have `make` installed:

```bash
make add      # Run addition example
make subtract # Run subtraction example
make multiply # Run multiplication example
make divide   # Run division example
make test     # Run tests
```

## Error Handling

The calculator handles various error cases:

- Insufficient arguments
- Invalid number formats
- Division by zero
- Unsupported operations

## Learning Concepts Covered

This simple application demonstrates several Go basics:

- Package management
- Command-line argument parsing
- Function definitions
- Error handling
- Switch statements
- Basic arithmetic operations
- Input/output operations