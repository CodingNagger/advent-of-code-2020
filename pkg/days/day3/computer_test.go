package day3

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func testInput() days.Input {
	return days.Input{
		"..##.........##.........##.........##.........##.........##.......",
		"#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..",
		".#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.",
		"..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#",
		".#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.",
		"..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....",
		".#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#",
		".#........#.#........#.#........#.#........#.#........#.#........#",
		"#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...",
		"#...##....##...##....##...##....##...##....##...##....##...##....#",
		".#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#",
	}
}

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(testInput())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "7" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(testInput())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "336" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestCountTreesForDirection_Bug(t *testing.T) {
	res := countTreesForDirection(testInput(), direction{7, 1})

	if res != 4 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestCountTreesForDirection_TinyDebug(t *testing.T) {
	input := days.Input{
		"..#",
		".#.",
		"..#",
	}
	res := countTreesForDirection(input, direction{1, 1})

	if res != 2 {
		t.Fatalf("Wrong result for 1, 1 : %d", res)
	}

	res = countTreesForDirection(input, direction{3, 1})

	if res != 0 {
		t.Fatalf("Wrong result for 1, 1 : %d", res)
	}

	res = countTreesForDirection(input, direction{5, 1})

	if res != 0 {
		t.Fatalf("Wrong result for 5, 1 : %d", res)
	}

	res = countTreesForDirection(input, direction{7, 1})

	if res != 2 {
		t.Fatalf("Wrong result for 7, 1 : %d", res)
	}

	res = countTreesForDirection(input, direction{1, 2})

	if res != 0 {
		t.Fatalf("Wrong result for 1, 2 : %d", res)
	}
}
