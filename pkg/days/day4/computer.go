package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/types"
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

	birthYearChecker      types.BoundsChecker
	issueYearChecker      types.BoundsChecker
	expirationYearChecker types.BoundsChecker
	inHeightChecker       types.BoundsChecker
	cmHeightChecker       types.BoundsChecker
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

func newPassport() passport {
	return passport{
		birthYearChecker:      types.BoundsChecker{Min: 1920, Max: 2002},
		issueYearChecker:      types.BoundsChecker{Min: 2010, Max: 2020},
		expirationYearChecker: types.BoundsChecker{Min: 2020, Max: 2030},
		cmHeightChecker:       types.BoundsChecker{Min: 150, Max: 193},
		inHeightChecker:       types.BoundsChecker{Min: 59, Max: 76},
	}
}

func parsePassports(input days.Input) []passport {
	currentPassport := newPassport()
	passports := []passport{}

	for _, line := range input {
		err := parsePassportFields(&currentPassport, line)

		if err == errNoMoreFields {
			passports = append(passports, currentPassport)
			currentPassport = newPassport()
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
	return p.birthYearChecker.ValidateString(p.birthYear)
}

func (p passport) hasValidIssueYear() bool {
	return p.issueYearChecker.ValidateString(p.issueYear)
}

func (p passport) hasValidExpirationYear() bool {
	return p.expirationYearChecker.ValidateString(p.expirationYear)
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
	return p.cmHeightChecker.ValidateString(p.getHeightNumber())
}

func (p passport) hasValidInHeight() bool {
	return p.inHeightChecker.ValidateString(p.getHeightNumber())
}

func (p passport) getHeightNumber() string {
	return p.height[:len(p.height)-2]
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
