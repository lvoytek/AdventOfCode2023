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

type Pull struct {
	reds int
	greens int
	blues int
}

type Game struct {
	id int
	pulls []Pull
}

func extractGame(line string) Game {
	re := regexp.MustCompile(`Game (\d+):`)
	gameNumber, _ := strconv.Atoi(re.FindStringSubmatch(line)[1])

	re = regexp.MustCompile(`(\d+)\s+(red|green|blue)`)
	var pulls []Pull

	for _, pullString := range(strings.Split(line, ";")) {
		var nextPull Pull
		colorSet := re.FindAllString(pullString, -1)

		for _, color := range(colorSet) {
			if strings.Contains(color, "red") {
				nextPull.reds = util.GetFirstIntInString(color)
			} else if strings.Contains(color, "green") {
				nextPull.greens = util.GetFirstIntInString(color)
			} else {
				nextPull.blues = util.GetFirstIntInString(color)
			}
		}

		pulls = append(pulls, nextPull)
	}

	return Game {
		id: gameNumber,
		pulls: pulls,
	}
}

func Part1(input string) string {
	lines := util.ListOfLines(input)

	const maxRed = 12
	const maxGreen = 13
	const maxBlue = 14

	gameSum := 0

	for _, line := range(lines) {
		nextGame := extractGame(line)

		gameIsValid := true

		for _, pull := range(nextGame.pulls) {
			if pull.reds > maxRed || pull.greens > maxGreen || pull.blues > maxBlue {
				gameIsValid = false
				break
			}
		}

		if gameIsValid {
			gameSum += nextGame.id
		}

	}

	return fmt.Sprint(gameSum)
}

func Part2(input string) string {
	lines := util.ListOfLines(input)

	sumOfPower := 0

	for _, line := range(lines) {
		nextGame := extractGame(line)

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, pull := range(nextGame.pulls) {
			if pull.reds > maxRed {
				maxRed = pull.reds
			}

			if pull.greens > maxGreen {
				maxGreen = pull.greens
			}

			if pull.blues > maxBlue {
				maxBlue = pull.blues
			}
		}

		sumOfPower += maxRed * maxGreen * maxBlue
	}

	return fmt.Sprint(sumOfPower)
}
