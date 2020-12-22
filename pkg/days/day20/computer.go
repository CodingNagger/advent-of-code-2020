package day20

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/codingnagger/advent-of-code-2020/pkg/days"
)

// Computer of the Advent of code 2020 Day 20
type Computer struct {
	tileIDMatcher   *regexp.Regexp
	cameras         []*camera
	sideLength      int
	arrangedCameras [][]*camera
	image           [][]rune
	monsterRunes    [][]rune
	monsterPresence [][]bool
	imageSize       int
	allEdges        map[string]int
	adjustedCameras map[int]bool
}

type location struct {
	x, y int
}

type camera struct {
	tiles    [][]rune
	idNumber int
}

// NewComputer desc
func NewComputer() Computer {
	return Computer{tileIDMatcher: regexp.MustCompile(`\d+`), adjustedCameras: map[int]bool{}}
}

// Part1 of Day 20
func (d *Computer) Part1(input days.Input) (days.Result, error) {
	d.loadCameras(input)
	d.arrangeCameras()

	return days.Result(fmt.Sprint(d.multiplyCornerCameraIds())), nil
}

// Part2 of Day 20
func (d *Computer) Part2(input days.Input) (days.Result, error) {
	d.Part1(input)
	d.createImage()

	if !d.transformImageToSeeSeaMonsters() {
		panic("Should have some monsters here")
	}

	return days.Result(fmt.Sprint(d.calculateWaterRoughness())), nil
}

func (d *Computer) calculateWaterRoughness() int {
	count := 0

	for j := 0; j < d.imageSize; j++ {
		for i := 0; i < d.imageSize; i++ {
			if d.image[j][i] == '#' && !d.monsterPresence[j][i] {
				count++
			}
		}
	}

	return count
}

func (d *Computer) arrangeCameras() {
	d.initialiseArrangeCameras()

	nextTopLeftCamera := d.fitCamera(d.findCamerasAtEdges()[0], "", "")

	row, col := 0, 0
	topEdge, leftEdge := "", nextTopLeftCamera.rightEdge()
	d.arrangedCameras[row][col] = nextTopLeftCamera
	d.adjustedCameras[nextTopLeftCamera.idNumber] = true
	col++

	for len(d.adjustedCameras) < len(d.cameras) {
		for _, cursor := range d.cameras {
			if _, ok := d.adjustedCameras[cursor.idNumber]; ok {
				continue
			}

			if candidate := d.fitCamera(cursor, leftEdge, topEdge); candidate != nil {
				d.adjustedCameras[candidate.idNumber] = true
				d.arrangedCameras[row][col] = candidate

				leftEdge = candidate.rightEdge()

				if col == d.sideLength-1 {
					col = 0
					row++
					topEdge, leftEdge = nextTopLeftCamera.bottomEdge(), ""

					nextTopLeftCamera = nil
					break
				} else {
					col++

					if row > 0 && d.arrangedCameras[row-1][col] != nil {
						topEdge = d.arrangedCameras[row-1][col].bottomEdge()
					}
				}

				if nextTopLeftCamera == nil {
					nextTopLeftCamera = candidate
				}
			}
		}
	}
}

func (d *Computer) transformIntoTopLeftCorner(wishedTopLeft *camera) *camera {
	for _, c := range wishedTopLeft.getRotationsOfCamera() {
		if d.allEdges[c.topEdge()] == 1 && d.allEdges[c.leftEdge()] == 1 && d.allEdges[c.rightEdge()] == 2 && d.allEdges[c.bottomEdge()] == 2 {
			d.adjustedCameras[c.idNumber] = true
			return c
		}
	}

	return nil
}

func (d *Computer) findNextNeighbourWith(leftEdge, topEdge string) (*camera, bool) {
	for _, candidate := range d.cameras {
		if d.adjustedCameras[candidate.idNumber] {
			continue
		}

		for _, c := range candidate.getCombinationsOfCamera() {
			leftMatch := (leftEdge == "" && d.allEdges[c.leftEdge()] == 1) || leftEdge == c.leftEdge()
			topMatch := (topEdge == "" && d.allEdges[c.topEdge()] == 1) || topEdge == c.topEdge()

			if leftMatch && topMatch {
				d.adjustedCameras[candidate.idNumber] = true
				return c, true
			}
		}
	}

	return nil, false
}

func (d *Computer) fitCamera(cam *camera, leftEdge, topEdge string) *camera {
	for _, c := range cam.getCombinationsOfCamera() {
		leftMatch := (leftEdge == "" && d.allEdges[c.leftEdge()] == 1) || leftEdge == c.leftEdge()
		topMatch := (topEdge == "" && d.allEdges[c.topEdge()] == 1) || topEdge == c.topEdge()

		if leftMatch && topMatch {
			d.adjustedCameras[cam.idNumber] = true
			return c
		}
	}

	return nil
}

