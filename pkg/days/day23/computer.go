package day23

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 23
type Computer struct {
}

type cupsGame struct {
	cursor    *node
	state     *circle
	maxValues []int64
	nodes     map[int64]*node
}

type circle struct {
	start *node
	end   *node
	count int64
}

type node struct {
	value int64
	prev  *node
	next  *node
}

// Part1 of Day 23
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	g := newCupsGame(input[0])
	g.playNRound(100)
	return days.Result(g.collectCups()), nil
}

// Part2 of Day 23
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	g := newOneMillionCupsGame(input[0])
	g.playNRound(10000000)
	one := g.nodes[1]
	return days.Result(fmt.Sprint(one.next.value * one.next.next.value)), nil
}

func (c *circle) push(val int64) *node {
	n := &node{value: val}

	if c.start == nil {
		c.start, c.end = n, n

		c.count = 1
	} else {
		n.prev = c.end
		n.next = c.start

		c.start.prev = n
		c.end.next = n
		c.end = n

		c.count++
	}

	return n
}

func newCupsGame(inputLine string) *cupsGame {
	digits := strings.Split(inputLine, "")

	circle := &circle{}
	nodes := map[int64]*node{}

	for _, potentialDigit := range digits {
		d, _ := strconv.Atoi(potentialDigit)
		digit := int64(d)
		nodes[digit] = circle.push(digit)
	}

	return &cupsGame{state: circle, cursor: circle.start, maxValues: []int64{9, 8, 7, 6, 5}, nodes: nodes}
}

func newOneMillionCupsGame(inputLine string) *cupsGame {
	g := newCupsGame(inputLine)

	for i := 10; i <= 1000000; i++ {
		digit := int64(i)
		n := g.state.push(digit)
		g.nodes[digit] = n
	}

	fmt.Printf("Node 9 - %d - Node 10 - %d - Node 1000000 - %d\n", g.nodes[9].value, g.nodes[10].value, g.nodes[1000000].value)

	g.maxValues = []int64{1000000, 1000000 - 1, 1000000 - 2, 1000000 - 3, 1000000 - 4}

	return g
}

func (c *cupsGame) collectCups() string {
	collectedCups := ""

	oneNode := c.nodes[1]

	for cursor := oneNode.next; cursor != oneNode; cursor = cursor.next {
		collectedCups += fmt.Sprint(cursor.value)
	}

	return collectedCups
}

func (c *cupsGame) playNRound(n int) {
	for i := 0; i < n; i++ {
		c.playRound()
	}
}

func (c *cupsGame) playRound() {
	search := c.current() - 1
	pickedUp := map[int64]bool{}

	firstPickUp := c.cursor.next
	searchCursor := firstPickUp

	for i := 0; i < 3; i++ {
		pickedUp[searchCursor.value] = true
		searchCursor = searchCursor.next
	}

	lastPickUp := searchCursor.prev

	pickedUp[c.current()] = true

	for i := 0; i < 3; i++ {
		if pickedUp[search] {
			search--
		}
	}

	var destination *node

	if !pickedUp[search] && search > 0 {
		destination = c.nodes[search]
	} else {
		for i := 0; i < 5; i++ {
			potentialMaxValue := c.maxValues[i]

			if !pickedUp[potentialMaxValue] {
				destination = c.nodes[potentialMaxValue]
				break
			}
		}
	}

	c.cursor.next = lastPickUp.next
	c.cursor.next.prev = c.cursor

	firstPickUp.prev = destination
	lastPickUp.next = destination.next
	destination.next.prev = lastPickUp
	destination.next = firstPickUp

	c.cursor = c.cursor.next
}

func (c *cupsGame) current() int64 {
	return c.cursor.value
}
