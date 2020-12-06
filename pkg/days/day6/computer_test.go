package day6

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1_SingleUserSingleGroup(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"abcx",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "4" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_MultipleUsersSingleGroup(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"abcx",
		"abcy",
		"abcz",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "6" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "11" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_SingleUserSingleGroup(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"abcx",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "4" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_MultipleUsersSingleGroup(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"abcx",
		"abcy",
		"abcz",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "3" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "6" {
		t.Fatalf("Wrong result: %s", res)
	}
}
