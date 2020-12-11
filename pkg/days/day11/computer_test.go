package day11

import (
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "26" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "37" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_Small(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"...",
		"###",
		"#.#",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "4" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_SmallBig(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"LLLLL",
		"LLLLL",
		"LLLLL",
		"LLLLL",
		"LLLLL",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "9" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_Small2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"...",
		"L.L",
		"L.L",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "4" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_SmallestFullyOccupiedSample(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"#",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart1_SmallestFullyEmptySample(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"L",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "1" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestInitialState(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})

	res := countOccupiedSeats(seatMap)

	if res != 0 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestFirstRound_Part1(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})

	newMap, didUpdate := executeRound(seatMap, shouldUpdateSeat)

	if !didUpdate {
		t.Fatalf("map unchanged")
	}

	res := countOccupiedSeats(newMap)

	if res != 71 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestFirst2Rounds_Part1(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	})

	seatMap, didUpdate := executeRound(seatMap, shouldUpdateSeat)

	if !didUpdate {
		t.Fatalf("map unchanged")
	}

	seatMap, didUpdate = executeRound(seatMap, shouldUpdateSeat)

	if !didUpdate {
		t.Fatalf("map unchanged")
	}

	res := countOccupiedSeats(seatMap)

	if res != 20 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestCountOccupiedSeats(t *testing.T) {
	res := countOccupiedSeats([][]rune{
		{'L', '.', '#'},
		{'#', '#', '#'},
		{'L', '.', '#'},
	})

	if res != 5 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestCountOccupiedSeatsAround(t *testing.T) {
	seatMap := [][]rune{
		{'L', '.', '#'},
		{'#', '#', '#'},
		{'L', '.', '#'},
	}

	testPositionsWithExpectedResult := [][]int{
		{0, 0, 2},
		{1, 1, 4},
		{2, 0, 2},
		{1, 2, 3},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}

func TestCountOccupiedVisibleSeatsAround(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	})

	testPositionsWithExpectedResult := [][]int{
		{4, 3, 8},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countVisibleOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}

func TestCountOccupiedVisibleSeatsAround2(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		".............",
		".L.L.#.#.#.#.",
		".............",
	})

	testPositionsWithExpectedResult := [][]int{
		{1, 1, 0},
		{1, 3, 1},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countVisibleOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}

func TestCountOccupiedVisibleSeatsAround3(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"...L...",
		"##...##",
		"#.#.#.#",
		".##.##.",
	})

	testPositionsWithExpectedResult := [][]int{
		{3, 3, 0},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countVisibleOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}

func TestCountOccupiedVisibleSeatsAround4(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		"###.##.",
		"#.#.###",
		"##...##",
		"#..L...",
		"##..###",
		"###.#.#",
		".##.##.",
	})

	testPositionsWithExpectedResult := [][]int{
		{3, 3, 5},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countVisibleOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}

func TestCountOccupiedVisibleSeatsAround5(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		".##.##.",
		"#.#.#.#",
		"##...##",
		"#L.L.L#",
		"##...##",
		"#.#.#.#",
		".##.##.",
	})

	testPositionsWithExpectedResult := [][]int{
		{3, 3, 0},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countVisibleOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}

func TestCountOccupiedVisibleSeatsAround7(t *testing.T) {
	seatMap := createSeatMap(days.Input{
		"#####",
		"#####",
		"#####",
		"#####",
		"#####",
	})

	testPositionsWithExpectedResult := [][]int{
		{1, 0, 5},
	}

	for _, test := range testPositionsWithExpectedResult {
		res := countVisibleOccupiedSeatsAround(test[0], test[1], seatMap)

		if res != test[2] {
			t.Fatalf("Wrong result for [%d, %d]: %d != %d", test[0], test[1], res, test[2])
		}
	}
}
