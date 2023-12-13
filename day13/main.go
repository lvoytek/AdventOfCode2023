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

func doRowsMatch(matrix []string, rowX int, rowY int) bool {
	for i := 0; i < len(matrix[rowX]); i++ {
		if matrix[rowX][i] != matrix[rowY][i] {
			return false
		}
	}

	return true
}

func doRowsMatchPossibleSmudge(matrix []string, rowX int, rowY int) (bool, bool) {
	foundSmudge := false

	for i := 0; i < len(matrix[rowX]); i++ {
		if matrix[rowX][i] != matrix[rowY][i] {
			if !foundSmudge {
				foundSmudge = true
			} else {
				return false, false
			}
		}
	}

	return true, foundSmudge
}

func doColsMatch(matrix []string, colX int, colY int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][colX] != matrix[i][colY] {
			return false
		}
	}

	return true
}

func doColsMatchPossibleSmudge(matrix []string, colX int, colY int) (bool, bool) {
	foundSmudge := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][colX] != matrix[i][colY] {
			if !foundSmudge {
				foundSmudge = true
			} else {
				return false, false
			}
		}
	}

	return true, foundSmudge
}

func checkRowMirror(matrix []string, highMid int, hasSmudge bool) bool {
	iRange := util.MinInt(highMid, len(matrix) - highMid)

	if hasSmudge {
		foundSmudge := false

		for i := 0; i < iRange; i++ {
			valid, newSmudge := doRowsMatchPossibleSmudge(matrix, highMid - i - 1, highMid + i)
			if !valid || (foundSmudge && newSmudge) {
				return false
			} else if newSmudge {
				foundSmudge = true
			}
		}

		if !foundSmudge {
			return false
		}
	} else {
		for i := 0; i < iRange; i++ {
			if !doRowsMatch(matrix, highMid - i - 1, highMid + i) {
				return false
			}
		}
	}

	return true
}

func checkColMirror(matrix []string, highMid int, hasSmudge bool) bool {
	iRange := util.MinInt(highMid, len(matrix[0]) - highMid)

	if hasSmudge {
		foundSmudge := false

		for i := 0; i < iRange; i++ {
			valid, newSmudge := doColsMatchPossibleSmudge(matrix, highMid - i - 1, highMid + i)
			if !valid || (foundSmudge && newSmudge) {
				return false
			} else if newSmudge {
				foundSmudge = true
			}
		}

		if !foundSmudge {
			return false
		}
	} else {
		for i := 0; i < iRange; i++ {
			if !doColsMatch(matrix, highMid - i - 1, highMid + i) {
				return false
			}
		}
	}

	return true
}

func runPart(input string, useSmudge bool) string {
	chunks := util.ListOfLineChunks(input)
	totalNotes := 0

	for _, matrix := range(chunks) {
		foundRow := false
		for i := 1; i < len(matrix); i++ {
			if checkRowMirror(matrix, i, useSmudge) {
				totalNotes += i * 100
				foundRow = true
				break
			}
		}

		if foundRow {
			continue
		}

		for i := 1; i < len(matrix[0]); i++ {
			if checkColMirror(matrix, i, useSmudge) {
				totalNotes += i
				break
			}
		}
	}

	return fmt.Sprint(totalNotes)
}

func Part1(input string) string {
	return runPart(input, false)
}

func Part2(input string) string {
	return runPart(input, true)
}
