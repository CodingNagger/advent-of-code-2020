package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/golang-collections/collections/stack"
)

// Computer of the Advent of code 2020 Day 14
type Computer struct {
}

// Part1 of Day 14
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	var bitmask string
	memory := make(map[string]string)
	res := int64(0)

	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			bitmask = parseBitmask(line)
		} else {
			key, value := parseInstruction(line)
			memory[key] = applyBitmask(bitmask, parseBinaryStringFromDecimalString(value))
		}
	}

	for _, value := range memory {
		res += parseIntFromBinaryString(value)
	}

	return days.Result(fmt.Sprint(res)), nil
}

// Part2 of Day 14
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	var bitmask string
	memory := make(map[string]string)
	res := int64(0)

	for _, line := range input {
		if strings.HasPrefix(line, "mask") {
			bitmask = parseBitmask(line)
		} else {
			key, value := parseInstruction(line)
			addresses := parseAddressesToWriteTo(bitmask, key)

			for _, address := range addresses {
				memory[address] = value
			}
		}
	}

	for _, value := range memory {
		v, _ := strconv.ParseInt(value, 10, 64)
		res += v
	}

	return days.Result(fmt.Sprint(res)), nil
}

func parseAddressesToWriteTo(bitmask string, decimalAddress string) []string {
	res := []string{}

	s := stack.New()
	s.Push(applyV2Bitmask(bitmask, parseBinaryStringFromDecimalString(decimalAddress)))

	for s.Len() > 0 {
		current := s.Pop().(string)

		if strings.Contains(current, "X") {
			s.Push(strings.Replace(current, "X", "1", 1))
			s.Push(strings.Replace(current, "X", "0", 1))
		} else {
			res = append(res, current)
		}
	}

	return res
}

func applyV2Bitmask(bitmask string, decimalAddress string) string {
	mask := []rune(reverseString(bitmask))
	number := []rune(reverseString(decimalAddress))
	res := make([]rune, len(bitmask))
	nn := len(number)
	nm := len(mask)

	for i := 0; i < nn; i++ {
		if mask[i] != '0' {
			res[i] = mask[i]
		} else {
			res[i] = number[i]
		}
	}

	for i := nn; i < nm; i++ {
		if mask[i] != '0' {
			res[i] = mask[i]
		} else {
			res[i] = '0'
		}
	}

	r := reverseString(string(res))
	return r
}

func parseBitmask(line string) string {
	return strings.TrimSpace(strings.Split(line, "=")[1])
}

// mem[7] = 101 returns 7, 101
func parseInstruction(line string) (key string, value string) {
	halves := strings.Split(line, "=")

	value = strings.TrimSpace(halves[1])

	memBlock := strings.TrimSpace(halves[0])

	key = memBlock[4 : len(memBlock)-1]

	return
}

func parseBinaryStringFromDecimalString(decimal string) string {
	intVal, _ := strconv.ParseInt(decimal, 10, 64)
	return strconv.FormatInt(intVal, 2)
}

func applyBitmask(bitmask string, binaryNumber string) string {
	mask := []rune(reverseString(bitmask))
	number := []rune(reverseString(binaryNumber))
	res := make([]rune, len(bitmask))
	nn := len(number)
	nm := len(mask)

	for i := 0; i < nn; i++ {
		if mask[i] != 'X' {
			res[i] = mask[i]
		} else {
			res[i] = number[i]
		}
	}

	for i := nn; i < nm; i++ {
		if mask[i] != 'X' {
			res[i] = mask[i]
		} else {
			res[i] = '0'
		}
	}

	r := reverseString(string(res))
	// fmt.Printf("applyBitmask(%s, %s) = %s / %d\n", bitmask, binaryNumber, r, parseIntFromBinaryString(r))
	return r
}

func parseIntFromBinaryString(binaryNumber string) int64 {
	res, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return res
}
func reverseString(s string) string {
	runes := []rune(s)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
	}
	return string(runes)
}
