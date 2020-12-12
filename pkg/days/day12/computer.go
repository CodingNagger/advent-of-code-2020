package day12

import (
	"fmt"
	"math"
	"strconv"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 12
type Computer struct {
}

const (
	north   = 'N'
	west    = 'W'
	east    = 'E'
	south   = 'S'
	forward = 'F'
	left    = 'L'
	right   = 'R'
)

type point struct {
	x int
	y int
}

type waypointVector struct {
	n int
	w int
	e int
	s int
}

type boat struct {
	face     rune
	position point
}

type weirdBoat struct {
	position  point
	direction waypointVector
}

// Part1 of Day 12
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	b := boat{}
	b.travel(input)
	return days.Result(fmt.Sprint(manhattanDistance(b.position))), nil
}

// Part2 of Day 12
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	b := weirdBoat{}
	b.travel(input)
	return days.Result(fmt.Sprint(manhattanDistance(b.position))), nil
}

func manhattanDistance(position point) int {
	return int(math.Abs(float64(position.x)) + math.Abs(float64(position.y)))
}

func (b *weirdBoat) travel(input days.Input) {
	b.position = point{0, 0}
	b.direction = waypointVector{n: 1, e: 10, s: 0, w: 0}

	for _, instruction := range input {
		b.move(instruction)
		// fmt.Printf("Move %s leaves boat at %v with waypoint N: %d E: %d S: %d W: %d\n", instruction, b.position, b.direction.n, b.direction.e, b.direction.s, b.direction.w)
	}
}

func (b *weirdBoat) move(instruction string) {
	action := []rune(instruction)[0]
	value, _ := strconv.Atoi(instruction[1:])

	switch action {
	case north:
		fallthrough
	case east:
		fallthrough
	case south:
		fallthrough
	case west:
		b.moveWaypoint(action, value)
	case forward:
		b.moveTowardsWaypoint(value)
	case left:
		b.rotateWaypointLeftBy(value)
	case right:
		b.rotateWaypointRightBy(value)
	}
}

func (b *weirdBoat) moveTowardsWaypoint(distance int) {
	b.moveWithDirection(north, b.direction.n*distance)
	b.moveWithDirection(south, b.direction.s*distance)
	b.moveWithDirection(west, b.direction.w*distance)
	b.moveWithDirection(east, b.direction.e*distance)
}

func (b *weirdBoat) moveWithDirection(direction rune, distance int) {
	switch direction {
	case north:
		b.position.y -= distance
	case east:
		b.position.x += distance
	case south:
		b.position.y += distance
	case west:
		b.position.x -= distance
	}
}

func (b *weirdBoat) rotateWaypointRightBy(degrees int) {
	turnCounts := degrees / 90

	for i := 0; i < turnCounts; i++ {
		b.direction.n, b.direction.e, b.direction.s, b.direction.w =
			b.direction.w, b.direction.n, b.direction.e, b.direction.s
	}
}

func (b *weirdBoat) rotateWaypointLeftBy(degrees int) {
	turnCounts := degrees / 90

	for i := 0; i < turnCounts; i++ {
		b.direction.n, b.direction.e, b.direction.s, b.direction.w =
			b.direction.e, b.direction.s, b.direction.w, b.direction.n
	}
}

func (b *weirdBoat) moveWaypoint(direction rune, distance int) {
	switch direction {
	case north:
		b.direction.n += distance
	case east:
		b.direction.e += distance
	case south:
		b.direction.s += distance
	case west:
		b.direction.w += distance
	}
}

func (b *boat) travel(input days.Input) {
	b.face = east
	b.position = point{0, 0}

	for _, instruction := range input {
		b.move(instruction)
	}
}

func (b *boat) move(instruction string) {
	action := []rune(instruction)[0]
	value, _ := strconv.Atoi(instruction[1:])

	switch action {
	case north:
		fallthrough
	case east:
		fallthrough
	case south:
		fallthrough
	case west:
		b.moveWithDirection(action, value)
	case forward:
		b.moveWithDirection(b.face, value)
	case left:
		b.turnLeftBy(value)
	case right:
		b.turnRightBy(value)
	}
}

func (b *boat) moveWithDirection(direction rune, distance int) {
	switch direction {
	case north:
		b.position.y -= distance
	case east:
		b.position.x += distance
	case south:
		b.position.y += distance
	case west:
		b.position.x -= distance
	}
}

func (b *boat) turnLeftBy(degrees int) {
	turnCounts := degrees / 90

	for i := 0; i < turnCounts; i++ {
		b.face = rotateFaceLeft(b.face)
	}
}

func rotateFaceLeft(face rune) rune {
	switch face {
	case north:
		return west
	case east:
		return north
	case south:
		return east
	case west:
		return south
	}

	return face
}

func (b *boat) turnRightBy(degrees int) {
	turnCounts := degrees / 90

	for i := 0; i < turnCounts; i++ {
		b.face = rotateFaceRight(b.face)
	}
}

func rotateFaceRight(face rune) rune {
	switch face {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}

	return face
}
