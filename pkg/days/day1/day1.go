package day1

import (
	"fmt"
	"strconv"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 1
type Computer struct {
}

// Part2 of Day1
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	intinput := []int{}

	for _, i := range input {
		intval, _ := strconv.Atoi(i)
		intinput = append(intinput, intval)
	}

	for _, i := range intinput {
		for _, j := range intinput {
			if i == j {
				continue
			}

			for _, k := range intinput {
				if k == i {
					continue
				}

				if k == j {
					continue
				}

				if i+j+k == 2020 {
					return days.Result(fmt.Sprint(i * j * k)), nil
				}
			}
		}
	}

	return "", fmt.Errorf("not found")
}

// Part1 of Day1
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	intinput := []int{}

	for _, i := range input {
		intval, _ := strconv.Atoi(i)
		intinput = append(intinput, intval)
	}

	for _, i := range intinput {
		for _, j := range intinput {
			if i == j {
				continue
			}

			if i+j == 2020 {
				return days.Result(fmt.Sprint(i * j)), nil
			}
		}
	}

	return "", fmt.Errorf("not found")
}
