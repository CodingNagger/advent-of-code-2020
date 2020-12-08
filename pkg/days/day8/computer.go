package day8

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 8
type Computer struct {
}

// Part1 of Day 8
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	program := ParseProgram(input)
	program.execute()
	return days.Result(fmt.Sprint(program.accumulator)), nil
}

// Part2 of Day 8
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	multiverse := ParseProgram(input).createMultiverse()

	for _, program := range multiverse {
		program.execute()

		if !program.hitInfiniteLoop {
			return days.Result(fmt.Sprint(program.accumulator)), nil
		}
	}

	return "", fmt.Errorf("No answer found")
}
