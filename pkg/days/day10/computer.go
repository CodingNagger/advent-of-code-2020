package day10

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

var countedPaths map[int]int

// Computer of the Advent of code 2020 Day 9
type Computer struct {
}

// Part1 of Day 9
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	res := calculateJoltageDistributionProduct(createSortedIntegersFromInput(input))
	return days.Result(fmt.Sprint(res)), nil
}

// Part2 of Day 9
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	return days.Result(fmt.Sprint(countPathsFromStart(input))), nil
}

func calculateJoltageDistributionProduct(joltages []int) int {
	sort.IntSlice(joltages).Sort()
	distribution := map[int]int{
		1: 1,
		3: 1,
	}

	for i := 1; i < len(joltages); i++ {
		diff := joltages[i] - joltages[i-1]

		distribution[diff]++
	}

	return distribution[1] * distribution[3]
}

func countPathsFromStart(input days.Input) int {
	countedPaths = make(map[int]int)

	joltages := createSortedIntegersFromInput(input)

	road := make(map[int]bool)

	for _, joltage := range joltages {
		road[joltage] = true
	}

	road[joltages[len(joltages)-1]+3] = true

	return countPathsFrom(0, road)
}

func countPathsFrom(start int, road map[int]bool) int {
	count, visited := countedPaths[start]

	if visited {
		return count
	}

	res := 0
	noCandidate := true

	for i := 1; i <= 3; i++ {
		candidate := start + i
		_, ok := road[candidate]
		if ok {
			res += countPathsFrom(candidate, road)
			noCandidate = false
		}
	}

	if noCandidate {
		res++
	}

	countedPaths[start] = res

	return res
}

func createSortedIntegersFromInput(input days.Input) []int {
	values := []int{}

	for _, line := range input {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}

	sort.IntSlice(values).Sort()

	return values
}
