package day8

import (
	"strconv"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

type command struct {
	verb      string
	parameter int
}

type commandPermutation struct {
	c     command
	index int
}

// A Program to interpret
type Program struct {
	accumulator     int
	instructions    []command
	hitInfiniteLoop bool
}

func newProgram(instructions []command) Program {
	return Program{0, instructions, false}
}

func (p *Program) execute() {
	p.hitInfiniteLoop = false
	cursor := 0
	executed := make(map[int]bool)

	for !executed[cursor] && cursor != len(p.instructions) {
		executed[cursor] = true
		command := p.instructions[cursor]

		switch command.verb {
		case "acc":
			p.accumulator += command.parameter
			cursor++
		case "jmp":
			cursor += command.parameter
		case "nop":
			cursor++
		default:
		}
	}

	p.hitInfiniteLoop = cursor != len(p.instructions)
}

func (p Program) createMultiverse() []Program {
	res := []Program{p}

	possiblePermutations := findPossiblePermutations(p)

	for _, permutation := range possiblePermutations {
		res = append(res, createProgramWithPermutation(p, permutation))
	}

	return res
}

func findPossiblePermutations(p Program) []commandPermutation {
	permutations := []commandPermutation{}

	for index, instruction := range p.instructions {
		switch instruction.verb {
		case "jmp":
			permutations = append(permutations, commandPermutation{command{"nop", instruction.parameter}, index})
		case "nop":
			permutations = append(permutations, commandPermutation{command{"jmp", instruction.parameter}, index})
		default:
		}
	}

	return permutations
}

func createProgramWithPermutation(p Program, permutation commandPermutation) Program {
	newInstructions := []command{}

	for index, instruction := range p.instructions {
		if index == permutation.index {
			newInstructions = append(newInstructions, permutation.c)
		} else {
			newInstructions = append(newInstructions, instruction)
		}
	}

	return newProgram(newInstructions)
}

// ParseProgram parses a program from input
func ParseProgram(input days.Input) Program {
	res := Program{}

	for _, line := range input {
		res.instructions = append(res.instructions, parseCommand(line))
	}

	return res
}

func parseCommand(commandLine string) command {
	verb := commandLine[:3]
	parameter, _ := strconv.Atoi(commandLine[4:])
	return command{verb, parameter}
}
