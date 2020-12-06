package day5

import (
	"fmt"
	"sort"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

var (
	errorValueNotFound = fmt.Errorf("Value not found")
)

// Computer of the Advent of code 2020 Day 5
type Computer struct {
}

type pass struct {
	row  int
	seat int
	ID   int
}

// Part1 of Day5
func (d *Computer) Part1(input days.Input) (days.Result, error) {

	maxID := parsePass(input[0]).ID

	for _, encodedPass := range input {
		pass := parsePass(encodedPass)
		if pass.ID > maxID {
			maxID = pass.ID
		}
	}

	return days.Result(fmt.Sprint(maxID)), nil
}

// Part2 of Day5
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	passesIDs := []int{}

	for _, encodedPass := range input {
		passesIDs = append(passesIDs, parsePass(encodedPass).ID)
	}

	sort.Ints(passesIDs)

	for i := 0; i < len(passesIDs)-1; i++ {
		if passesIDs[i+1] != passesIDs[i]+1 {
			return days.Result(fmt.Sprint(passesIDs[i] + 1)), nil
		}
	}

	return "", errorValueNotFound
}

func parsePass(seat string) pass {
	res := pass{}

	res.row, _ = findRow(seat[:7])
	res.seat, _ = findSeat(seat[7:])
	res.ID = res.row*8 + res.seat
	return res
}

func findRow(binaryRow string) (int, error) {
	return find(binaryRow, 'F', 'B', 128)
}

func findSeat(binarySeat string) (int, error) {
	return find(binarySeat, 'L', 'R', 8)
}

func find(path string, left rune, right rune, size int) (int, error) {
	min, max := 0, size-1

	for _, c := range path {
		newSize := size / 2

		if c == left {
			max -= newSize
		} else {
			min += newSize
		}

		size = newSize

		if size == 1 {
			return min, nil
		}
	}

	return 0, errorValueNotFound
}
