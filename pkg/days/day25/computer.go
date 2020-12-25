package day25

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 25
type Computer struct {
}

// Part1 of Day 25
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	publicKeys := inputparser.ParseNumbers(input)

	cardLoopSize := findLoopSize(publicKeys[0])
	doorLoopSize := findLoopSize(publicKeys[1])

	cardEncriptionKey := transform(cardLoopSize, publicKeys[1])
	doorEncriptionKey := transform(doorLoopSize, publicKeys[0])

	if cardEncriptionKey != doorEncriptionKey {
		fmt.Printf("cardEncriptionKey %d - doorEncriptionKey %d", cardEncriptionKey, doorEncriptionKey)
		panic("Keys mismatch ")
	}

	return days.Result(fmt.Sprint(cardEncriptionKey)), nil
}

// func findLoopSizeAndSubjectNumber(publicKey int) (int, int) {
// 	candidateSubjectNumber := 2

// 	for {
// 		loopSize := 0

// 		value := 1

// 		for value < publicKey {
// 			value = (value * candidateSubjectNumber) % 20201227

// 			loopSize++

// 			if value == publicKey {
// 				return loopSize, candidateSubjectNumber
// 			}
// 		}

// 		candidateSubjectNumber++
// 	}
// }

func transform(cardLoopSize, subjectNumber int) int {
	value := 1
	for i := 0; i < cardLoopSize; i++ {
		value = (value * subjectNumber) % 20201227
	}
	return value
}

func findLoopSize(publicKey int) int {
	value := 1

	loopSize := 0

	for value != publicKey {
		value = (value * 7) % 20201227

		loopSize++

		if value == publicKey {
			break
		}
	}

	return loopSize
}