func (d *Computer) initialiseArrangeCameras() {
	d.arrangedCameras = make([][]*camera, d.sideLength)

	for i := 0; i < d.sideLength; i++ {
		d.arrangedCameras[i] = make([]*camera, d.sideLength)
	}
}

func (d *Computer) findCamerasAtEdges() []*camera {
	d.allEdges = map[string]int{}

	for _, c := range d.cameras {
		for _, edge := range c.allEdges() {
			d.allEdges[edge]++
		}
	}

	edgeCameras := []*camera{}

	for _, c := range d.cameras {
		uniqueEdges := 0

		for _, edge := range c.allEdges() {
			if d.isUniqueEdge(edge) {
				uniqueEdges++
			}
		}

		if uniqueEdges == 4 {
			edgeCameras = append(edgeCameras, c)
		}
	}

	return edgeCameras
}

func (d *Computer) isUniqueEdge(edge string) bool {
	return d.allEdges[edge]+d.allEdges[flipEdge(edge)] == 2
}

func (d *Computer) transformImageToSeeSeaMonsters() bool {
	d.loadSeaMonsterRunes()

	combinations := getCombinationsOfTiles(d.image)
	for _, combination := range combinations {
		foundMonsters := 0

		for j := 0; j < len(combination); j++ {
			for i := 0; i < len(combination[j]); i++ {
				potentialMonsterBitLocations := []location{}

				assumeMonster := true

			MonsterDetection:
				for y := 0; y < len(d.monsterRunes); y++ {
					for x := 0; x < len(d.monsterRunes[y]); x++ {
						l := location{x: i + x, y: j + y}

						if l.x >= d.imageSize {
							assumeMonster = false
							break MonsterDetection
						}

						if l.y >= d.imageSize {
							assumeMonster = false
							break MonsterDetection
						}

						if d.monsterRunes[y][x] == '#' {
							if combination[l.y][l.x] != '#' {
								assumeMonster = false
								break MonsterDetection
							}

							potentialMonsterBitLocations = append(potentialMonsterBitLocations, l)
						}

					}
				}

				if !assumeMonster {
					continue
				}

				foundMonsters++

				for _, l := range potentialMonsterBitLocations {
					combination[l.y][l.x] = 'O'
				}
			}
		}

		if foundMonsters > 0 {
			d.image = combination
			return true
		}
	}

	return false
}

func (d *Computer) loadSeaMonsterRunes() {
	monsterStrings := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	d.monsterRunes = [][]rune{}

	for _, s := range monsterStrings {
		d.monsterRunes = append(d.monsterRunes, []rune(s))
	}
}

func (d *Computer) createImage() {
	d.removeTileBorders()

	cameraSize := len(d.arrangedCameras[0][0].tiles)
	d.imageSize = d.sideLength * cameraSize

	d.image = make([][]rune, d.imageSize)
	d.monsterPresence = make([][]bool, d.imageSize)

	for i := 0; i < d.imageSize; i++ {
		d.image[i] = make([]rune, d.imageSize)
		d.monsterPresence[i] = make([]bool, d.imageSize)
	}

	for row := 0; row < d.sideLength; row++ {
		for col := 0; col < d.sideLength; col++ {
			for i := 0; i < cameraSize; i++ {
				for j := 0; j < cameraSize; j++ {
					d.image[row*cameraSize+i][col*cameraSize+j] = d.arrangedCameras[row][col].tiles[i][j]
				}
			}
		}
	}
}

func printTiles(tiles [][]rune) {
	fmt.Println()
	for i := 0; i < len(tiles); i++ {
		fmt.Printf("%s\n", string(tiles[i]))
	}
}

func (d *Computer) removeTileBorders() {
	for row := 0; row < d.sideLength; row++ {
		for col := 0; col < d.sideLength; col++ {
			d.arrangedCameras[row][col].removeTileBorders()
		}
	}
}

func (c *camera) removeTileBorders() {
	newSideLength := len(c.tiles) - 2

	newTiles := make([][]rune, newSideLength)

	for row := 0; row < newSideLength; row++ {
		newTiles[row] = make([]rune, newSideLength)

		for col := 0; col < newSideLength; col++ {
			newTiles[row][col] = c.tiles[row+1][col+1]
		}
	}

	c.tiles = newTiles
}

func (d *Computer) multiplyCornerCameraIds() int {
	return d.arrangedCameras[0][0].idNumber * d.arrangedCameras[0][d.sideLength-1].idNumber *
		d.arrangedCameras[d.sideLength-1][0].idNumber * d.arrangedCameras[d.sideLength-1][d.sideLength-1].idNumber
}

