package day20

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
	"github.com/codingnagger/advent-of-code-2020/pkg/foundation/inputparser"
)

func getTestInput() days.Input {
	return inputparser.ReadInput("../../../assets/input/day20_test.txt")
}

func getRealInput() days.Input {
	return inputparser.ReadInput("../../../assets/input/day20.txt")
}

func TestPart1(t *testing.T) {
	testDay := NewComputer()

	res, err := testDay.Part1(getTestInput())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "20899048083289" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := NewComputer()

	res, err := testDay.Part2(getTestInput())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "273" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestLoadCameras_TestInput(t *testing.T) {
	testDay := NewComputer()

	testDay.loadCameras(getTestInput())

	res := testDay.cameras

	if len(res) != 9 {
		t.Fatalf("Wrong result: %v", res)
	}
}

func TestRemoveBorders(t *testing.T) {

}

func TestLoadCameras_RealInput(t *testing.T) {
	testDay := NewComputer()

	testDay.loadCameras(getRealInput())

	res := testDay.cameras

	if len(res) != 144 {
		t.Fatalf("Wrong result: %v", res)
	}

	for i, c := range testDay.cameras {
		if len(c.tiles) != 10 {
			t.Fatalf("%d - Wrong line count", i)
		}

		for j, tileColumns := range c.tiles {
			if len(tileColumns) != 10 {
				t.Fatalf("%d - %d- Wrong column count", i, j)
			}
		}
	}
}

func TestLoadSingleCamera(t *testing.T) {
	testDay := NewComputer()

	res := testDay.loadSingleCamera("Tile 2311:\n" +
		"..##.#..#.\n" +
		"##..#.....\n" +
		"#...##..#.\n" +
		"####.#...#\n" +
		"##.##.###.\n" +
		"##...#.###\n" +
		".#.#.#..##\n" +
		"..#....#..\n" +
		"###...#.#.\n" +
		"..###..###")

	if res.idNumber != 2311 {
		t.Fatalf("Wrong id: %v", res)
	}

	if len(res.tiles) != 10 {
		t.Fatalf("Wrong tiles height: %v", res)
	}

	if len(res.tiles[0]) != 10 {
		t.Fatalf("Wrong tiles width: %v", res)
	}

	if res.tiles[0][9] != '.' {
		t.Fatalf("Wrong tiles width: %v", res)
	}
}
