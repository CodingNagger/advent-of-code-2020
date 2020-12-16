package types

import (
	"testing"
)

func TestNewBoundsChecker(t *testing.T) {
	var Min, Max int = 1, 10

	checker := NewBoundsChecker(Min, Max)

	if checker.Max != Max || checker.Min != Min {
		t.Fatalf("Wrong initilisation: %v", checker)
	}
}

func TestNewBoundsCheckerWhenMinMaxSwap(t *testing.T) {
	var Min, Max int = 1, 10

	checker := NewBoundsChecker(Max, Min)

	if checker.Max != Max || checker.Min != Min {
		t.Fatalf("Wrong initilisation: %v", checker)
	}
}

func TestValidate_ValidValues(t *testing.T) {
	var Min, Max int = 1, 10
	testValues := []int{Min, Max, 5}

	checker := NewBoundsChecker(Min, Max)

	for _, value := range testValues {
		if !checker.Validate(value) {
			t.Fatalf("Validation should have passed for: %d", value)
		}
	}
}

func TestValidate_InvalidValues(t *testing.T) {
	var Min, Max int = 98, 358
	testValues := []int{1, 10, 5, 558, Min - 1, Max + 1}

	checker := NewBoundsChecker(Min, Max)

	for _, value := range testValues {
		if checker.Validate(value) {
			t.Fatalf("Validation should have failed for: %d", value)
		}
	}
}

func TestValidateString_ValidValues(t *testing.T) {
	var Min, Max int = 2010, 2020
	testValues := []string{"2010", "2020"}

	checker := NewBoundsChecker(Min, Max)

	for _, value := range testValues {
		if !checker.ValidateString(value) {
			t.Fatalf("Validation should have passed for: %s", value)
		}
	}
}