func (d *Computer) loadCameras(input days.Input) {
	d.cameras = []*camera{}

	for _, c := range strings.Split(strings.Join(input, "\n"), "\n\n") {
		d.cameras = append(d.cameras, d.loadSingleCamera(c))
	}

	d.sideLength = int(math.Sqrt(float64(len(d.cameras))))
}

func (d *Computer) loadSingleCamera(cameraInput string) *camera {
	c := camera{}

	lines := strings.Split(cameraInput, "\n")

	c.idNumber, _ = strconv.Atoi(d.tileIDMatcher.FindString(lines[0]))

	for _, line := range lines[1:] {
		c.tiles = append(c.tiles, []rune(line))
	}

	return &c
}

func (c *camera) topEdge() string {
	edge := make([]rune, len(c.tiles))

	for i := 0; i < len(c.tiles); i++ {
		edge[i] = c.tiles[0][i]
	}

	return string(edge)
}

func (c *camera) bottomEdge() string {
	edge := make([]rune, len(c.tiles))

	for i := 0; i < len(c.tiles); i++ {
		edge[i] = c.tiles[len(c.tiles[0])-1][i]
	}

	return string(edge)
}

func (c *camera) leftEdge() string {
	edge := make([]rune, len(c.tiles))

	for i := 0; i < len(c.tiles); i++ {
		edge[i] = c.tiles[i][0]
	}

	return string(edge)
}

func (c *camera) rightEdge() string {
	edge := make([]rune, len(c.tiles))

	for i := 0; i < len(c.tiles); i++ {
		edge[i] = c.tiles[i][len(c.tiles[0])-1]
	}

	return string(edge)
}

func (c *camera) allEdges() []string {
	return []string{
		c.topEdge(),
		c.bottomEdge(),
		c.leftEdge(),
		c.rightEdge(),
		flipEdge(c.topEdge()),
		flipEdge(c.bottomEdge()),
		flipEdge(c.leftEdge()),
		flipEdge(c.rightEdge()),
	}
}

func flipEdge(edge string) string {
	bits := []rune(edge)
	flipped := make([]rune, len(bits))

	for i := len(bits)/2 - 1; i >= 0; i-- {
		opp := len(bits) - 1 - i
		flipped[i], flipped[opp] = bits[opp], bits[i]
	}

	return string(flipped)
}

func (c *camera) getCombinationsOfCamera() []*camera {
	combinations := []*camera{}

	for _, tilesCombination := range getCombinationsOfTiles(c.tiles) {
		combinations = append(combinations, &camera{idNumber: c.idNumber, tiles: tilesCombination})
	}

	return combinations
}

func (c *camera) getRotationsOfCamera() []*camera {
	rotations := []*camera{}

	for _, tilesRotation := range getRotationsOfTiles(c.tiles) {
		rotations = append(rotations, &camera{idNumber: c.idNumber, tiles: tilesRotation})
	}

	return rotations
}

func getCombinationsOfTiles(original [][]rune) [][][]rune {
	combinations := make([][][]rune, 16)

	i := 0
	rotations := getRotationsOfTiles(original)
	n := len(rotations)

	for _, rotation := range rotations {
		combinations[i*n] = rotation
		combinations[i*n+1] = flipTilesVertically(rotation)
		combinations[i*n+2] = flipTilesHorizontally(rotation)
		combinations[i*n+3] = flipTilesVertically(flipTilesHorizontally(rotation))
		i++
	}

	return combinations
}

func getRotationsOfTiles(original [][]rune) [][][]rune {
	rotations := make([][][]rune, 4)

	rotation := original

	for i := 0; i < 4; i++ {
		rotations[i] = rotation
		rotation = rotateTiles(rotation)
	}

	return rotations
}

func flipTilesVertically(tiles [][]rune) [][]rune {
	n := len(tiles)
	res := make([][]rune, n)

	for i := 0; i < n; i++ {
		res[i] = tiles[n-i-1]
	}

	return res
}

func flipTilesHorizontally(tiles [][]rune) [][]rune {
	n := len(tiles)
	res := make([][]rune, n)

	for i := 0; i < n; i++ {
		res[i] = make([]rune, n)

		for j := 0; j < n; j++ {
			res[i][j] = tiles[i][n-j-1]
		}
	}

	return res
}

func rotateTiles(tiles [][]rune) [][]rune {
	n := len(tiles)
	res := make([][]rune, n)

	for i := 0; i < n; i++ {
		res[i] = make([]rune, n)

		for j := 0; j < n; j++ {
			res[i][j] = tiles[n-j-1][i]
		}
	}

	return res
}

func (c *camera) validateDimensionsBeforeCheck(cc *camera) {
	if len(c.tiles) != len(cc.tiles) || len(c.tiles[0]) != len(cc.tiles[0]) {
		panic("This can't be")
	}
}
