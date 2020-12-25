package main

import (
	"fmt"
	"log"
	"time"

	"github.com/codingnagger/advent-of-code-2020/pkg/days/day24"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"
)

func main() {
	start := time.Now()

	today := &day24.Computer{}

	input := inputparser.ReadInput("./assets/input/day24.txt")

	res, err := today.Part2(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", res)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution took %s\n", elapsed)
}
