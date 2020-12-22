package inputparser

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// ParseNumbers returns a list of int from a list of string
func ParseNumbers(lines []string) []int {
	values := []int{}

	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}

	return values
}

// ParseCsvNumbers returns a list of int from CSV string, if a number parsing fails it is replaced with fallbackValue
func ParseCsvNumbers(line string, fallbackValue int) []int {
	entries := strings.Split(line, ",")
	ids := []int{}

	for _, entry := range entries {
		numberValue, err := strconv.Atoi(entry)

		if err != nil {
			ids = append(ids, fallbackValue)
		} else {
			ids = append(ids, numberValue)
		}
	}

	return ids
}

// ReadInput loads input from a filename
func ReadInput(filename string) days.Input {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
