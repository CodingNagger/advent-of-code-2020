package day16

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/types"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "71" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"class: 0-1 or 4-19",
		"departure a: 120-130 or 140-150",
		"row: 0-5 or 8-19",
		"departure b: 220-230 or 240-250",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13,121,221",
		"",
		"nearby tickets:",
		"3,9,18,122,222",
		"15,1,5,123,223",
		"5,14,9,124,224",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "26741" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_ButWithRiskOfUnmatchedFields(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"departure a: 5-9 or 13-18",
		"departure b: 220-230 or 240-250",
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,14,13,221",
		"",
		"nearby tickets:",
		"3,9,15,18,222",
		"15,1,16,5,223",
		"5,14,17,9,224",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "3094" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_ButWithBiggerRiskOfUnmatchedFields(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"departure a: 1-3 or 4-5",
		"class: 0-1 or 2-3",
		"",
		"your ticket:",
		"1,8",
		"",
		"nearby tickets:",
		"1,2",
		"2,3",
		"3,4",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "8" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestFindFieldIndexes(t *testing.T) {
	testDay := &Computer{}

	testDay.load(days.Input{
		"class: 0-1 or 4-19",
		"departure a: 120-130 or 140-150",
		"row: 0-5 or 8-19",
		"departure b: 220-230 or 240-250",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13,121,221",
		"",
		"nearby tickets:",
		"3,9,18,122,222",
		"15,1,5,123,223",
		"5,14,9,124,224",
	}, false)

	fields := testDay.findFieldIndexes()

	if fields["departure a"] != 3 || fields["departure b"] != 4 {
		t.Fatalf("Wrong result: %v", fields)
	}
}

func TestParseRule(t *testing.T) {
	expectedResult := rule{
		name: "row",
		ranges: []types.BoundsChecker{
			{Min: 6, Max: 11},
			{Min: 33, Max: 44},
		},
	}
	result := parseRule("row: 6-11 or 33-44")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestIsRule(t *testing.T) {
	result := isRule("row row: 6-11 or 33-44")

	if !result {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestIsTicket(t *testing.T) {
	result := isTicket("40,4,50")

	if !result {
		t.Fatalf("Wrong result: %v", result)
	}
}
