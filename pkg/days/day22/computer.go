package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

type winnerFinder func(*combat) (int, int, int)

const (
	P1 = 0
	P2 = 1
)

// Computer of the Advent of code 2020 Day 22
type Computer struct {
	players []*deck
}

type combat struct {
	players         []*deck
	determineWinner winnerFinder
	previousRounds  map[string]bool
}

type deck struct {
	cards []int
}

// Part1 of Day 22
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	d.loadGame(input)

	return days.Result(fmt.Sprint(d.playCombat())), nil
}

// Part2 of Day 22
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	d.loadGame(input)

	return days.Result(fmt.Sprint(d.playRecursiveCombat())), nil
}

func (d *Computer) playRecursiveCombat() int {
	c := combat{players: d.players, determineWinner: determineWinnerRecursively}

	return d.players[c.play()].score()
}

func determineWinnerRecursively(game *combat) (int, int, int) {

	p1Play := game.players[P1].popCard()
	p2Play := game.players[P2].popCard()

	if p1Play <= len(game.players[P1].cards) && p2Play <= len(game.players[P2].cards) {
		subGame := combat{players: []*deck{
			newDeckFrom(game.players[P1].cards[:p1Play]),
			newDeckFrom(game.players[P2].cards[:p2Play]),
		}, determineWinner: determineWinnerRecursively}

		subGame.play()

		if P1 == subGame.findWinner() {
			return P1, p1Play, p2Play
		}
	} else if p1Play > p2Play {
		return P1, p1Play, p2Play
	}

	return P2, p2Play, p1Play
}

func roundState(players []*deck) string {
	return fmt.Sprintf("%s|%s", players[P1].state(), players[P2].state())
}

func (d *deck) state() string {
	res := []string{}

	for _, c := range d.cards {
		res = append(res, fmt.Sprint(c))
	}

	return strings.Join(res, ";")
}

func (d *Computer) playCombat() int {
	c := combat{players: d.players, determineWinner: determineWinnerRegularly}
	c.play()
	return c.calculateWinnerScore()
}

func (d *Computer) loadGame(input days.Input) {
	cardDetectorRegex := regexp.MustCompile(`^\d+$`)
	d.players = make([]*deck, 2)
	d.players[P1] = newDeck()
	d.players[P2] = newDeck()
	currentPlayer := P1

	for _, line := range input {
		if "Player 2:" == line {
			currentPlayer = P2
		} else if cardDetectorRegex.MatchString(line) {
			card, _ := strconv.Atoi(line)
			d.players[currentPlayer].pushCard(card)
		}
	}
}

func (c *combat) play() int {
	c.previousRounds = map[string]bool{}

	for c.players[P1].hasCards() && c.players[P2].hasCards() {
		newRound := roundState(c.players)

		if _, ok := c.previousRounds[newRound]; ok {
			return P1
		}

		c.previousRounds[newRound] = true

		c.playRound()
	}

	return c.findWinner()
}

func (c *combat) playRound() {
	roundWinner, firstCard, secondCard := c.determineWinner(c)

	c.players[roundWinner].pushCard(firstCard)
	c.players[roundWinner].pushCard(secondCard)
}

func determineWinnerRegularly(game *combat) (int, int, int) {
	p1Play := game.players[P1].popCard()
	p2Play := game.players[P2].popCard()

	if p1Play > p2Play {
		return P1, p1Play, p2Play
	}

	return P2, p2Play, p1Play
}

func (c *combat) findWinner() int {
	if c.players[P1].hasCards() {
		return P1
	}
	return P2
}

func (c *combat) calculateWinnerScore() int {
	return c.players[c.findWinner()].score()
}

func newDeck() *deck {
	return &deck{[]int{}}
}

func newDeckFrom(cards []int) *deck {
	copy := newDeck()

	for _, card := range cards {
		copy.pushCard(card)
	}

	return copy
}

func (d *deck) pushCard(card int) {
	d.cards = append(d.cards, card)
}

func (d *deck) popCard() int {
	topCard := d.cards[0]
	d.cards = d.cards[1:]
	return topCard
}

func (d *deck) hasCards() bool {
	return len(d.cards) > 0
}

func (d *deck) score() int {
	res := 0
	highestCard := len(d.cards)

	for _, card := range d.cards {
		res += card * highestCard
		highestCard--
	}

	return res
}
