package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/types"
)

// Computer of the Advent of code 2020 Day 16
type Computer struct {
	rules         []rule
	myTicket      ticket
	nearbyTickets []ticket
}

type rule struct {
	name   string
	ranges []types.BoundsChecker
}

type ticket []int

// Part1 of Day 16
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	d.load(input, true)
	return days.Result(fmt.Sprint(d.calculcateScanningErrorRate())), nil
}

// Part2 of Day 16
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	d.load(input, false)
	fields := d.findFieldIndexes()

	res := 1

	for key, index := range fields {
		if strings.HasPrefix(key, "departure") {
			res *= d.myTicket[index]
		}
	}
	return days.Result(fmt.Sprint(res)), nil
}

func (d *Computer) findFieldIndexes() map[string]int {
	ticketSize := len(d.myTicket)

	res := make(map[string]int, ticketSize)
	lockedFields := make(map[int]bool, ticketSize)

	for len(res) != ticketSize {
		matches := make(map[string]ticket, ticketSize)

		for _, rule := range d.rules {
			for i := 0; i < ticketSize; i++ {
				if lockedFields[i] {
					continue
				}

				allValid := true

				for _, ticket := range d.nearbyTickets {
					allValid = allValid && rule.validate(ticket[i])
				}

				if allValid {
					indexes, ok := matches[rule.name]

					if ok {
						matches[rule.name] = append(indexes, i)
					} else {
						matches[rule.name] = []int{i}
					}
				}
			}

			if len(matches[rule.name]) == 1 {
				lockedIndex := matches[rule.name][0]
				res[rule.name] = lockedIndex
				lockedFields[lockedIndex] = true
			}
		}
	}

	return res
}

func (rule *rule) validate(field int) bool {
	isValid := false
	for _, r := range rule.ranges {
		isValid = isValid || r.Validate(field)
	}
	return isValid
}

func (d *Computer) load(input days.Input, keepInvalid bool) {
	d.rules = []rule{}
	d.myTicket = ticket{}
	d.nearbyTickets = []ticket{}

	for _, line := range input {
		if isRule(line) {
			d.rules = append(d.rules, parseRule(line))
		} else if isTicket(line) {
			if len(d.myTicket) == 0 {
				d.myTicket = parseTicketFromLine(line)
			} else {
				ticket := parseTicketFromLine(line)
				ok, _ := d.validateTicket(ticket)

				if ok || keepInvalid {
					d.nearbyTickets = append(d.nearbyTickets, ticket)
				}
			}
		}
	}
}

func (d *Computer) calculcateScanningErrorRate() int {
	res := 0

	for _, ticket := range d.nearbyTickets {
		_, badFields := d.validateTicket(ticket)

		for _, field := range badFields {
			res += field
		}
	}

	return res
}

func (d *Computer) validateTicket(ticket []int) (bool, []int) {
	badFields := []int{}

	for _, field := range ticket {
		validForAny := false

		for _, rule := range d.rules {
			validForAny = validForAny || rule.validate(field)
		}

		if !validForAny {
			badFields = append(badFields, field)
		}
	}

	return len(badFields) == 0, badFields
}

func isTicket(line string) bool {
	res, _ := regexp.MatchString("^\\d+(,\\d+)*$", line)
	return res
}

func isRule(line string) bool {
	res, _ := regexp.MatchString("^\\w+( \\w+)*: \\d+-\\d+ or \\d+-\\d+$", line)
	return res
}

// row: 6-11 or 33-44
func parseRule(line string) rule {
	halves := strings.Split(line, ":")

	name := halves[0]
	rangesHalves := strings.Split(halves[1], "or")
	ranges := []types.BoundsChecker{}

	for _, rangeHalf := range rangesHalves {
		minMax := strings.Split(rangeHalf, "-")

		min, _ := strconv.Atoi(strings.TrimSpace(minMax[0]))
		max, _ := strconv.Atoi(strings.TrimSpace(minMax[1]))

		ranges = append(ranges, types.BoundsChecker{Min: min, Max: max})
	}

	return rule{name, ranges}
}

func parseTicketFromLine(line string) []int {
	return inputparser.ParseCsvNumbers(line, 0)
}
