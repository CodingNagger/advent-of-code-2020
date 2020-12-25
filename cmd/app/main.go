package main

import (
	"fmt"
	"log"
	"time"

	"github.com/codingnagger/advent-of-code-2020/pkg/days/day25"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"
)

func main() {
	start := time.Now()

	today := &day25.Computer{}

	input := inputparser.ReadInput("./assets/input/day25.txt")

	res, err := today.Part1(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", res)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
