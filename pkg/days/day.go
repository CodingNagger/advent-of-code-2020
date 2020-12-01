package days

// A Day that can compute a result from an input
type Day interface {
	Part1(Input) (Result, error)
	Part2(Input) (Result, error)
}

// An Input is used to compute a Result for a Day
type Input []string

// The Result of a Day's computation
type Result string
