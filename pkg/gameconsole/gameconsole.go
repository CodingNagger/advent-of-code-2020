package gameconsole

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

var (
	errorInfiniteLoop     = fmt.Errorf("Hit an infinite loop")
	errorCommandNotParsed = fmt.Errorf("Could not parse command")
)

type commandReplacement struct {
	c     command
	index int
}

// A GameConsole to interpret
type GameConsole struct {
	accumulator    int
	instructions   []command
	swappedCommand *commandReplacement
	cursor         int
}

func newGameConsole(instructions []command) GameConsole {
	return GameConsole{0, instructions, nil, 0}
}

// Execute runs the GameConsole instructions
func (g *GameConsole) Execute() (int, error) {
	g.cursor = 0
	g.accumulator = 0
	executed := make(map[int]bool)

	for !executed[g.cursor] && g.cursor != len(g.instructions) {
		executed[g.cursor] = true

		if g.swappedCommand != nil && g.cursor == g.swappedCommand.index {
			g.swappedCommand.c.Apply(g)
		} else {
			g.instructions[g.cursor].Apply(g)
		}
	}

	if g.cursor != len(g.instructions) {
		return g.accumulator, errorInfiniteLoop
	}

	return g.accumulator, nil
}

// ExecutePermutationsAndFix runs program permutations until one can terminate naturally
func (g GameConsole) ExecutePermutationsAndFix() (int, error) {
	res, err := g.Execute()

	if err == nil {
		return res, nil
	}

	possiblePermutations := findPossiblePermutations(g)

	for _, permutation := range possiblePermutations {
		g.swappedCommand = &permutation

		res, err := g.Execute()

		if err == nil {
			return res, nil
		}
	}

	return -1, fmt.Errorf("No answer found")
}

func findPossiblePermutations(g GameConsole) []commandReplacement {
	permutations := []commandReplacement{}

	for index, instruction := range g.instructions {
		switch instruction.(type) {
		case jmpCommand:
			permutations = append(permutations, commandReplacement{createNopCommand(instruction.GetOriginalParameters()), index})
		case nopCommand:
			permutations = append(permutations, commandReplacement{createJmpCommand(instruction.GetOriginalParameters()), index})
		default:
		}
	}

	return permutations
}

// Load parses and laods a gameconsole from input
func Load(input days.Input) (*GameConsole, error) {
	res := GameConsole{}

	for _, line := range input {
		command, err := parseCommand(line)

		if err != nil {
			return nil, err
		}

		res.instructions = append(res.instructions, command)
	}

	return &res, nil
}
