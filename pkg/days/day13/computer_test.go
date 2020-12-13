package day13

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"939",
		"7,13,x,x,59,x,31,19",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "295" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"939",
		"7,13,x,x,59,x,31,19",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1068781" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"_",
		"67,7,59,61",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "754018" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"_",
		"67,x,7,59,61",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "779210" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_3(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"_",
		"67,7,x,59,61",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1261476" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_4(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"_",
		"1789,37,47,1889",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1202161486" {
		t.Fatalf("Wrong result: %s", res)
	}
}
