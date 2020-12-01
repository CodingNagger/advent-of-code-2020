package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/days/day1"
)

func main() {
	today := &day1.Computer{}

	input := readInput("./assets/input/day1.txt")

	res, err := today.Part2(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func readInput(filename string) days.Input {
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
