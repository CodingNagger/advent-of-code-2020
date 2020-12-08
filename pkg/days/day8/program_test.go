package day8

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestParseCommand(t *testing.T) {
	expectedResult := command{"acc", -99}
	result := parseCommand("acc -99")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestParseProgram(t *testing.T) {
	expectedResult := newProgram([]command{
		command{"nop", 0},
		command{"acc", 1},
		command{"jmp", 4},
	})
	result := ParseProgram(days.Input{
		"nop +0",
		"acc +1",
		"jmp +4",
	})

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestCreateProgramMultiverse(t *testing.T) {
	originalProgram := newProgram([]command{
		command{"nop", 0},
		command{"acc", 1},
		command{"jmp", 4},
	})

	expectedResult := []Program{
		originalProgram,
		newProgram([]command{
			command{"jmp", 0},
			command{"acc", 1},
			command{"jmp", 4},
		}),
		newProgram([]command{
			command{"nop", 0},
			command{"acc", 1},
			command{"nop", 4},
		}),
	}

	result := originalProgram.createMultiverse()

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}
