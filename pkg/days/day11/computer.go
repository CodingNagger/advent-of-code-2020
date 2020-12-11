package day11

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

type seatUpdateCheck func(row int, col int, seatMap [][]rune) (bool, rune)

// Computer of the Advent of code 2020 Day 11
type Computer struct {
}

// Part1 of Day 11
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	res := executeRounds(input, shouldUpdateSeat)
	return days.Result(fmt.Sprint(res)), nil
}

// Part2 of Day 11
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	res := executeRounds(input, shouldUpdateSeatWithTolerance)
	return days.Result(fmt.Sprint(res)), nil
}

func executeRounds(input days.Input, check seatUpdateCheck) int {
	seatMap := createSeatMap(input)

	didUpdate := true

	count := 1
	// fmt.Printf("Start:\n%s\n", renderMap(seatMap))

	for didUpdate {
		seatMap, didUpdate = executeRound(seatMap, check)
		// fmt.Printf("Round %d:\n%s\n", count, renderMap(seatMap))
		count++
	}

	return countOccupiedSeats(seatMap)
}

func renderMap(seatMap [][]rune) string {
	res := ""
	for _, r := range seatMap {
		res += fmt.Sprintf("%s\n", string(r))
	}
	return res
}

func executeRound(seatMap [][]rune, check seatUpdateCheck) ([][]rune, bool) {
	didUpdate := false
	newMap := [][]rune{}

	for r, row := range seatMap {
		newRow := []rune{}

		for s := range row {
			ok, newSeat := check(r, s, seatMap)

			if ok {
				didUpdate = true
			}

			newRow = append(newRow, newSeat)
		}

		newMap = append(newMap, newRow)
	}

	return newMap, didUpdate
}

func createSeatMap(input days.Input) [][]rune {
	seatMap := [][]rune{}

	for _, line := range input {
		seatMap = append(seatMap, []rune(line))
	}

	return seatMap
}

func countOccupiedSeats(seatMap [][]rune) int {
	res := 0

	for _, row := range seatMap {
		for _, col := range row {
			if col == '#' {
				res++
			}
		}
	}

	return res
}

func countOccupiedSeatsAround(row int, col int, seatMap [][]rune) int {
	res := 0

	for r := row - 1; r <= row+1; r++ {
		if r >= 0 && r < len(seatMap) {
			for s := col - 1; s <= col+1; s++ {
				if s >= 0 && s < len(seatMap[r]) {
					if row == r && col == s {
						continue
					}

					if isSeatOccupied(r, s, seatMap) {
						res++
					}
				}
			}
		}
	}

	return res
}

func shouldUpdateSeat(row int, col int, seatMap [][]rune) (bool, rune) {
	if isSeat(row, col, seatMap) {
		occupiedSeats := countOccupiedSeatsAround(row, col, seatMap)

		if isSeatOccupied(row, col, seatMap) {
			if occupiedSeats >= 4 {
				return true, 'L'
			}
		} else {
			if occupiedSeats == 0 {
				return true, '#'
			}
		}
	}

	return false, seatMap[row][col]
}

func isSeat(row int, col int, seatMap [][]rune) bool {
	if row < 0 || row > len(seatMap) || col < 0 || col > len(seatMap[row]) {
		return false
	}

	return seatMap[row][col] != '.'
}

func isSeatOccupied(row int, col int, seatMap [][]rune) bool {
	return seatMap[row][col] == '#'
}

func shouldUpdateSeatWithTolerance(row int, col int, seatMap [][]rune) (bool, rune) {
	if isSeat(row, col, seatMap) {
		occupiedSeats := countVisibleOccupiedSeatsAround(row, col, seatMap)

		if isSeatOccupied(row, col, seatMap) {
			if occupiedSeats >= 5 {
				return true, 'L'
			}
		} else {
			if occupiedSeats == 0 {
				return true, '#'
			}
		}
	}

	return false, seatMap[row][col]
}

func countVisibleOccupiedSeatsAround(row int, col int, seatMap [][]rune) int {
	res := 0

TopLeft:
	for shift := 1; row-shift >= 0 && col-shift >= 0; shift++ {
		r := row - shift
		s := col - shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break TopLeft
		}
	}
Top:
	for shift := 1; row-shift >= 0; shift++ {
		r := row - shift
		s := col

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break Top
		}
	}
TopRight:
	for shift := 1; row-shift >= 0 && col+shift < len(seatMap[row-shift]); shift++ {
		r := row - shift
		s := col + shift

		if isSeat(r, s, seatMap) {
			// fmt.Printf("Check seat row %d - col %d\n", r, s)
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break TopRight
		}
	}
Left:
	for shift := 1; col-shift >= 0; shift++ {
		r := row
		s := col - shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break Left
		}
	}
Right:
	for shift := 1; col+shift < len(seatMap[row]); shift++ {
		r := row
		s := col + shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break Right
		}
	}
BottomLeft:
	for shift := 1; row+shift < len(seatMap) && col-shift >= 0; shift++ {
		r := row + shift
		s := col - shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break BottomLeft
		}
	}
Bottom:
	for shift := 1; row+shift < len(seatMap); shift++ {
		r := row + shift
		s := col

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break Bottom
		}
	}
BottomRight:
	for shift := 1; row+shift < len(seatMap) && col+shift < len(seatMap[row+shift]); shift++ {
		r := row + shift
		s := col + shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				// fmt.Printf("Found seat row %d - col %d\n", r, s)
				res++
			}
			break BottomRight
		}
	}

	return res
}
