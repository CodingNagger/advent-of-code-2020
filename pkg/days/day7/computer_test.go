package day7

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "4" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestParseRules_SingleRule(t *testing.T) {
	input := days.Input{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	}
	expectedResult := map[string]bagRuleset{
		"light red": bagRuleset{
			bagRule{1, "bright white"},
			bagRule{2, "muted yellow"},
		},
	}
	result := parseRules(input)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseRules_SingleEmptyRule(t *testing.T) {
	input := days.Input{
		"dotted black bags contain no other bags.",
	}
	expectedResult := map[string]bagRuleset{
		"dotted black": bagRuleset{},
	}
	result := parseRules(input)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestPart2_FirstPartData(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "32" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "126" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_Impractical(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"shiny gold bags contain 2 yellow bags, 2 red bags.",
		"yellow bags contain 1 black bag.",
		"black bags contain no other bags.",
		"red bags contain 2 yellow bags, 1 green bags.",
		"green bags contain 2 yellow bags.",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "24" {
		t.Fatalf("Wrong result: %s", res)
	}
}
