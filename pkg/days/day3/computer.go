package day3

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 1
type Computer struct {
}

type direction struct {
	dx int
	dy int
}

// Part2 of Day1
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	directions := []direction{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := 1

	for _, direction := range directions {
		res *= countTreesForDirection(input, direction)
	}

	return days.Result(fmt.Sprint(res)), nil
}

// Part1 of Day1
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	return days.Result(fmt.Sprint(countTreesForDirection(input, direction{3, 1}))), nil
}

func countTreesForDirection(input days.Input, d direction) int {
	res := 0
	x := 0
	toboggan := make([][]rune, len(input))

	for y, row := range input {
		toboggan[y] = []rune(row)
	}

	for y := d.dy; y < len(toboggan); y += d.dy {
		cols := toboggan[y]

		x += d.dx

		if x >= len(cols) {
			x = x % len(cols)
		}

		if cols[x] == '#' {
			res++
		}
	}

	return res
}
