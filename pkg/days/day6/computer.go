package day6

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 6
type Computer struct {
}

// Part1 of Day 6
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	return days.Result(fmt.Sprint(yesCountAnyone(input))), nil
}

// Part2 of Day 6
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	return days.Result(fmt.Sprint(yesCountEveryone(input))), nil
}

func yesCountAnyone(input days.Input) int {
	res := 0

	groupYes := make(map[rune]bool)

	for _, line := range input {
		if len(line) == 0 {
			res += len(groupYes)
			groupYes = make(map[rune]bool)
		} else {
			for _, char := range line {
				groupYes[char] = true
			}
		}
	}

	res += len(groupYes)

	return res
}

func yesCountEveryone(input days.Input) int {
	res := 0

	groupYes := make(map[rune]int)
	groupSize := 0

	for _, line := range input {
		if len(line) == 0 {
			res += countGroupYes(groupYes, groupSize)
			groupYes = make(map[rune]int, groupSize)
			groupSize = 0
		} else {
			for _, char := range line {
				v, ok := groupYes[char]

				if ok {
					groupYes[char] = v + 1
				} else {
					groupYes[char] = 1
				}
			}

			groupSize++
		}
	}

	res += countGroupYes(groupYes, groupSize)

	return res
}

func countGroupYes(groupYes map[rune]int, size int) int {
	count := 0
	for _, yesCount := range groupYes {
		if yesCount == size {
			count++
		}
	}

	return count
}
