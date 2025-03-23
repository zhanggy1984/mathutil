package mathutil

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestAverage(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Average(numbers)
	if result != 3.0 {
		t.Errorf("Expected 3.0, got %f", result)
	}
}

func TestCircleArea(t *testing.T) {
	result := CircleArea(2.0)
	if result != 12.566370614359172 {
		t.Errorf("Expected 12.566370614359172, got %f", result)
	}
}
