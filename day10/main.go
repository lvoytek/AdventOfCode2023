package main

import (
	_ "embed"
	"flag"
	"fmt"
	"reflect"
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

// Location on the map
type Coord struct {
	row int
	col int
}

// A single pipe in a given location with connections
type Pipe struct {
	pipeType rune
	location Coord
	connectedPipes []*Pipe
	visited bool
}

// Directions for pipe route tracing
const (
	Up int = 0
	Right = 1
	Down = 2
	Left = 3
)

/* Create a list of the pipes contained in the map */
func createPipes(lines []string) (pipes []Pipe) {
	for row, line := range(lines) {
		for col, item := range(line) {
			if item != '.' {
				pipes = append(pipes, Pipe {
					pipeType: item,
					location: Coord {
						row: row,
						col: col,
					},
					visited: false,
				})
			}
		}
	}

	return pipes
}

/* Find all possible locations that a given pipe can connect to */
func getAdjacentPipeCoords(pipe Pipe) (coords []Coord) {
	if strings.ContainsRune("|LJS", pipe.pipeType) {
		coords = append(coords, Coord {
			row: pipe.location.row - 1,
			col: pipe.location.col,
		})
	}

	if strings.ContainsRune("-LFS", pipe.pipeType) {
		coords = append(coords, Coord {
			row: pipe.location.row,
			col: pipe.location.col + 1,
		})
	}

	if strings.ContainsRune("|7FS", pipe.pipeType) {
		coords = append(coords, Coord {
			row: pipe.location.row + 1,
			col: pipe.location.col,
		})
	}

	if strings.ContainsRune("-J7S", pipe.pipeType) {
		coords = append(coords, Coord {
			row: pipe.location.row,
			col: pipe.location.col - 1,
		})
	}

	return coords
}

/* For each pipe find all other pipes actually connected to it and add them to their list */
func connectPipes(pipes []Pipe) []Pipe {
	for i, currentPipe := range(pipes) {
		availableCoords := getAdjacentPipeCoords(currentPipe)

		for j, pipe := range(pipes) {
			for _, coord := range(availableCoords) {
				if reflect.DeepEqual(coord, pipe.location) {
					pipes[i].connectedPipes = append(pipes[i].connectedPipes, &pipes[j])
					break
				}
			}
		}
	}

	return pipes
}

/* Recursively search through pipes to find the longest loop connected to S */
func findLongestLoop(pipes []*Pipe) []*Pipe {
	currentPipe := pipes[len(pipes) - 1]
	if currentPipe.pipeType == 'S' && len(pipes) > 1 {
		return pipes
	}

	greatestDistance := 0
	var greatestPipeline []*Pipe
	routeEnds := true

	currentPipe.visited = true

	for _, nextPipe := range(currentPipe.connectedPipes) {
		if !nextPipe.visited || nextPipe.pipeType == 'S' {
			routeEnds = false
			nextPipeline := findLongestLoop(append(pipes, nextPipe))

			if len(nextPipeline) > greatestDistance {
				greatestDistance = len(nextPipeline)
				greatestPipeline = nextPipeline
			}
		}
	}

	currentPipe.visited = false

	if routeEnds {
		return []*Pipe{}
	}

	return greatestPipeline
}

/* Recursively fill an empty area on the map with a given character */
func fillArea(row int, col int, fillWith rune, cleanedMap [][]rune) [][]rune {
	if row < 0 || row >= len(cleanedMap) || col < 0 || col >= len(cleanedMap[row]) || cleanedMap[row][col] != '.' {
		return cleanedMap
	}

	cleanedMap[row][col] = fillWith

	if row > 0 && cleanedMap[row - 1][col] == '.' {
		cleanedMap = fillArea(row - 1, col, fillWith, cleanedMap)
	}

	if row < len(cleanedMap) - 1 && cleanedMap[row + 1][col] == '.' {
		cleanedMap = fillArea(row + 1, col, fillWith, cleanedMap)
	}

	if col > 0 && cleanedMap[row][col - 1] == '.' {
		cleanedMap = fillArea(row, col - 1, fillWith, cleanedMap)
	}

	if col < len(cleanedMap[row]) - 1 && cleanedMap[row][col + 1] == '.' {
		cleanedMap = fillArea(row, col + 1, fillWith, cleanedMap)
	}

	return cleanedMap
}

/* Find distance from S to furthest away point on the longest pipe loop */
func Part1(input string) string {
	// Parse input and extract all pipes from the map
	lines := util.ListOfLines(input)
	pipes := createPipes(lines)
	pipes = connectPipes(pipes)

	var startPipe *Pipe

	for _, pipe := range(pipes) {
		if pipe.pipeType == 'S' {
			startPipe = &pipe
			break
		}
	}

	// Get the length of the longest loop then divide by 2 for furthest
	return fmt.Sprint(len(findLongestLoop([]*Pipe{startPipe})) / 2)
}

/* Find the area located within the longest pipe loop */
func Part2(input string) string {
	// Parse input and extract all pipes from the map
	lines := util.ListOfLines(input)
	pipes := createPipes(lines)
	pipes = connectPipes(pipes)

	var startPipe *Pipe

	for _, pipe := range(pipes) {
		if pipe.pipeType == 'S' {
			startPipe = &pipe
			break
		}
	}

	// Find longest loop then remove all other pipes from the map
	loop := findLongestLoop([]*Pipe{startPipe})
	var cleanedMap [][]rune

	for row, line := range(lines) {
		cleanedMap = append(cleanedMap, []rune{})
		for col, _ := range(line) {
			inLoop := false
			for _, pipe := range(loop) {
				if row == pipe.location.row && col == pipe.location.col {
					cleanedMap[row] = append(cleanedMap[row], pipe.pipeType)
					inLoop = true
					break
				}
			}

			if !inLoop {
				cleanedMap[row] = append(cleanedMap[row], '.')
			}
		}
	}

	// Trace the longest loop and fill each side with a distinct character
	var direction int

	if loop[1].location.row < loop[0].location.row {
		direction = Up
	} else if loop[1].location.col > loop[0].location.col {
		direction = Right
	} else if loop[1].location.row > loop[0].location.row {
		direction = Down
	} else {
		direction = Left
	}

	for _, pipe := range(loop) {
		prevDirection := direction

		if pipe.pipeType == 'L' {
			if direction == Left {
				direction = Up
			} else if direction == Down {
				direction = Right
			}
		} else if pipe.pipeType == 'J' {
			if direction == Right {
				direction = Up
			} else if direction == Down {
				direction = Left
			}
		} else if pipe.pipeType == '7' {
			if direction == Right {
				direction = Down
			} else if direction == Up {
				direction = Left
			}
		} else if pipe.pipeType == 'F' {
			if direction == Up {
				direction = Right
			} else if direction == Left {
				direction = Down
			}
		}

		if direction == Up {
			cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'B', cleanedMap)
		} else if direction == Right {
			cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'B', cleanedMap)
		} else if direction == Down {
			cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'B', cleanedMap)
		} else {
			cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'B', cleanedMap)
		}

		if prevDirection != direction {
			if prevDirection == Up {
				cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'B', cleanedMap)
			} else if prevDirection == Right {
				cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'B', cleanedMap)
			} else if prevDirection == Down {
				cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'B', cleanedMap)
			} else {
				cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'B', cleanedMap)
			}
		}
	}

	// Determine which side is the inside
	lookFor := 'A'

	for _, location := range(cleanedMap[0]) {
		if location == 'A' {
			lookFor = 'B'
			break
		}
	}

	// Count the internal area
	totalInside := 0

	for _, line := range(cleanedMap) {
		for _, location := range(line) {
			if location == lookFor {
				totalInside ++
			}
		}
	}

	return fmt.Sprint(totalInside)
}
