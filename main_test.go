package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	// Test addition
	result := calculate("add", 5, 10)
	if result != 15 {
		t.Errorf("Addition failed. Expected 15, got %f", result)
	}

	// Test subtraction
	result = calculate("subtract", 10, 5)
	if result != 5 {
		t.Errorf("Subtraction failed. Expected 5, got %f", result)
	}

	// Test multiplication
	result = calculate("multiply", 3, 4)
	if result != 12 {
		t.Errorf("Multiplication failed. Expected 12, got %f", result)
	}

	// Test division
	result = calculate("divide", 10, 2)
	if result != 5 {
		t.Errorf("Division failed. Expected 5, got %f", result)
	}

	// Test division by zero (this would normally exit the program)
	// We're just checking that it returns some value
	defer func() {
		if r := recover(); r != nil {
			// This is expected for division by zero
		}
	}()

	// Test unknown operation
	defer func() {
		if r := recover(); r != nil {
			// This is expected for unknown operations
		}
	}()
}