package day9

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"
)

// Computer of the Advent of code 2020 Day 9
type Computer struct {
}

// Part1 of Day 9
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	res, _ := findUnrulyNumber(25, inputparser.ParseNumbers(input))
	return days.Result(fmt.Sprint(res)), nil
}

// Part2 of Day 9
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	res, _ := findEncryptionWeakness(25, inputparser.ParseNumbers(input))
	return days.Result(fmt.Sprint(res)), nil
}

func findEncryptionWeakness(preamble int, values []int) (int, error) {
	unrulyNumber, _ := findUnrulyNumber(preamble, values)
	var sum, min, max, count int

	for i := 0; i < len(values); i++ {
		sum = values[i]
		min = values[i]
		max = values[i]
		count = 1

		for j := i + 1; j < len(values) && sum < unrulyNumber; j++ {
			sum += values[j]
			count++

			if values[j] < min {
				min = values[j]
			} else if values[j] > max {
				max = values[j]
			}
		}

		if sum == unrulyNumber && count > 1 {
			return min + max, nil
		}
	}

	return 0, fmt.Errorf("Not found")
}

func findUnrulyNumber(preamble int, values []int) (int, error) {
	for i := preamble; i < len(values); i++ {
		possibleSums := make(map[int]int)
		respectRules := false

	SumsCheck:
		for j := i - preamble; j < i; j++ {
			_, ok := possibleSums[values[j]]

			if ok {
				respectRules = true
				break SumsCheck
			} else {
				possibleSums[values[i]-values[j]] = values[j]
			}
		}

		if !respectRules {
			return values[i], nil
		}
	}
	return 0, fmt.Errorf("Not found")
}
