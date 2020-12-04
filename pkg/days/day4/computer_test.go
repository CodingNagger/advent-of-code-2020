package day4

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func testInput() days.Input {
	return days.Input{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}
}

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(testInput())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_OnlyInvalid(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "0" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_OnlyValid(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestParsePassports(t *testing.T) {
	res := parsePassports(testInput())

	if len(res) != 4 {
		t.Fatalf("Wrong size: %d", len(res))
	}
}

func TestParsePassportsFields(t *testing.T) {
	p := passport{}
	err := parsePassportFields(&p,
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if p.eyeColor != "gry" {
		t.Fatalf("Wrong eye color: %s", p.eyeColor)
	}

	if p.passportID != "860033327" {
		t.Fatalf("Wrong passport id: %s", p.passportID)
	}

}

func TestPart2_InvalidPassports(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "0" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2_ValidPassports(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "4" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestValidBirthYear(t *testing.T) {
	validBirthYears := []string{"1920", "2002", "1991"}
	invalidBirthYears := []string{"1919", "2003", "2158", "1852"}

	for _, year := range validBirthYears {
		if !(passport{birthYear: year}).hasValidBirthYear() {
			t.Fatalf("birthYear - Should have accepted %s", year)
		}
	}

	for _, year := range invalidBirthYears {
		if (passport{birthYear: year}).hasValidBirthYear() {
			t.Fatalf("birthYear - Should have rejected %s", year)
		}
	}
}

func TestValidIssueYear(t *testing.T) {
	validIssueYears := []string{"2010", "2016", "2020"}
	invalidIssueYears := []string{"2009", "1919", "2003", "2158",
		"1852", "2021"}

	for _, year := range validIssueYears {
		if !(passport{issueYear: year}).hasValidIssueYear() {
			t.Fatalf("issueYear - Should have accepted %s", year)
		}
	}

	for _, year := range invalidIssueYears {
		if (passport{issueYear: year}).hasValidIssueYear() {
			t.Fatalf("issueYear - Should have rejected %s", year)
		}
	}
}

func TestValidExpirationYear(t *testing.T) {
	validExpirationYears := []string{"2020", "2024", "2030"}
	invalidExpirationYears := []string{"2009", "2019", "2031", "2158"}

	for _, year := range validExpirationYears {
		if !(passport{expirationYear: year}).hasValidExpirationYear() {
			t.Fatalf("expirationYear - Should have accepted %s", year)
		}
	}

	for _, year := range invalidExpirationYears {
		if (passport{expirationYear: year}).hasValidExpirationYear() {
			t.Fatalf("expirationYear - Should have rejected %s", year)
		}
	}
}

func TestHeight(t *testing.T) {
	validHeights := []string{"60in", "190cm", "59in", "76in", "150cm", "193cm"}
	invalidHeights := []string{"190in", "190", "59", "58in", "194cm"}

	for _, height := range validHeights {
		if !(passport{height: height}).hasValidHeight() {
			t.Fatalf("expirationYear - Should have accepted %s", height)
		}
	}

	for _, height := range invalidHeights {
		if (passport{height: height}).hasValidHeight() {
			t.Fatalf("expirationYear - Should have rejected %s", height)
		}
	}
}

func TestHairColor(t *testing.T) {
	validColors := []string{"#123abc"}
	invalidColors := []string{"#123abz", "123abc"}

	for _, color := range validColors {
		if !(passport{hairColor: color}).hasValidHairColor() {
			t.Fatalf("hairColor - Should have accepted %s", color)
		}
	}

	for _, color := range invalidColors {
		if (passport{hairColor: color}).hasValidHairColor() {
			t.Fatalf("hairColor - Should have rejected %s", color)
		}
	}
}

func TestEyeColor(t *testing.T) {
	validColors := []string{"brn"}
	invalidColors := []string{"ttt", "wat"}

	for _, color := range validColors {
		if !(passport{eyeColor: color}).hasValidEyeColor() {
			t.Fatalf("eyeColor - Should have accepted %s", color)
		}
	}

	for _, color := range invalidColors {
		if (passport{eyeColor: color}).hasValidEyeColor() {
			t.Fatalf("eyeColor - Should have rejected %s", color)
		}
	}
}

func TestPassportID(t *testing.T) {
	validPassportIDs := []string{"000000001"}
	invalidPassportIDs := []string{"0123456789", "0000000001"}

	for _, passportid := range validPassportIDs {
		if !(passport{passportID: passportid}).hasValidPassportID() {
			t.Fatalf("passportID - Should have accepted %s", passportid)
		}
	}

	for _, passportid := range invalidPassportIDs {
		if (passport{passportID: passportid}).hasValidPassportID() {
			t.Fatalf("passportID - Should have rejected %s", passportid)
		}
	}
}
