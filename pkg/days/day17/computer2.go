package day17

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

type hypercube struct {
	x, y, z, w int
}

type pocketHyperDimension struct {
	hypercubes map[hypercube]bool
}

// Part2 of Day 17
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	p := createPocketHyperDimension(input)
	return days.Result(fmt.Sprint(p.runCyclesAndReturnActiveCount(6))), nil
}

func createPocketHyperDimension(input days.Input) pocketHyperDimension {
	p := pocketHyperDimension{}
	p.hypercubes = map[hypercube]bool{}

	for y := 0; y < len(input); y++ {
		cols := []rune(input[y])
		for x := 0; x < len(cols); x++ {
			c := hypercube{x: x, y: y, z: 0, w: 0}
			p.hypercubes[c] = cols[x] == '#'
		}
	}

	return p
}

func (p *pocketHyperDimension) runCyclesAndReturnActiveCount(count int) int {
	for i := 0; i < count; i++ {
		p.cycle()
	}

	return p.countActivehypercubes()
}

func (p *pocketHyperDimension) cycle() {
	nextState := map[hypercube]bool{}

	p.addInvisibleNeighboursToPlane()

	for hypercube, isActive := range p.hypercubes {
		activeNeighbours := p.countActiveNeighbours(hypercube)
		nextState[hypercube] = activeNeighbours == 3 || (isActive && activeNeighbours == 2)
		// fmt.Printf("%v from state %v to %v\n", p.hypercubes[hypercube], hypercube, nextState[hypercube])
	}

	p.hypercubes = nextState
}

func (p *pocketHyperDimension) addInvisibleNeighboursToPlane() {
	hypercubesToAdd := []hypercube{}

	for hypercube := range p.hypercubes {
		for _, neighbour := range hypercube.getNeighbours() {
			_, ok := p.hypercubes[neighbour]

			if !ok {
				hypercubesToAdd = append(hypercubesToAdd, neighbour)
			}
		}
	}

	for _, c := range hypercubesToAdd {
		p.hypercubes[c] = false
	}
}

func (p *pocketHyperDimension) countActiveNeighbours(c hypercube) int {
	count := 0
	for _, n := range c.getNeighbours() {
		if p.hypercubes[n] {
			count++
		}
	}
	return count
}

func (p *pocketHyperDimension) countActivehypercubes() int {
	count := 0
	for _, isActive := range p.hypercubes {
		if isActive {
			count++
		}
	}
	return count
}

func (c *hypercube) isNeighbour(candidate hypercube) bool {
	for _, v := range c.getNeighbourVectors() {
		if c.x+v.x == candidate.x && c.y+v.y == candidate.y && c.z+v.z == candidate.z {
			return true
		}
	}

	return false
}

func (c hypercube) getNeighbours() []hypercube {
	res := []hypercube{}

	for _, v := range c.getNeighbourVectors() {
		res = append(res, hypercube{c.x + v.x, c.y + v.y, c.z + v.z, c.w + v.w})
	}

	return res
}

