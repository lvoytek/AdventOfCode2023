package main

import (
	_ "embed"
	"flag"
	"fmt"

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

type Coord struct {
	row int
	col int
}

type Space struct {
	isGalaxy bool
	horizontal int
	vertical int
}

func createGalaxyMatrix(lines []string) (galaxyMatrix [][]bool) {
	for i, line := range(lines) {
		galaxyMatrix = append(galaxyMatrix, []bool{})

		for _, location := range(line) {
			galaxyMatrix[i] = append(galaxyMatrix[i], location == '#')
		}
	}

	return galaxyMatrix
}

func inflateSpace(galaxyMatrix [][]bool, inflationAmount int) [][]Space {
	var rowDistances []int

	for _, line := range(galaxyMatrix) {
		hasGalaxy := false

		for _, location := range(line) {
			if location {
				hasGalaxy = true
				break
			}
		}

		if hasGalaxy {
			rowDistances = append(rowDistances, 1)
		} else {
			rowDistances = append(rowDistances, inflationAmount)
		}
	}


	var colDistances []int

	for col := 0; col < len(galaxyMatrix[0]); col++ {
		hasGalaxy := false

		for _, line := range(galaxyMatrix) {
			if line[col] {
				hasGalaxy = true
				break
			}
		}

		if hasGalaxy {
			colDistances = append(colDistances, 1)
		} else {
			colDistances = append(colDistances, inflationAmount)
		}
	}

	var spaceMatrix [][]Space

	for row, line := range(galaxyMatrix) {
		spaceMatrix = append(spaceMatrix, []Space{})

		for col, location := range(line) {
			spaceMatrix[row] = append(spaceMatrix[row], Space {
				isGalaxy: location,
				vertical: rowDistances[row],
				horizontal: colDistances[col],
			})
		}
	}

	return spaceMatrix
}

func getGalaxyCoords(spaceMatrix [][]Space) (galaxyCoords []Coord) {
	inflatedRow := 0

	for _, line := range(spaceMatrix) {
		inflatedCol := 0

		for _, location := range(line) {
			if location.isGalaxy {
				galaxyCoords = append(galaxyCoords, Coord {
					row: inflatedRow,
					col: inflatedCol,
				})
			}
			inflatedCol += location.horizontal
		}

		inflatedRow += line[0].vertical
	}

	return galaxyCoords
}

func getDistanceBetween(a Coord, b Coord) int {
	return util.AbsInt(b.col - a.col) + util.AbsInt(b.row - a.row)
}

func runPart(input string, inflationAmount int) string {
	lines := util.ListOfLines(input)
	galaxies := createGalaxyMatrix(lines)
	spaceMatrix := inflateSpace(galaxies, inflationAmount)

	galaxyCoords := getGalaxyCoords(spaceMatrix)

	totalDistances := 0

	for i, currentGalaxy := range(galaxyCoords) {
		for _, otherGalaxy := range(galaxyCoords[i+1:]) {
			totalDistances += getDistanceBetween(currentGalaxy, otherGalaxy)
		}
	}

	return fmt.Sprint(totalDistances)
}

func Part1(input string) string {
	return runPart(input, 2)
}

func Part2(input string) string {
	return runPart(input, 1000000)
}
