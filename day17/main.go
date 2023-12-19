package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"

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

// A location on the grid
type Coord struct {
	row int
	col int
}

// The state of part of a path
type NodeState struct {
	location Coord
	direction int
	moveStreak int
}

// Map of direction integers to directional grid movement
var Directions = []Coord{
	0: {-1, 0},
	1: {0, 1},
	2: {1, 0},
	3: {0, -1},
}

/* Use depth first search to find the shortest weighted path from start to end with directional limits */
func findLowestHeatLoss(grid [][]int, minMoveStreak int, maxMoveStreak int) int {
	queue := []NodeState{NodeState{Coord{0,0}, Right, 0}, NodeState{Coord{0,0}, Down, 0}}
	var heatLosses = make(map[NodeState]int)

	end := Coord{len(grid)-1, len(grid[0])-1}
	minHeatLoss := math.MaxInt

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.location == end && current.moveStreak >= minMoveStreak {
			minHeatLoss = util.MinInt(heatLosses[current], minHeatLoss)
		}

		directions := []int{current.direction - 1, current.direction, (current.direction + 1) % 4}

		if directions[0] == -1 {
			directions[0] = 3
		}

		for _, direction := range(directions) {
			directionMove := Directions[direction]
			nextCoord := Coord{current.location.row + directionMove.row, current.location.col + directionMove.col}

			if nextCoord.row < 0 || nextCoord.row >= len(grid) || nextCoord.col < 0 || nextCoord.col >= len(grid[0]) {
				continue
			}

			nextHeatLoss := heatLosses[current] + grid[nextCoord.row][nextCoord.col]
			nextMoveStreak := 1

			if direction == current.direction {
				nextMoveStreak = current.moveStreak + 1
			}

			if (direction == current.direction && current.moveStreak >= maxMoveStreak) ||
				(direction != current.direction && current.moveStreak < minMoveStreak) {
					continue
				}

			nextState := NodeState{
				location: nextCoord,
				direction: direction,
				moveStreak: nextMoveStreak,
			}

			if heatLoss, found := heatLosses[nextState]; !found || heatLoss > nextHeatLoss {
				heatLosses[nextState] = nextHeatLoss
				queue = append(queue, nextState)
			}
		}
	}

	return minHeatLoss
}

/* Find the lowest heat loss path with a max of 3 steps in one direction */
func Part1(input string) string {
	grid := util.MatrixOfDigits(input, -1)
	return fmt.Sprint(findLowestHeatLoss(grid, 1, 3))
}

/* Find the lowest heat loss path with a min of 4 and max of 10 steps in one direction */
func Part2(input string) string {
	grid := util.MatrixOfDigits(input, -1)
	return fmt.Sprint(findLowestHeatLoss(grid, 4, 10))
}
