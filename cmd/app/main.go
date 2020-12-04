package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/days/day4"
)

func main() {
	start := time.Now()

	today := &day4.Computer{}

	input := readInput("./assets/input/day4.txt")

	res, err := today.Part2(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", res)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
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
