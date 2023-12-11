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

type Coord struct {
	row int
	col int
}

type Pipe struct {
	pipeType rune
	location Coord
	connectedPipes []*Pipe
	visited bool
}

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

func Part1(input string) string {
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

	return fmt.Sprint(len(findLongestLoop([]*Pipe{startPipe})) / 2)
}

func Part2(input string) string {
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

	//   0
	// 3   1
	//   2

	var direction int

	if loop[1].location.row < loop[0].location.row {
		direction = 0
	} else if loop[1].location.col > loop[0].location.col {
		direction = 1
	} else if loop[1].location.row > loop[0].location.row {
		direction = 2
	} else {
		direction = 3
	}

	for _, pipe := range(loop) {
		prevDirection := direction

		if pipe.pipeType == 'L' {
			if direction == 3 {
				direction = 0
			} else if direction == 2 {
				direction = 1
			}
		} else if pipe.pipeType == 'J' {
			if direction == 1 {
				direction = 0
			} else if direction == 2 {
				direction = 3
			}
		} else if pipe.pipeType == '7' {
			if direction == 1 {
				direction = 2
			} else if direction == 0 {
				direction = 3
			}
		} else if pipe.pipeType == 'F' {
			if direction == 0 {
				direction = 1
			} else if direction == 3 {
				direction = 2
			}
		}

		if direction == 0 {
			cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'B', cleanedMap)
		} else if direction == 1 {
			cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'B', cleanedMap)
		} else if direction == 2 {
			cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'B', cleanedMap)
		} else {
			cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'A', cleanedMap)
			cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'B', cleanedMap)
		}

		if prevDirection != direction {
			if prevDirection == 0 {
				cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'B', cleanedMap)
			} else if prevDirection == 1 {
				cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'B', cleanedMap)
			} else if prevDirection == 2 {
				cleanedMap = fillArea(pipe.location.row, pipe.location.col + 1, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row, pipe.location.col - 1, 'B', cleanedMap)
			} else {
				cleanedMap = fillArea(pipe.location.row + 1, pipe.location.col, 'A', cleanedMap)
				cleanedMap = fillArea(pipe.location.row - 1, pipe.location.col, 'B', cleanedMap)
			}
		}
	}

	lookFor := 'A'

	for _, location := range(cleanedMap[0]) {
		if location == 'A' {
			lookFor = 'B'
			break
		}
	}

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
