package day17

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		".#.",
		"..#",
		"###",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "112" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		".#.",
		"..#",
		"###",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "848" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestCreatePocketDimension(t *testing.T) {
	input := days.Input{
		".#.",
		"..#",
		"###",
	}

	d := createPocketDimension(input)

	activeCubesCount := d.countActiveCubes()

	if activeCubesCount != 5 {
		t.Fatalf("Wrong active count: %d", activeCubesCount)
	}
}

func TestCycle(t *testing.T) {
	input := days.Input{
		".#.",
		"..#",
		"###",
	}

	d := createPocketDimension(input)
	d.cycle()

	activeCubesCount := d.countActiveCubes()

	if activeCubesCount != 11 {
		t.Fatalf("Wrong active count: %d", activeCubesCount)
	}
}

func TestGetNeighbours(t *testing.T) {
	c := cube{1, 1, 0}

	neighbours := c.getNeighbours()

	if len(neighbours) != 26 {
		t.Fatalf("Wrong neighbours count: %d", len(neighbours))
	}

	dedup := map[cube]bool{}

	for _, n := range neighbours {
		dedup[n] = true
	}

	if len(dedup) != 26 {
		t.Fatalf("Wrong dedup count: %d", len(dedup))
	}
}

func TestGetNeighbourVectors(t *testing.T) {
	c := cube{0, 0, 0}

	vectors := c.getNeighbourVectors()

	if len(vectors) != 26 {
		t.Fatalf("Wrong vectors count: %d", len(vectors))
	}

	dedup := map[cube]bool{}

	for _, n := range vectors {
		dedup[n] = true
	}

	if len(dedup) != 26 {
		t.Fatalf("Wrong dedup count: %d", len(dedup))
	}
}

func TestCountActiveNeighbours(t *testing.T) {
	tests := [][]int{
		// x, y, z, expectedResult
		[]int{0, 0, 0, 1},
		[]int{1, 1, 0, 5},
		[]int{2, 2, 0, 2},
	}

	input := days.Input{
		".#.",
		"..#",
		"###",
	}

	d := createPocketDimension(input)

	for _, test := range tests {
		c := cube{x: test[0], y: test[1], z: test[2]}
		activeNeighbours := d.countActiveNeighbours(c)

		if activeNeighbours != test[3] {
			t.Fatalf("Wrong neighbour count: %d for %v", activeNeighbours, c)
		}
	}
}

func TestIsNeighbour_ValidNeighbours(t *testing.T) {
	c := cube{1, 2, 3}
	validNeighbours := []cube{
		cube{2, 2, 2},
		cube{0, 2, 3},
	}

	for _, n := range validNeighbours {
		if !c.isNeighbour(n) {
			t.Fatalf("Wrong result: %v is a neighbour", n)
		}
	}
}
