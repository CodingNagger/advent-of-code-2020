package day2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/types"
)

// Computer of the Advent of code 2020 Day 1
type Computer struct {
}

type policyCheck struct {
	bounds   types.BoundsChecker
	char     rune
	password string
}

// Part2 of Day1
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	res := 0

	for _, candidate := range input {
		policyCheck := createPolicyCheck(candidate)

		if policyCheck.isValidWithCorrectPolicy() {
			res++
		}
	}

	return days.Result(fmt.Sprint(res)), nil
}

// Part1 of Day1
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	res := 0

	for _, candidate := range input {
		policyCheck := createPolicyCheck(candidate)

		if policyCheck.isValid() {
			res++
		}
	}

	return days.Result(fmt.Sprint(res)), nil
}

func createPolicyCheck(candidate string) policyCheck {
	res := policyCheck{}

	// 1-3 a: abcde
	parts := strings.Split(candidate, " ")

	res.password = parts[2]
	res.char, _ = utf8.DecodeRuneInString(parts[1])

	minMax := strings.Split(parts[0], "-")

	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	res.bounds = types.NewBoundsChecker(min, max)

	return res
}

func (p *policyCheck) isValid() bool {
	count := 0

	for _, char := range p.password {
		if p.char != char {
			continue
		}

		count++
	}

	return p.bounds.Validate(count)
}

func (p *policyCheck) isValidWithCorrectPolicy() bool {
	chars := []rune(p.password)

	return (chars[p.bounds.Min-1] == p.char || chars[p.bounds.Max-1] == p.char) &&
		chars[p.bounds.Min-1] != chars[p.bounds.Max-1]
}
