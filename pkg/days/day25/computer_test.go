package day25

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"5764801",
		"17807724",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "14897079" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestFindLoopSize(t *testing.T) {
	if loopSize := findLoopSize(17807724); loopSize != 11 {
		t.Fatalf("Wrong result: %d", loopSize)
	}

	if loopSize := findLoopSize(5764801); loopSize != 8 {
		t.Fatalf("Wrong result: %d", loopSize)
	}
}

func TestTransform(t *testing.T) {
	if res := transform(8, 17807724); res != 14897079 {
		t.Fatalf("Wrong result key transform: %d", res)
	}

	if res := transform(11, 5764801); res != 14897079 {
		t.Fatalf("Wrong result door transform: %d", res)
	}
}
