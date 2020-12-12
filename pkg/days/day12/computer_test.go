package day12

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "25" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "286" {
		t.Fatalf("Wrong result: %s", res)
	}
}
