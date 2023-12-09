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

func checkAllZeros(values []int) bool {
	allZeros := true

	for _, value := range(values) {
		if value != 0 {
			allZeros = false
			break
		}
	}

	return allZeros
}

func getDifferenceSet(values []int) []int {
	var newValues []int

	for i := 0; i < len(values) - 1; i++  {
		newValues = append(newValues, values[i+1] - values[i])
	}

	return newValues
}

func extrapolateLast(values []int) int {
	if checkAllZeros(values) {
		return 0
	}

	newValues := getDifferenceSet(values)
	return values[len(values) - 1] + extrapolateLast(newValues)
}

func extrapolateFirst(values []int) int {
	if checkAllZeros(values) {
		return 0
	}

	newValues := getDifferenceSet(values)
	return values[0] - extrapolateFirst(newValues)
}

func Part1(input string) string {
	values := util.MatrixOfInts(input)
	totalNexts := 0

	for _, intSet := range(values) {
		totalNexts += extrapolateLast(intSet)
	}

	return fmt.Sprint(totalNexts)
}

func Part2(input string) string {
	values := util.MatrixOfInts(input)
	totalPrevs := 0

	for _, intSet := range(values) {
		totalPrevs += extrapolateFirst(intSet)
	}

	return fmt.Sprint(totalPrevs)
}
