package day8refactored

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/gameconsole"
)

// Computer of the Advent of code 2020 Day 8
type Computer struct {
}

// Part1 of Day 8
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	g, err := gameconsole.Load(input)

	if err != nil {
		return "", err
	}

	res, _ := g.Execute()

	return days.Result(fmt.Sprint(res)), nil
}

// Part2 of Day 8
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	g, err := gameconsole.Load(input)

	if err != nil {
		return "", err
	}

	res, _ := g.ExecutePermutationsAndFix()

	return days.Result(fmt.Sprint(res)), nil
}
