package day10

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1_Short(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"16",
		"10",
		"15",
		"5",
		"1",
		"11",
		"7",
		"19",
		"6",
		"12",
		"4",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "35" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"28",
		"33",
		"18",
		"42",
		"31",
		"14",
		"46",
		"20",
		"48",
		"47",
		"24",
		"23",
		"49",
		"45",
		"19",
		"38",
		"39",
		"11",
		"1",
		"32",
		"25",
		"35",
		"8",
		"17",
		"7",
		"9",
		"4",
		"2",
		"34",
		"10",
		"3",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "220" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_Short(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"16",
		"10",
		"15",
		"5",
		"1",
		"11",
		"7",
		"19",
		"6",
		"12",
		"4",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "8" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"28",
		"33",
		"18",
		"42",
		"31",
		"14",
		"46",
		"20",
		"48",
		"47",
		"24",
		"23",
		"49",
		"45",
		"19",
		"38",
		"39",
		"11",
		"1",
		"32",
		"25",
		"35",
		"8",
		"17",
		"7",
		"9",
		"4",
		"2",
		"34",
		"10",
		"3",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "19208" {
		t.Fatalf("Wrong result: %s", res)
	}
}
