package main

import (
	_ "embed"
	"flag"
	"fmt"
	"hash/fnv"
	"regexp"

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

func tiltWest(dish [][]rune) [][]rune {
	for row, line := range(dish) {
		moveTo := 0

		for col, rock := range(line) {
			if rock == '#' {
				moveTo = col + 1
			} else if rock == 'O' {
				dish[row][col] = '.'
				dish[row][moveTo] = 'O'
				moveTo ++
			}
		}
	}

	return dish
}

func tiltNorth(dish [][]rune) [][]rune {
	for col := 0; col < len(dish[0]); col++ {
		moveTo := 0

		for row, line := range(dish) {
			if line[col] == '#' {
				moveTo = row + 1
			} else if line[col] == 'O' {
				dish[row][col] = '.'
				dish[moveTo][col] = 'O'
				moveTo ++
			}
		}
	}

	return dish
}

func tiltSouth(dish [][]rune) [][]rune {
	for col := 0; col < len(dish[0]); col++ {
		moveTo := len(dish) - 1

		for row := len(dish) - 1; row >= 0; row -- {
			line := dish[row]

			if line[col] == '#' {
				moveTo = row - 1
			} else if line[col] == 'O' {
				dish[row][col] = '.'
				dish[moveTo][col] = 'O'
				moveTo --
			}
		}
	}

	return dish
}

func tiltEast(dish [][]rune) [][]rune {
	for row, line := range(dish) {
		moveTo := len(line) - 1

		for col := len(line) - 1; col >= 0; col -- {
			rock := line[col]

			if rock == '#' {
				moveTo = col - 1
			} else if rock == 'O' {
				dish[row][col] = '.'
				dish[row][moveTo] = 'O'
				moveTo --
			}
		}
	}

	return dish
}

func runCycle(dish [][]rune) [][]rune {
	dish = tiltNorth(dish)
	dish = tiltWest(dish)
	dish = tiltSouth(dish)
	dish = tiltEast(dish)
	return dish
}

func hashDish(dish [][]rune) uint32 {
	var dishList []rune

	for _, line := range(dish) {
		dishList = append(dishList, line...)
	}

	hash := fnv.New32a()
	hash.Write([]byte(string(dishList)))
	return hash.Sum32()
}

func dishExistsInHashSet(newDish uint32, hashes []uint32) (bool, int) {
	for i, hash := range(hashes) {
		if hash == newDish {
			return true, i
		}
	}

	return false, 0
}

func getLoad(dish [][]rune) int {
	totalLoad := 0
	rockRe := regexp.MustCompile(`O`)

	for row, lineRunes := range(dish) {
		line := string(lineRunes)
		totalLoad += len(rockRe.FindAllStringIndex(line, -1)) * (len(dish) - row)
	}

	return totalLoad
}

func Part1(input string) string {
	dish := util.MatrixOfCharacters(input)
	movedDish := tiltNorth(dish)
	return fmt.Sprint(getLoad(movedDish))
}

func Part2(input string) string {
	const cyclesToDo = 1000000000

	dish := util.MatrixOfCharacters(input)

	var cyclesSoFar int
	var hashedDishes []uint32
	var matchingDishIndex int

	for cyclesSoFar = 1; cyclesSoFar < cyclesToDo; cyclesSoFar++ {
		dish = runCycle(dish)

		movedDishHash := hashDish(dish)
		dishExists, mDI := dishExistsInHashSet(movedDishHash, hashedDishes)

		if dishExists {
			matchingDishIndex = mDI
			break
		}

		hashedDishes = append(hashedDishes, movedDishHash)
	}

	if cyclesSoFar < cyclesToDo {
		cyclesLeft := (cyclesToDo - cyclesSoFar) % (cyclesSoFar - matchingDishIndex - 1)

		for i := 0; i < cyclesLeft; i++ {
			dish = runCycle(dish)
		}
	}

	return fmt.Sprint(getLoad(dish))
}
