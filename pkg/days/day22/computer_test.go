package day22

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "306" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "291" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_InfiniteLoopCheck(t *testing.T) {
	testDay := &Computer{}

	_, _ = testDay.Part2(days.Input{
		"Player 1:",
		"43",
		"19",
		"",
		"Player 2:",
		"2",
		"29",
		"14",
	})
}

// func TestPart2_RealStuff(t *testing.T) {
// 	testDay := &Computer{}

// 	res, err := testDay.Part2(inputparser.ReadInput("../../../assets/input/day22.txt"))

// 	if err != nil {
// 		t.Fatalf(err.Error())
// 	}

// 	t.Logf("Result: %s", res)
// }
