package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 14
type Computer struct {
}

// Part1 of Day 14
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	return days.Result(fmt.Sprint(runGame(input, 2020))), nil
}

// Part2 of Day 14
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	return days.Result(fmt.Sprint(runGame(input, 30000000))), nil
}

func runGame(input days.Input, lastTurn int) int {
	startingNumbers := parseStartingNumbers(input)

	spokenTurns := make(map[int]int, len(startingNumbers))
	spokenCounts := make(map[int]int, len(startingNumbers))
	previousSpokenTurns := make(map[int]int, len(startingNumbers))

	cursor := 0
	var current, previous int

	for cursor < lastTurn {
		previous = current

		if cursor < len(startingNumbers) {
			current = startingNumbers[cursor]
		} else {
			count := spokenCounts[previous]

			if count == 1 {
				current = 0
			} else if previousSpokenTurns[previous] == 0 {
				current = spokenTurns[previous] - count + 1
			} else {
				current = spokenTurns[previous] - previousSpokenTurns[previous]
			}
		}

		_, ok := spokenCounts[current]

		if ok {
			spokenCounts[current]++
		} else {
			spokenCounts[current] = 1
		}

		cursor++

		previousSpokenTurns[current] = spokenTurns[current]
		spokenTurns[current] = cursor
	}

	return current
}

func parseStartingNumbers(input days.Input) []int {
	values := strings.Split(input[0], ",")
	res := []int{}

	for _, value := range values {
		number, _ := strconv.Atoi(value)
		res = append(res, number)
	}

	return res
}
