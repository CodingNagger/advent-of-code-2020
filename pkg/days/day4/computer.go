package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 4
type Computer struct {
}

var errNoMoreFields = fmt.Errorf("No fields parsed")

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

// Part2 of Day4
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	count := 0

	passports := parsePassports(input)

	for _, passport := range passports {
		if passport.isValid() {
			count++
		}
	}

	return days.Result(fmt.Sprint(count)), nil
}

// Part1 of Day4
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	count := 0

	passports := parsePassports(input)

	for _, passport := range passports {
		if passport.hasMandatoryFields() {
			count++
		}
	}

	return days.Result(fmt.Sprint(count)), nil
}

func parsePassports(input days.Input) []passport {
	currentPassport := passport{}
	passports := []passport{}

	for _, line := range input {
		err := parsePassportFields(&currentPassport, line)

		if err == errNoMoreFields {
			passports = append(passports, currentPassport)
			currentPassport = passport{}
		}
	}

	passports = append(passports, currentPassport)

	return passports
}

func parsePassportFields(p *passport, line string) error {
	if len(line) == 0 {
		return errNoMoreFields
	}

	parts := strings.Fields(line)

	for _, part := range parts {
		pair := strings.Split(part, ":")

		switch pair[0] {
		case "byr":
			p.birthYear = pair[1]
		case "iyr":
			p.issueYear = pair[1]
		case "eyr":
			p.expirationYear = pair[1]
		case "hgt":
			p.height = pair[1]
		case "hcl":
			p.hairColor = pair[1]
		case "ecl":
			p.eyeColor = pair[1]
		case "pid":
			p.passportID = pair[1]
		case "cid":
			p.countryID = pair[1]
		}
	}

	return nil
}

func (p passport) hasMandatoryFields() bool {
	return len(p.birthYear) > 0 &&
		len(p.issueYear) > 0 &&
		len(p.expirationYear) > 0 &&
		len(p.height) > 0 &&
		len(p.hairColor) > 0 &&
		len(p.eyeColor) > 0 &&
		len(p.passportID) > 0
}

func (p passport) isValid() bool {
	return p.hasMandatoryFields() &&
		p.hasValidBirthYear() &&
		p.hasValidIssueYear() &&
		p.hasValidExpirationYear() &&
		p.hasValidHeight() &&
		p.hasValidHairColor() &&
		p.hasValidEyeColor() &&
		p.hasValidPassportID()
}

func (p passport) hasValidBirthYear() bool {
	b, err := strconv.Atoi(p.birthYear)

	if err != nil {
		return false
	}

	return b >= 1920 && b <= 2002
}

func (p passport) hasValidIssueYear() bool {
	b, err := strconv.Atoi(p.issueYear)

	if err != nil {
		return false
	}

	return b >= 2010 && b <= 2020
}

func (p passport) hasValidExpirationYear() bool {
	b, err := strconv.Atoi(p.expirationYear)

	if err != nil {
		return false
	}

	return b >= 2020 && b <= 2030
}

func (p passport) hasValidHeight() bool {
	if strings.HasSuffix(p.height, "in") {
		return p.hasValidInHeight()
	} else if strings.HasSuffix(p.height, "cm") {
		return p.hasValidCmHeight()
	}

	return false
}

func (p passport) hasValidCmHeight() bool {
	h, err := strconv.Atoi(p.height[:len(p.height)-2])

	if err != nil {
		return false
	}

	return h >= 150 && h <= 193
}

func (p passport) hasValidInHeight() bool {
	h, err := strconv.Atoi(p.height[:len(p.height)-2])

	if err != nil {
		return false
	}

	return h >= 59 && h <= 76
}

func (p passport) hasValidHairColor() bool {
	res, _ := regexp.MatchString("^#[a-f0-9]{6}$", p.hairColor)
	return res
}

func (p passport) hasValidEyeColor() bool {
	res, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", p.eyeColor)
	return res
}

func (p passport) hasValidPassportID() bool {
	res, _ := regexp.MatchString("^[0-9]{9}$", p.passportID)
	return res
}
