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

const (
	Up int = 0
	Right = 1
	Down = 2
	Left = 3
)

type LightLocation struct {
	row int
	col int
	direction int
	complete bool
}

/* Follow a laser down its path on the grid and report all locations it touched */
func shineLaser(grid []string, initialConfig LightLocation) [][]bool {
	var energizedTiles [][]bool
	var prevDirections [][]int

	// Initialize the return grid and directions previously traversed
	for row := 0; row < len(grid); row++ {
		energizedTiles = append(energizedTiles, []bool{})
		prevDirections = append(prevDirections, []int{})

		for col := 0; col < len(grid[0]); col++ {
			energizedTiles[row] = append(energizedTiles[row], false)
			prevDirections[row] = append(prevDirections[row], -1)
		}
	}

	// Run through laser beam until all branches are terminated or looped
	locations := []LightLocation{ initialConfig }

	for {
		var locationsToAppend []LightLocation
		allComplete := true

		for i, location := range(locations) {
			if location.complete || location.row >= len(grid) || location.row < 0 ||
				location.col >= len(grid[0]) || location.col < 0 ||
				prevDirections[location.row][location.col] == location.direction {

				locations[i].complete = true
				continue
			}

			allComplete = false
			energizedTiles[location.row][location.col] = true

			if grid[location.row][location.col] == '.' {
				prevDirections[location.row][location.col] = location.direction
			} else if grid[location.row][location.col] == '/' {
				if location.direction % 2 == 0 {
					locations[i].direction ++
				} else {
					locations[i].direction --
				}
			} else if grid[location.row][location.col] == '\\' {
				if location.direction == Up {
					locations[i].direction = Left
				} else if location.direction == Left {
					locations[i].direction = Up
				} else if location.direction == Down {
					locations[i].direction = Right
				} else {
					locations[i].direction = Down
				}
			} else if grid[location.row][location.col] == '-' {
				if location.direction % 2 == 0 {
					locations[i].direction = Left
					locationsToAppend = append(locationsToAppend, LightLocation {
						row: location.row,
						col: location.col + 1,
						direction: Right,
						complete: false,
					})
				}
			} else if grid[location.row][location.col] == '|' {
				if location.direction % 2 == 1 {
					locations[i].direction = Up
					locationsToAppend = append(locationsToAppend, LightLocation {
						row: location.row + 1,
						col: location.col,
						direction: Down,
						complete: false,
					})
				}
			}

			if locations[i].direction == Up {
				locations[i].row --
			} else if locations[i].direction == Down {
				locations[i].row ++
			} else if locations[i].direction == Left {
				locations[i].col --
			} else {
				locations[i].col ++
			}

		}

		if allComplete {
			return energizedTiles
		}

		locations = append(locations, locationsToAppend...)
	}
}

/* Find the total grid charge as a sum of charged locations */
func findTotalCharge(energizedTiles [][]bool) int {
	totalCharged := 0

	for _, line := range(energizedTiles) {
		for _, tile := range(line) {
			if tile {
				totalCharged ++
			}
		}
	}

	return totalCharged
}

/* Starting along each edge, find the path with the greatest grid charge and get its charge */
func findBestConfigCharge(grid []string) int {
	bestConfigCharge := 0

	baseLaser := LightLocation {
		row: 0,
		col: 0,
		direction: Down,
		complete: false,
	}

	// Horizontal
	for i, _ := range(grid[0]) {
		baseLaser.col = i

		baseLaser.row = 0
		baseLaser.direction = Down
		nextCharge := findTotalCharge(shineLaser(grid, baseLaser))

		if nextCharge > bestConfigCharge {
			bestConfigCharge = nextCharge
		}

		baseLaser.row = len(grid) - 1
		baseLaser.direction = Up
		nextCharge = findTotalCharge(shineLaser(grid, baseLaser))

		if nextCharge > bestConfigCharge {
			bestConfigCharge = nextCharge
		}
	}

	// Vertical
	for i, _ := range(grid) {
		baseLaser.row = i

		baseLaser.col = 0
		baseLaser.direction = Right
		nextCharge := findTotalCharge(shineLaser(grid, baseLaser))

		if nextCharge > bestConfigCharge {
			bestConfigCharge = nextCharge
		}

		baseLaser.col = len(grid[0]) - 1
		baseLaser.direction = Left
		nextCharge = findTotalCharge(shineLaser(grid, baseLaser))

		if nextCharge > bestConfigCharge {
			bestConfigCharge = nextCharge
		}
	}

	return bestConfigCharge
}

/* Find the total grid charge for the laser starting in the top left */
func Part1(input string) string {
	lines := util.ListOfLines(input)
	energizedTiles := shineLaser(lines, LightLocation {
		row: 0,
		col: 0,
		direction: Right,
		complete: false,
	})

	return fmt.Sprint(findTotalCharge(energizedTiles))
}

/* Find the greatest possible grid charge */
func Part2(input string) string {
	lines := util.ListOfLines(input)
	return fmt.Sprint(findBestConfigCharge(lines))
}
