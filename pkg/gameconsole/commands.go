package gameconsole

import (
	"strconv"
	"strings"
)

type command interface {
	Apply(*GameConsole)
	GetOriginalParameters() string
}

type baseCommand struct {
	originalParameters string
}

func (c baseCommand) GetOriginalParameters() string {
	return c.originalParameters
}

type nopCommand struct {
	baseCommand
}

func (nopCommand) Apply(g *GameConsole) {
	g.cursor++
}

type accCommand struct {
	baseCommand
	increment int
}

func (c accCommand) Apply(g *GameConsole) {
	g.accumulator += c.increment
	g.cursor++
}

type jmpCommand struct {
	baseCommand
	directedDistance int
}

func (c jmpCommand) Apply(g *GameConsole) {
	g.cursor += c.directedDistance
}

func parseCommand(commandLine string) (command, error) {
	verb := commandLine[:3]
	parameters := strings.TrimSpace(commandLine[3:])

	switch verb {
	case "acc":
		return createAccCommand(parameters), nil
	case "jmp":
		return createJmpCommand(parameters), nil
	case "nop":
		return createNopCommand(parameters), nil
	default:
		return nil, errorCommandNotParsed
	}
}

func createAccCommand(parameters string) accCommand {
	increment, _ := strconv.Atoi(parameters)
	return accCommand{baseCommand{parameters}, increment}
}

func createJmpCommand(parameters string) jmpCommand {
	increment, _ := strconv.Atoi(parameters)
	return jmpCommand{baseCommand{parameters}, increment}
}

func createNopCommand(parameters string) nopCommand {
	return nopCommand{baseCommand{parameters}}
}
