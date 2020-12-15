package day15

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"0,3,6",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "436" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"0,3,6",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "175594" {
		t.Fatalf("Wrong result: %s", res)
	}
}
