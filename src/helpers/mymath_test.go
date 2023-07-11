package mymath

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(12, 27)
	expected := 39
	if result != expected {
		t.Errorf("Sum result was incorrect, got: %d, expected: %d", result, expected)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(27, 12)
	expected := 15
	if result != expected {
		t.Errorf("Substract result was incorrect, got: %d, expected: %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(12, 27)
	expected := 324
	if result != expected {
		t.Errorf("Multiply result was incorrect, got: %d, expected: %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(27, 12)
	expected := float32(2.25)
	if result != expected {
		t.Errorf("Divide result was incorrect, got: %f, expected: %f", result, expected)
	}
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Divide did not panic as expected")
		}
	}()
	_ = Divide(27, 0)
}
