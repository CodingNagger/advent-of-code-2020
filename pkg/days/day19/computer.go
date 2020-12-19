package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

var rules map[ruleReference]rule

var rulesRegex map[string]string

// Computer of the Advent of code 2020 Day 19
type Computer struct {
	messages []string
	AddLoops bool
}

type rule interface {
	validate(value string) (bool, int) // if true, returns the length of the matched string
}

type charRule struct {
	char string
}

type orRuleReference []andRuleReference
type andRuleReference []ruleReference

type ruleReference int

// Part1 of Day 19
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	d.parseInput(input)
	count := 0

	for _, message := range d.messages {
		ok, length := rules[0].validate(message)
		if ok && length == len(message) {
			count++
			// // // fmt.Printf("%s is valid\n", message)
		} else {
			// // // fmt.Printf("%s is invalid\n", message)
		}
	}

	return days.Result(fmt.Sprint(count)), nil
}

// Part2 of Day 19
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	regexp0 := d.parseInputAsRegexes(input)

	count := 0

	for _, message := range d.messages {
		// ok, length := rules[0].validate(message)
		// if ok && length == len(message) {
		if regexp0.MatchString(message) {
			count++
			// // // fmt.Printf("%s is valid against %s\n", message, createRegex(rules[0]))
		} else {
			// fmt.Printf("%s is invalid against %v\n", message, regexp0)
		}
	}

	return days.Result(fmt.Sprint(count)), nil
}

func (d *Computer) parseInputAsRegexes(input days.Input) *regexp.Regexp {
	rulesRegex = map[string]string{}
	d.messages = []string{}

	for _, line := range input {
		if isRule(line) {
			ref, parsedRule := parseRuleAsRegex(line)
			rulesRegex[ref] = parsedRule
		} else if isMessage(line) {
			d.messages = append(d.messages, line)
		}
	}

	rulesRegex["8"] = "42 +"
	rulesRegex["11"] = "42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 (?: 42 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31 )? 31"

	numberReplaceRegex := regexp.MustCompile(`\d+`)

	for {
		done := true

		for k := range rulesRegex {
			rulesRegex[k] = numberReplaceRegex.ReplaceAllStringFunc(rulesRegex[k], func(val string) string {
				done = false
				return rulesRegex[val]
			})
		}

		if done {
			break
		}
	}

	return regexp.MustCompile("^" + strings.ReplaceAll(strings.ReplaceAll(rulesRegex["0"], "\"", ""), " ", "") + "$")
}

func parseRuleAsRegex(line string) (string, string) {
	keyValue := strings.Split(line, ":")
	return keyValue[0], fmt.Sprintf("( %s )", strings.TrimSpace(keyValue[1]))
}

func isMessage(line string) bool {
	res, _ := regexp.MatchString("^(a|b)+$", line)
	return res
}

func isRule(line string) bool {
	res, _ := regexp.MatchString("\\d+: .+", line)
	return res
}

func (r charRule) validate(value string) (bool, int) {
	if value == r.char {
		return true, 1
	}

	return false, -1
}

func (r orRuleReference) validate(value string) (bool, int) {
	for _, rr := range r {
		check, length := rr.validate(value)
		if check {
			return true, length
		}
	}

	return false, -1
}

func (r andRuleReference) validate(value string) (bool, int) {
	fullLength := 0

	for _, rr := range r {
		check, length := rules[rr].validate(value[fullLength : len(value)-fullLength])
		if !check {
			return false, -1
		}

		fullLength += length
	}

	return true, fullLength
}

/*
"0: 4 1 3":     rule0,
"1: 4 4 | 3 3": rule1,
"2: 4 3 | 3 4": rule2,
"3: \"a\"":     rule3,
"4: \"b\"":     rule4,
*/
func (d *Computer) parseInput(input days.Input) {
	rules = map[ruleReference]rule{}
	d.messages = []string{}

	for _, line := range input {
		if isRule(line) {
			ref, parsedRule := parseRule(line)
			rules[ref] = parsedRule
		} else if isMessage(line) {
			d.messages = append(d.messages, line)
		}
	}
}

func parseRule(line string) (ruleReference, rule) {
	keyValue := strings.Split(line, ":")

	key, _ := strconv.Atoi(keyValue[0])
	value := strings.TrimSpace(keyValue[1])

	var parsedRule rule

	charCheck, _ := regexp.MatchString("^\"a|b\"$", value)

	if charCheck {
		parsedRule = charRule{value[1:2]}
	} else if strings.Contains(value, "|") {
		rules := strings.Split(value, "|")
		parsedRule = orRuleReference{
			parseAndRulesReferences(rules[0]),
			parseAndRulesReferences(rules[1]),
		}
	} else {
		parsedRule = parseAndRulesReferences(value)
	}

	return ruleReference(key), parsedRule
}

func parseAndRulesReferences(ruleset string) andRuleReference {
	refs := andRuleReference{}

	for _, ref := range strings.Split(strings.TrimSpace(ruleset), " ") {
		numericRef, _ := strconv.Atoi(ref)
		refs = append(refs, ruleReference(numericRef))
	}

	return refs
}
