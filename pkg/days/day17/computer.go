package day17

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 17
type Computer struct {
}

type cube struct {
	x, y, z int
}

type pocketDimension struct {
	cubes map[cube]bool
}

// Part1 of Day 17
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	p := createPocketDimension(input)
	return days.Result(fmt.Sprint(p.runCyclesAndReturnActiveCount(6))), nil
}

func createPocketDimension(input days.Input) pocketDimension {
	p := pocketDimension{}
	p.cubes = map[cube]bool{}

	for y := 0; y < len(input); y++ {
		cols := []rune(input[y])
		for x := 0; x < len(cols); x++ {
			c := cube{x: x, y: y, z: 0}
			p.cubes[c] = cols[x] == '#'
		}
	}

	return p
}

func (p *pocketDimension) runCyclesAndReturnActiveCount(count int) int {
	for i := 0; i < count; i++ {
		p.cycle()
	}

	return p.countActiveCubes()
}

func (p *pocketDimension) cycle() {
	nextState := map[cube]bool{}

	p.addInvisibleNeighboursToPlane()

	for cube, isActive := range p.cubes {
		activeNeighbours := p.countActiveNeighbours(cube)
		nextState[cube] = activeNeighbours == 3 || (isActive && activeNeighbours == 2)
		fmt.Printf("%v from state %v to %v\n", p.cubes[cube], cube, nextState[cube])
	}

	p.cubes = nextState
}

func (p *pocketDimension) addInvisibleNeighboursToPlane() {
	cubesToAdd := []cube{}

	for cube := range p.cubes {
		for _, neighbour := range cube.getNeighbours() {
			_, ok := p.cubes[neighbour]

			if !ok {
				cubesToAdd = append(cubesToAdd, neighbour)
			}
		}
	}

	for _, c := range cubesToAdd {
		p.cubes[c] = false
	}
}

func (p *pocketDimension) countActiveNeighbours(c cube) int {
	count := 0
	for _, n := range c.getNeighbours() {
		if p.cubes[n] {
			count++
		}
	}
	return count
}

func (p *pocketDimension) countActiveCubes() int {
	count := 0
	for _, isActive := range p.cubes {
		if isActive {
			count++
		}
	}
	return count
}

func (c *cube) isNeighbour(candidate cube) bool {
	for _, v := range c.getNeighbourVectors() {
		if c.x+v.x == candidate.x && c.y+v.y == candidate.y && c.z+v.z == candidate.z {
			return true
		}
	}

	return false
}

func (c cube) getNeighbours() []cube {
	res := []cube{}

	for _, v := range c.getNeighbourVectors() {
		res = append(res, cube{c.x + v.x, c.y + v.y, c.z + v.z})
	}

	return res
}

func (c *cube) getNeighbourVectors() []cube {
	return []cube{
		cube{x: -1, y: -1, z: -1}, cube{x: 0, y: -1, z: -1}, cube{x: 1, y: -1, z: -1},
		cube{x: -1, y: 0, z: -1}, cube{x: 0, y: 0, z: -1}, cube{x: 1, y: 0, z: -1},
		cube{x: -1, y: 1, z: -1}, cube{x: 0, y: 1, z: -1}, cube{x: 1, y: 1, z: -1},
		cube{x: -1, y: -1, z: 0}, cube{x: 0, y: -1, z: 0}, cube{x: 1, y: -1, z: 0},
		cube{x: -1, y: 0, z: 0} /*, cube{x: 0, y: 0, z: 0} */, cube{x: 1, y: 0, z: 0},
		cube{x: -1, y: 1, z: 0}, cube{x: 0, y: 1, z: 0}, cube{x: 1, y: 1, z: 0},
		cube{x: -1, y: -1, z: 1}, cube{x: 0, y: -1, z: 1}, cube{x: 1, y: -1, z: 1},
		cube{x: -1, y: 0, z: 1}, cube{x: 0, y: 0, z: 1}, cube{x: 1, y: 0, z: 1},
		cube{x: -1, y: 1, z: 1}, cube{x: 0, y: 1, z: 1}, cube{x: 1, y: 1, z: 1},
	}
}
