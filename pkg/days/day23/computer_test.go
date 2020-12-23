package day23

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"389125467",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "67384529" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"389125467",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "149245887792" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestNewCupsGame(t *testing.T) {
	g := newCupsGame("389125467")

	cups := g.collectCups()
	if g.collectCups() != "25467389" {
		t.Fatalf("Wrong state: %v", cups)
	}

	if g.current() != 3 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.current())
	}
}

func TestPlayRound(t *testing.T) {
	g := newCupsGame("389125467")
	g.playRound()

	cups := g.collectCups()
	if g.collectCups() != "54673289" {
		t.Fatalf("Wrong state: %v", cups)
	}

	if g.current() != 2 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.current())
	}
}

func TestPlayTwoRounds(t *testing.T) {
	g := newCupsGame("389125467")
	g.playNRound(2)

	cups := g.collectCups()
	if g.collectCups() != "32546789" {
		t.Fatalf("Wrong state: %v", cups)
	}

	if g.current() != 5 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.cursor)
	}
}

func TestPlayThreeRounds(t *testing.T) {
	g := newCupsGame("389125467")
	g.playNRound(3)

	cups := g.collectCups()
	if g.collectCups() != "34672589" {
		t.Fatalf("Wrong state: %v", cups)
	}

	if g.current() != 8 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.cursor)
	}
}

func TestPlaySixthRound(t *testing.T) {
	g := newCupsGame("389125467")
	g.playNRound(6)

	cups := g.collectCups()
	if g.collectCups() != "93672584" {
		t.Fatalf("Wrong state: %v", cups)
	}

	if g.current() != 9 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.cursor)
	}
}

func TestPlayTenRounds(t *testing.T) {
	g := newCupsGame("389125467")
	g.playNRound(10)

	if g.current() != 8 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.cursor)
	}

	if g.collectCups() != "92658374" {
		t.Fatalf("Wrong cup collection: %v", g.collectCups())
	}
}

func TestPlayElevenRounds(t *testing.T) {
	g := newCupsGame("389125467")
	g.playNRound(11)

	cups := g.collectCups()
	if cups != "92637458" {
		t.Fatalf("Wrong state: %d - %s", g.current(), cups)
	}

	if g.current() != 1 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.cursor)
	}
}

func TestPlayTwelveRounds(t *testing.T) {
	g := newCupsGame("389125467")
	g.playNRound(12)

	cups := g.collectCups()
	if cups != "37458926" {
		t.Fatalf("Wrong state: %d - %s", g.current(), cups)
	}

	if g.current() != 3 {
		t.Fatalf("Wrong cursor not positioned correctly: %v", g.current())
	}
}
