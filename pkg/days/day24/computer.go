package day24

import (
	"fmt"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 24
type Computer struct {
	tiles map[coordinates]colour
}

type direction int
type colour int

type coordinates struct {
	x, y, z int
}

const (
	northEast direction = iota
	east
	southEast
	southWest
	west
	northWest

	black colour = iota
	white
)

// Part1 of Day 23
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	d.loadTiles(input)
	return days.Result(fmt.Sprint(d.countBlackTiles())), nil
}

// Part2 of Day 23
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	d.loadTiles(input)
	d.runExhibit()
	return days.Result(fmt.Sprint(d.countBlackTiles())), nil
}

func (d *Computer) countBlackTiles() int {
	count := 0

	for _, position := range d.tiles {
		if position == black {
			count++
		}
	}

	return count
}

func (d *Computer) runExhibit() {
	for i := 0; i < 100; i++ {
		d.updateTilesArt()
	}
}

func (d *Computer) updateTilesArt() {
	nextState := map[coordinates]colour{}

	d.addInvisibleNeighbours()

	for tile := range d.tiles {
		nextState[tile] = d.nextTileColour(tile)
	}

	d.tiles = nextState
}

func (d *Computer) addInvisibleNeighbours() {
	positions := []coordinates{}

	for tile := range d.tiles {
		for _, neighbour := range tile.neighbours() {
			_, ok := d.tiles[neighbour]

			if !ok {
				positions = append(positions, neighbour)
			}
		}
	}

	for _, p := range positions {
		d.tiles[p] = white
	}
}

func (d *Computer) nextTileColour(position coordinates) colour {
	adjacentBlacks := d.countBlackAdjacentTiles(position)

	if d.checkBlackness(position) {
		if adjacentBlacks == 0 || adjacentBlacks > 2 {
			return white
		}

		return black
	}

	if adjacentBlacks == 2 {
		return black
	}

	return white
}

func (d *Computer) countBlackAdjacentTiles(position coordinates) int {
	count := 0

	for _, neighbour := range position.neighbours() {
		if d.checkBlackness(neighbour) {
			count++
		}
	}

	return count
}

func (c *coordinates) neighbours() []coordinates {
	neighbours := []coordinates{}

	for _, i := range allDirections() {
		neighbours = append(neighbours, c.derive(vectorForDirection(i)))
	}

	return neighbours
}

func (c coordinates) derive(vector coordinates) coordinates {
	return coordinates{x: c.x + vector.x, y: c.y + vector.y, z: c.z + vector.z}
}

func (d *Computer) checkBlackness(position coordinates) bool {
	existing, ok := d.tiles[position]
	return ok && existing == black
}

func (d *Computer) loadTiles(input days.Input) {
	d.tiles = map[coordinates]colour{}

	for _, instructions := range input {
		t := processTileInstructions(instructions)

		if d.checkBlackness(t) {
			d.tiles[t] = white
		} else {
			d.tiles[t] = black
		}
	}
}

func processTileInstructions(instructions string) coordinates {
	directions := parseDirections(instructions)
	cursor := coordinates{x: 0, y: 0, z: 0}

	for _, d := range directions {
		cursor = cursor.derive(vectorForDirection(d))
	}

	return cursor
}

func vectorForDirection(d direction) coordinates {
	switch d {
	case northEast:
		return coordinates{x: 1, y: 0, z: -1}
	case east:
		return coordinates{x: 1, y: -1, z: 0}
	case southEast:
		return coordinates{x: 0, y: -1, z: 1}
	case southWest:
		return coordinates{x: -1, y: 0, z: 1}
	case west:
		return coordinates{x: -1, y: 1, z: 0}
	case northWest:
		return coordinates{x: 0, y: 1, z: -1}
	}

	panic("No other direction available")
}

func allDirections() []direction {
	return []direction{northEast, east, southEast, southWest, west, northWest}
}

func parseDirections(instructions string) []direction {
	elements := []rune(instructions)
	directions := []direction{}

	for i := 0; i < len(elements); i++ {
		switch elements[i] {
		case 'e':
			directions = append(directions, east)
		case 'w':
			directions = append(directions, west)
		case 's':
			switch elements[i+1] {
			case 'e':
				directions = append(directions, southEast)
			case 'w':
				directions = append(directions, southWest)
			}
			i++
		case 'n':
			switch elements[i+1] {
			case 'e':
				directions = append(directions, northEast)
			case 'w':
				directions = append(directions, northWest)
			}
			i++
		}
	}
	return directions
}
