package types

import "strconv"

// BoundsChecker allows to validate that methods are between a given min and max value
type BoundsChecker struct {
	Min int
	Max int
}

// NewBoundsChecker returns a new BoundsChecker
func NewBoundsChecker(min, max int) BoundsChecker {
	if min > max {
		min, max = max, min
	}

	return BoundsChecker{min, max}
}

// Validate returns true if the value is within bounds
func (b *BoundsChecker) Validate(value int) bool {
	return b.Min <= value && b.Max >= value
}

// ValidateString returns true if the value is within bounds
func (b *BoundsChecker) ValidateString(value string) bool {
	v, err := strconv.Atoi(value)

	if err != nil {
		return false
	}

	return b.Validate(v)
}
