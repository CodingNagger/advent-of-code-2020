package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/golang-collections/collections/stack"
)

// Computer of the Advent of code 2020 Day 18
type Computer struct {
}

type expression interface {
	reduce() int
}

type constantExp int
type sumExp []expression
type productExp []expression

var lineCount = 0

// Part1 of Day 18
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	sum := 0

	for _, line := range input {
		lineCount++
		sum += parseExpression(line).reduce()
	}

	return days.Result(fmt.Sprint(sum)), nil
}

// Part2 of Day 18
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	sum := 0

	for _, line := range input {
		lineCount++
		v, _ := strconv.Atoi(processInnerMostParentheses(line))
		sum += v
	}

	return days.Result(fmt.Sprint(sum)), nil
}

func (e constantExp) reduce() int {
	return int(e)
}

func (e sumExp) reduce() int {
	sum := 0

	for _, exp := range e {
		sum += exp.reduce()
	}

	return sum
}

func (e productExp) reduce() int {
	product := 1

	for _, exp := range e {
		product *= exp.reduce()
	}

	return product
}

func parseAdvancedMathExpression(line string) expression {
	chars := append([]rune(strings.ReplaceAll(line, ")", " )")), ' ')

	var currentExp expression = nil
	currentNumber := []rune{}
	productStack := stack.New()

	for _, c := range chars {
		if c == '+' {
			product, ok := currentExp.(productExp)

			if ok {
				productStack.Push(productExp{product[0]})
				currentExp = sumExp{product[1]}
			} else {
				currentExp = sumExp{currentExp}
			}

		} else if c == '*' {
			currentExp = productExp{currentExp}
		} else if c == ' ' {
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(string(currentNumber))

				sum, ok := currentExp.(sumExp)

				if ok {
					sum = append(sum, constantExp(number))

					if productStack.Len() > 0 {
						product := productStack.Pop().(productExp)
						product = append(product, sum)
						currentExp = product
					} else {
						currentExp = sum
					}
				} else {
					product, ok := currentExp.(productExp)

					if ok {
						product = append(product, constantExp(number))
						currentExp = product
					} else {
						// fmt.Printf("Current expression %v is not a compound\n", currentExp)
						currentExp = constantExp(number)
					}
				}

				currentNumber = []rune{}
			}
		} else {
			currentNumber = append(currentNumber, c)
		}

		// fmt.Printf("Current expression %v\n", currentExp)
	}

	return currentExp
}

func processInnerMostParentheses(line string) string {
	for strings.Contains(line, "(") {
		// fmt.Printf("Tackling  %s\n", line)
		start := 1 + strings.LastIndex(line, "(")
		end := start + strings.Index(line[start:], ")")

		// fmt.Printf("Will process %s\n", line[start:end])

		processedInnerParenthesis := fmt.Sprint(parseAdvancedMathExpression(line[start:end]).reduce())

		if line != line[start-1:end+1] {
			line = fmt.Sprintf("%s%s%s", line[:start-1], processedInnerParenthesis, line[end+1:])
		} else {
			line = processedInnerParenthesis
		}
	}

	return fmt.Sprint(parseAdvancedMathExpression(line).reduce())
}

func parseExpression(line string) expression {
	chars := append([]rune(strings.ReplaceAll(line, ")", " )")), ' ')

	var currentExp expression = nil
	currentNumber := []rune{}
	stackedExpressions := stack.New()

	for _, c := range chars {
		if c == '+' {
			currentExp = sumExp{currentExp}
		} else if c == '*' {
			currentExp = productExp{currentExp}
		} else if c == ' ' {
			if len(currentNumber) > 0 {
				number, _ := strconv.Atoi(string(currentNumber))

				sum, ok := currentExp.(sumExp)

				if ok {
					sum = append(sum, constantExp(number))
					currentExp = sum
				} else {
					product, ok := currentExp.(productExp)

					if ok {
						product = append(product, constantExp(number))
						currentExp = product
					} else {
						// fmt.Printf("Current expression %v is not a compound\n", currentExp)
						currentExp = constantExp(number)
					}
				}

				currentNumber = []rune{}
			}
		} else if c == '(' {
			if currentExp == nil {
				currentExp = productExp{constantExp(1)}
			}

			stackedExpressions.Push(currentExp)
			currentExp = nil
		} else if c == ')' {
			poppedExp := stackedExpressions.Pop()

			sum, ok := poppedExp.(sumExp)

			if ok {
				sum = append(sum, currentExp)
				currentExp = sum
			} else {
				product, ok := poppedExp.(productExp)

				if ok {
					product = append(product, currentExp)
					currentExp = product
				} else {
					panic(fmt.Sprintf("What is this fuckery on line %d ??? %v - %v\n", lineCount, poppedExp, currentExp))
				}
			}
		} else {
			currentNumber = append(currentNumber, c)
		}

		// fmt.Printf("Current expression %v\n", currentExp)
	}

	return currentExp
}
