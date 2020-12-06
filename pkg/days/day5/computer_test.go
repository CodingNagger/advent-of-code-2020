package day5

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"BFFFBBFRRR",
		"BBFFBBFRLL",
		"FFFBBBFRRR",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "820" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestFindRow(t *testing.T) {
	res, err := findRow("FBFBBFF")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != 44 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestParsePass(t *testing.T) {
	res := parsePass("FBFBBFFRLR")

	if res.ID != 357 {
		t.Fatalf("Wrong pass ID: %d", res.ID)
	}
}

func TestFindSeat(t *testing.T) {
	res, err := findSeat("RLR")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != 5 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"BBFFBBFLLL",
		"BBFFBBFLLR",
		"BBFFBBFLRR",
		"BBFFBBFRLL",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "818" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2Range_WithPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"BBFFBBFLLL",
		"BBFFBBFLLR",
		"BBFFBBFLRR",
		"BBFFBBFRLL",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "820" {
		t.Fatalf("Wrong result: %s", res)
	}
}
