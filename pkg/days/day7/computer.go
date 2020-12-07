package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/golang-collections/collections/stack"
)

// Computer of the Advent of code 2020 Day 7
type Computer struct {
}

type bagRule struct {
	maxCount int
	colour   string
}

type bagRuleset []bagRule

// Part1 of Day 7
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	colours := findPossibleOuterContainerColours(parseRules(input), "shiny gold")

	return days.Result(fmt.Sprint(len(colours))), nil
}

// Part2 of Day 7
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	count := countBags(parseRules(input), bagRule{1, "shiny gold"})
	return days.Result(fmt.Sprint(count)), nil
}

func countBags(ruleset map[string]bagRuleset, start bagRule) int {
	res := 0

	s := stack.New()
	s.Push(start)

	for s.Len() > 0 {
		current := s.Pop().(bagRule)

		rules := ruleset[current.colour]

		for _, rule := range rules {
			factor := rule.maxCount * current.maxCount
			res += factor
			s.Push(bagRule{factor, rule.colour})
		}
	}

	return res
}

func findPossibleOuterContainerColours(ruleset map[string]bagRuleset, start string) []string {
	coloursFound := make(map[string]bool)
	visited := make(map[string]bool)
	s := stack.New()
	s.Push(start)

	for s.Len() > 0 {
		current := fmt.Sprint(s.Pop())

		if visited[current] {
			continue
		}

		visited[current] = true

		for key, rules := range ruleset {
			for _, rule := range rules {
				if rule.colour == current {
					s.Push(key)
					coloursFound[key] = true
				}
			}
		}
	}

	res := []string{}

	for colour := range coloursFound {
		res = append(res, colour)
	}

	return res
}

// light red bags contain 1 bright white bag, 2 muted yellow bags.
func parseRules(input days.Input) map[string]bagRuleset {
	res := make(map[string]bagRuleset)

	for _, rule := range input {
		keyValue := strings.Split(strings.ReplaceAll(strings.ReplaceAll(rule, "bags", ""), "bag", ""), "contain")
		key := strings.TrimSpace(keyValue[0])
		values := strings.Split(strings.ReplaceAll(keyValue[1], ".", ""), ",")
		rules := bagRuleset{}

		for _, value := range values {
			trimmedValue := strings.TrimSpace(value)

			if trimmedValue != "no other" {
				countAndColour := strings.SplitN(trimmedValue, " ", 2)
				count, _ := strconv.Atoi(countAndColour[0])
				rules = append(rules, bagRule{count, countAndColour[1]})
			}
		}

		res[key] = rules
	}

	return res
}
