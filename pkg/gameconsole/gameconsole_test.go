package gameconsole

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestExecute(t *testing.T) {
	g := newGameConsole([]command{
		createNopCommand("0"),
		createAccCommand("1"),
		createJmpCommand("4"),
		createAccCommand("3"),
		createJmpCommand("-3"),
		createAccCommand("-99"),
		createAccCommand("1"),
		createJmpCommand("-4"),
		createAccCommand("6"),
	})

	res, _ := g.Execute()

	if res != 5 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestExecutePermutationsAndFix(t *testing.T) {
	g := newGameConsole([]command{
		createNopCommand("0"),
		createAccCommand("1"),
		createJmpCommand("4"),
		createAccCommand("3"),
		createJmpCommand("-3"),
		createAccCommand("-99"),
		createAccCommand("1"),
		createJmpCommand("-4"),
		createAccCommand("6"),
	})

	res, err := g.ExecutePermutationsAndFix()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != 8 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestParseCommand(t *testing.T) {
	expectedResult := createAccCommand("-99")
	result, err := parseCommand("acc -99")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result: %v", result)
	}
}

func TestLoad(t *testing.T) {
	expectedResult := newGameConsole([]command{
		createNopCommand("+0"),
		createAccCommand("+1"),
		createJmpCommand("+4"),
	})

	result, err := Load(days.Input{
		"nop +0",
		"acc +1",
		"jmp +4",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !reflect.DeepEqual(*result, expectedResult) {
		t.Fatalf("Wrong result: %v should be %v", result, expectedResult)
	}
}
