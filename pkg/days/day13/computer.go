package day13

import (
	"fmt"
	"strconv"

	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 13
type Computer struct {
}

const (
	ignoredBus = -1
)

// Part1 of Day 13
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	timestamp := parseTimestamp(input[0])
	busIds := parseBusIds(input)

	selectedBus := busIds[0]
	min := timestamp - (timestamp % selectedBus) + selectedBus

	for i := 1; i < len(busIds); i++ {
		testValue := timestamp - (timestamp % busIds[i]) + busIds[i]

		if testValue > timestamp && testValue < min {
			min = testValue
			selectedBus = busIds[i]
		}
	}

	return days.Result(fmt.Sprint((min - timestamp) * selectedBus)), nil
}

// Part2 of Day 13
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	busIds := parseBusIds(input)
	cursor := 1
	timestamp := int64(0)
	increment := busIds[0]

	for cursor < len(busIds) {
		if busIds[cursor] == ignoredBus { // forgot to skip ignored busses so input failed but tests passed weirdly
			cursor++
			continue
		}
		timestamp += int64(increment)

		if (timestamp+int64(cursor))%int64(busIds[cursor]) == int64(0) {
			increment = lcm(increment, busIds[cursor])
			cursor++
		}
	}

	return days.Result(fmt.Sprint(timestamp)), nil
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	result := a * b / gcd(a, b)

	return result
}

func parseTimestamp(time string) int {
	res, _ := strconv.Atoi(time)
	return res
}

func parseBusIds(input days.Input) []int {
	return inputparser.ParseCsvNumbers(input[1], ignoredBus)
}
