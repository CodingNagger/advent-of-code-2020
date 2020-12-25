package day24

import (
	"reflect"
	"testing"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

func TestPart1(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part1(days.Input{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "10" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestPart2(t *testing.T) {
	testDay := &Computer{}

	res, err := testDay.Part2(days.Input{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != "2208" {
		t.Fatalf("Wrong result: %s", res)
	}
}

func TestProcessTileInstructionsToRoot(t *testing.T) {
	expectedResult := coordinates{x: 0, y: 0, z: 0}

	result := processTileInstructions("")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result %v", result)
	}
}

func TestProcessTileInstructionsToRoot_LongWay(t *testing.T) {
	expectedResult := coordinates{x: 0, y: 0, z: 0}

	result := processTileInstructions("nwwsweewewewewe")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result %v", result)
	}
}

func TestProcessTileInstructionsTo_IndirectSouthEast(t *testing.T) {
	expectedResult := coordinates{x: 0, y: 1, z: -1}

	result := processTileInstructions("esew")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf("Wrong result %v", result)
	}
}

func TestParseDirections(t *testing.T) {
	res := parseDirections("esew")

	if len(res) != 3 {
		t.Fatalf("Wrong size: %d - %v", len(res), res)
	}

	if res[0] != east || res[1] != southEast || res[2] != west {
		t.Fatalf("Wrong directions: %v", res)
	}
}