func (c *hypercube) getNeighbourVectors() []hypercube {
	return []hypercube{
		hypercube{x: -1, y: -1, z: -1, w: -1}, hypercube{x: 0, y: -1, z: -1, w: -1}, hypercube{x: 1, y: -1, z: -1, w: -1},
		hypercube{x: -1, y: 0, z: -1, w: -1}, hypercube{x: 0, y: 0, z: -1, w: -1}, hypercube{x: 1, y: 0, z: -1, w: -1},
		hypercube{x: -1, y: 1, z: -1, w: -1}, hypercube{x: 0, y: 1, z: -1, w: -1}, hypercube{x: 1, y: 1, z: -1, w: -1},
		hypercube{x: -1, y: -1, z: 0, w: -1}, hypercube{x: 0, y: -1, z: 0, w: -1}, hypercube{x: 1, y: -1, z: 0, w: -1},
		hypercube{x: -1, y: 0, z: 0, w: -1}, hypercube{x: 0, y: 0, z: 0, w: -1}, hypercube{x: 1, y: 0, z: 0, w: -1},
		hypercube{x: -1, y: 1, z: 0, w: -1}, hypercube{x: 0, y: 1, z: 0, w: -1}, hypercube{x: 1, y: 1, z: 0, w: -1},
		hypercube{x: -1, y: -1, z: 1, w: -1}, hypercube{x: 0, y: -1, z: 1, w: -1}, hypercube{x: 1, y: -1, z: 1, w: -1},
		hypercube{x: -1, y: 0, z: 1, w: -1}, hypercube{x: 0, y: 0, z: 1, w: -1}, hypercube{x: 1, y: 0, z: 1, w: -1},
		hypercube{x: -1, y: 1, z: 1, w: -1}, hypercube{x: 0, y: 1, z: 1, w: -1}, hypercube{x: 1, y: 1, z: 1, w: -1},

		hypercube{x: -1, y: -1, z: -1, w: 0}, hypercube{x: 0, y: -1, z: -1, w: 0}, hypercube{x: 1, y: -1, z: -1, w: 0},
		hypercube{x: -1, y: 0, z: -1, w: 0}, hypercube{x: 0, y: 0, z: -1, w: 0}, hypercube{x: 1, y: 0, z: -1, w: 0},
		hypercube{x: -1, y: 1, z: -1, w: 0}, hypercube{x: 0, y: 1, z: -1, w: 0}, hypercube{x: 1, y: 1, z: -1, w: 0},
		hypercube{x: -1, y: -1, z: 0, w: 0}, hypercube{x: 0, y: -1, z: 0, w: 0}, hypercube{x: 1, y: -1, z: 0, w: 0},
		hypercube{x: -1, y: 0, z: 0, w: 0} /*, hypercube{x: 0, y: 0, z: 0, w: 0} */, hypercube{x: 1, y: 0, z: 0, w: 0},
		hypercube{x: -1, y: 1, z: 0, w: 0}, hypercube{x: 0, y: 1, z: 0, w: 0}, hypercube{x: 1, y: 1, z: 0, w: 0},
		hypercube{x: -1, y: -1, z: 1, w: 0}, hypercube{x: 0, y: -1, z: 1, w: 0}, hypercube{x: 1, y: -1, z: 1, w: 0},
		hypercube{x: -1, y: 0, z: 1, w: 0}, hypercube{x: 0, y: 0, z: 1, w: 0}, hypercube{x: 1, y: 0, z: 1, w: 0},
		hypercube{x: -1, y: 1, z: 1, w: 0}, hypercube{x: 0, y: 1, z: 1, w: 0}, hypercube{x: 1, y: 1, z: 1, w: 0},

		hypercube{x: -1, y: -1, z: -1, w: 1}, hypercube{x: 0, y: -1, z: -1, w: 1}, hypercube{x: 1, y: -1, z: -1, w: 1},
		hypercube{x: -1, y: 0, z: -1, w: 1}, hypercube{x: 0, y: 0, z: -1, w: 1}, hypercube{x: 1, y: 0, z: -1, w: 1},
		hypercube{x: -1, y: 1, z: -1, w: 1}, hypercube{x: 0, y: 1, z: -1, w: 1}, hypercube{x: 1, y: 1, z: -1, w: 1},
		hypercube{x: -1, y: -1, z: 0, w: 1}, hypercube{x: 0, y: -1, z: 0, w: 1}, hypercube{x: 1, y: -1, z: 0, w: 1},
		hypercube{x: -1, y: 0, z: 0, w: 1}, hypercube{x: 0, y: 0, z: 0, w: 1}, hypercube{x: 1, y: 0, z: 0, w: 1},
		hypercube{x: -1, y: 1, z: 0, w: 1}, hypercube{x: 0, y: 1, z: 0, w: 1}, hypercube{x: 1, y: 1, z: 0, w: 1},
		hypercube{x: -1, y: -1, z: 1, w: 1}, hypercube{x: 0, y: -1, z: 1, w: 1}, hypercube{x: 1, y: -1, z: 1, w: 1},
		hypercube{x: -1, y: 0, z: 1, w: 1}, hypercube{x: 0, y: 0, z: 1, w: 1}, hypercube{x: 1, y: 0, z: 1, w: 1},
		hypercube{x: -1, y: 1, z: 1, w: 1}, hypercube{x: 0, y: 1, z: 1, w: 1}, hypercube{x: 1, y: 1, z: 1, w: 1},
	}
}
