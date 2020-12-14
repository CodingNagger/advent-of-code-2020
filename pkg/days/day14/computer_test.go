package day14

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "165" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_MultipleMasks(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[0] = 101",
		"mem[10] = 0",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "330" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "208" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestParseAddressesToWriteTo_42(t *testing.T) {
	expectedResult := []string{
		"000000000000000000000000000000011010",
		"000000000000000000000000000000011011",
		"000000000000000000000000000000111010",
		"000000000000000000000000000000111011",
	}

	result := parseAddressesToWriteTo("000000000000000000000000000000X1101X", "42")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseAddressesToWriteTo_26(t *testing.T) {
	expectedResult := []string{
		"000000000000000000000000000000010000",
		"000000000000000000000000000000010001",
		"000000000000000000000000000000010010",
		"000000000000000000000000000000010011",
		"000000000000000000000000000000011000",
		"000000000000000000000000000000011001",
		"000000000000000000000000000000011010",
		"000000000000000000000000000000011011",
	}

	result := parseAddressesToWriteTo("00000000000000000000000000000000X0XX", "26")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseIntFromBinaryString(t *testing.T) {
	res := parseIntFromBinaryString("000000000000000000000000000001001001")
	if res != int64(73) {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestParseInstruction(t *testing.T) {
	key, value := parseInstruction("mem[7] = 101")

	if value != "101" {
		t.Fatalf("Wrong value: %s", value)
	}

	if key != "7" {
		t.Fatalf("Wrong key: %s", key)
	}
}

func TestApplyBitmask(t *testing.T) {
	testSets := [][]string{
		{"000000000000000000000000000000001011", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", "000000000000000000000000000001001001"},
		{"000000000000000000000000000001100101", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", "000000000000000000000000000001100101"},
		{"000000000000000000000000000000000000", "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", "000000000000000000000000000001000000"},
	}

	for _, testSet := range testSets {
		res := applyBitmask(testSet[1], testSet[0])
		if res != testSet[2] {
			t.Fatalf("Wrong result: %s", res)
		}
	}
}

func TestApplyV2Bitmask(t *testing.T) {
	testSets := [][]string{
		{"000000000000000000000000000000101010", "000000000000000000000000000000X1001X", "000000000000000000000000000000X1101X"},
		{"000000000000000000000000000000011010", "00000000000000000000000000000000X0XX", "00000000000000000000000000000001X0XX"},
	}

	for _, testSet := range testSets {
		res := applyV2Bitmask(testSet[1], testSet[0])
		if res != testSet[2] {
			t.Fatalf("Wrong result: %s", res)
		}
	}
}
