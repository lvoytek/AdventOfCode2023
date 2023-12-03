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

type PartNumber struct {
	value int
	startCol int
	endCol int
	row int
}

type Symbol struct {
	value rune
	row int
	col int
}

func isSymbol(input rune) bool {
	return (int(input) < int('0') || int(input) > int('9')) && input != '.'
}

func isNumberAdjacentToSymbol(number PartNumber, symbol Symbol) bool {
	return symbol.row >= number.row - 1 && symbol.row <= number.row + 1 &&
		symbol.col >= number.startCol - 1 && symbol.col <= number.endCol + 1
}

func Part1(input string) string {
	lines := util.ListOfLines(input)

	var symbols []Symbol
	var numbers []PartNumber

	for row, line := range(lines) {
		inANumber := false

		for col, nextChar := range(line) {
			if isSymbol(nextChar) {
				symbols = append(symbols, Symbol {
					value: nextChar,
					row: row,
					col: col,
				})
			}

			isNumber := util.RuneIsNumber(nextChar)

			if isNumber && !inANumber {
				inANumber = true

				value := util.GetFirstIntInString(line[col:])
				numbers = append(numbers, PartNumber{
					value: value,
					row: row,
					startCol: col,
					endCol: col + len(fmt.Sprint(value)) - 1,
				})
			}

			if !isNumber {
				inANumber = false
			}
		}
	}

	sumOfPartNumbers := 0

	for _, number := range(numbers) {
		for _, symbol := range(symbols) {
			if isNumberAdjacentToSymbol(number, symbol) {
				sumOfPartNumbers += number.value
				break
			}
		}
	}

	return fmt.Sprint(sumOfPartNumbers)
}

func Part2(input string) string {
	lines := util.ListOfLines(input)

	var symbols []Symbol
	var numbers []PartNumber

	for row, line := range(lines) {
		inANumber := false

		for col, nextChar := range(line) {
			if isSymbol(nextChar) {
				symbols = append(symbols, Symbol {
					value: nextChar,
					row: row,
					col: col,
				})
			}

			isNumber := util.RuneIsNumber(nextChar)

			if isNumber && !inANumber {
				inANumber = true

				value := util.GetFirstIntInString(line[col:])
				numbers = append(numbers, PartNumber{
					value: value,
					row: row,
					startCol: col,
					endCol: col + len(fmt.Sprint(value)) - 1,
				})
			}

			if !isNumber {
				inANumber = false
			}
		}
	}

	sumOfGearRatios := 0

	for _, symbol := range(symbols) {
		if symbol.value == '*' {
			var adjacentNumbers []PartNumber

			for _, number := range(numbers) {
				if isNumberAdjacentToSymbol(number, symbol) {
					adjacentNumbers = append(adjacentNumbers, number)
				}
			}

			if len(adjacentNumbers) == 2 {
				sumOfGearRatios += adjacentNumbers[0].value * adjacentNumbers[1].value
			}
		}
	}

	return fmt.Sprint(sumOfGearRatios)
}
