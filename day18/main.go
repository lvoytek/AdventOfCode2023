package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lvoytek/AdventOfCode2023/util"
)

//go:embed input.txt
var inputData string

func main() {
	var doPart int
	flag.IntVar(&doPart, "p", 1, "Part number to run:")
	flag.Parse()

	if doPart != 2 {
		fmt.Println("Part 1 Output:", Part1(inputData))
	}

	if doPart != 1 {
		fmt.Println("Part 2 Output:", Part2(inputData))
	}
}

// Coordinate on a grid
type Coord struct {
	row int
	col int
}

// Map the final character of a color to a direction character
var ColorEndToDirection = map[rune]rune {
	'0': 'R',
	'1': 'D',
	'2': 'L',
	'3': 'U',
}

// Map direction characters to movement on a grid
var Directions = map[rune]Coord {
	'U': {-1, 0},
	'D': {1, 0},
	'L': {0, -1},
	'R': {0, 1},
}

// A set of digs made by an elf in a certain direction
type DigAction struct {
	direction rune
	distance int
}

/* Create a matrix of booleans showing the dig outline */
func createDigGrid(digActions []DigAction) [][]bool {
	maxLoc := Coord{0,0}
	minLoc := Coord{0,0}
	currentLoc := Coord{0,0}

	for _, digAction := range(digActions) {
		movement := Directions[digAction.direction]
		currentLoc.row += movement.row * digAction.distance
		currentLoc.col += movement.col * digAction.distance

		if currentLoc.row > maxLoc.row {
			maxLoc.row = currentLoc.row
		} else if currentLoc.row < minLoc.row {
			minLoc.row = currentLoc.row
		}

		if currentLoc.col > maxLoc.col {
			maxLoc.col = currentLoc.col
		} else if currentLoc.col < minLoc.col {
			minLoc.col = currentLoc.col
		}
	}

	var digGrid = make([][]bool, maxLoc.row + 1 - minLoc.row)

	for row := 0; row < len(digGrid); row++ {
		digGrid[row] = make([]bool, maxLoc.col + 1 - minLoc.col)
	}

	startLoc := Coord{-minLoc.row, -minLoc.col}
	digGrid[startLoc.row][startLoc.col] = true

	for _, digAction := range(digActions) {
		movement := Directions[digAction.direction]
		for i:=0; i < digAction.distance; i++ {
			startLoc.row += movement.row
			startLoc.col += movement.col
			digGrid[startLoc.row][startLoc.col] = true
		}
	}

	return digGrid
}

/* The recursive portion of grid filling */
func recursiveFill(current Coord, digGrid [][]bool) [][]bool{
	if digGrid[current.row][current.col] {
		return digGrid
	}

	digGrid[current.row][current.col] = true

	digGrid = recursiveFill(Coord{current.row + 1, current.col}, digGrid)
	digGrid = recursiveFill(Coord{current.row - 1, current.col}, digGrid)
	digGrid = recursiveFill(Coord{current.row, current.col + 1}, digGrid)
	digGrid = recursiveFill(Coord{current.row, current.col - 1}, digGrid)

	return digGrid
}

/* Fill the inside of an outline recursively */
func fillGrid(digGrid [][]bool) [][]bool {
	var initialLoc Coord

	for i, location := range(digGrid[0]) {
		if location && !digGrid[1][i] {
			initialLoc = Coord{1,i}
		}
	}

	digGrid = recursiveFill(initialLoc, digGrid)

	return digGrid
}

/* Get the set of coordinates for the corners of the outline */
func getDigCornerCoords(digActions []DigAction) []Coord {
	var coords []Coord
	currentCoord := Coord{0,0}

	for _, digAction := range(digActions) {
		coords = append(coords, currentCoord)

		movement := Directions[digAction.direction]
		currentCoord.row += movement.row * digAction.distance
		currentCoord.col += movement.col * digAction.distance
	}

	return coords
}

/* Use shoelace theorem to get the internal area within a set of coordinates */
func getAreaWithShoelace(coords []Coord) int64 {
	a1 := int64(coords[0].col) * int64(coords[len(coords)-1].row)
	a2 := int64(coords[0].row) * int64(coords[len(coords)-1].col)

	for i := 0; i < len(coords)-1; i++ {
		a1 += int64(coords[i].col) * int64(coords[i+1].row)
		a2 += int64(coords[i].row) * int64(coords[i+1].col)
	}

	totalArea := a1 - a2

	if totalArea < 0 {
		totalArea *= -1
	}

	totalArea = totalArea / 2

	return totalArea
}

/* Find the perimeter of the dig / the area of the outline */
func getPerimeter(digActions []DigAction) int64 {
	perimeter := int64(0)

	for _, digAction := range(digActions) {
		perimeter += int64(digAction.distance)
	}

	return perimeter
}

/* Find the volume of a dig outline and the area within to be digged */
func Part1(input string) string {
	lines := util.ListOfLines(input)

	var digActions []DigAction

	for _, line := range(lines) {
		splitLines := strings.Split(line, " ")
		digActions = append(digActions, DigAction{
			direction: rune(splitLines[0][0]),
			distance: util.GetFirstIntInString(splitLines[1]),
		})
	}

	grid := createDigGrid(digActions)
	grid = fillGrid(grid)

	totalVolume := 0

	for _, line := range(grid) {
		for _, current := range(line) {
			if current {
				totalVolume ++
			}
		}
	}

	return fmt.Sprint(totalVolume)
}

/* Find the same volume but convert colors to the outline instead */
func Part2(input string) string {
	lines := util.ListOfLines(input)

	var digActions []DigAction

	colorRe := regexp.MustCompile(`#[a-f0-9]+`)

	for _, line := range(lines) {
		color := colorRe.FindString(line)[1:]
		distance, _ := strconv.ParseUint(color[:len(color)-1], 16, 32)
		digActions = append(digActions, DigAction{
			direction: ColorEndToDirection[rune(color[len(color)-1])],
			distance: int(distance),
		})
	}

	coords := getDigCornerCoords(digActions)
	areaIncludingOutline := getAreaWithShoelace(coords) + getPerimeter(digActions) / 2 + 1

	return fmt.Sprint(areaIncludingOutline)
}
