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

func Part1(input string) string {
	lines := util.ListOfLines(input)
	total := 0

	for _, line := range lines {
		re := regexp.MustCompile(`\d`)
		match := re.FindAllString(line, -1)

		if len(match) > 0 {
			nextVal, _ := strconv.Atoi(match[0] + match[len(match)-1])
			total += nextVal
		}
	}

	return fmt.Sprint(total)
}

func Part2(input string) string {
	lines := util.ListOfLines(input)
	total := 0

	digitsSpelledOut := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range lines {
		re := regexp.MustCompile(`\d|` + strings.Join(digitsSpelledOut, "|"))

		if len(line) > 0 {
			firstDigit := re.FindString(line)
			secondDigit := ""

			for i := len(line) - 1; i >= 0; i-- {
				secondDigit = re.FindString(line[i:])

				if secondDigit != "" {
					break
				}
			}

			for i, spelledDigit := range digitsSpelledOut {
				if firstDigit == spelledDigit {
					firstDigit = fmt.Sprint(i)
				}

				if secondDigit == spelledDigit {
					secondDigit = fmt.Sprint(i)
				}
			}

			nextVal, _ := strconv.Atoi(firstDigit + secondDigit)
			total += nextVal
		}
	}

	return fmt.Sprint(total)
}
