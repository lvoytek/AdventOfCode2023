package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type Race struct {
	time int
	distance int
}

func calculateDistance(holdTime int, raceTime int) int {
	raceTime -= holdTime
	distance := raceTime * holdTime

	return distance
}

func Part1(input string) string {
	lines := util.ListOfLines(input)

	times := util.GetAllIntsInString(lines[0])
	distances := util.GetAllIntsInString(lines[1])

	var races []Race

	for i, time := range(times) {
		races = append(races, Race {
			time: time,
			distance: distances[i],
		})
	}

	possibletimesMult := 1

	for _, race := range(races) {
		possibleWins := 0
		for i := 1; i < race.time - 1; i++ {
			if calculateDistance(i, race.time) > race.distance {
				possibleWins ++
			}
		}
		possibletimesMult *= possibleWins
	}

	return fmt.Sprint(possibletimesMult)
}

func Part2(input string) string {
	lines := util.ListOfLines(input)

	time := util.GetFirstIntInString(strings.ReplaceAll(lines[0], " ", ""))
	distance := util.GetFirstIntInString(strings.ReplaceAll(lines[1], " ", ""))

	possibleWins := 0

	for i := 1; i < time; i++ {
		if calculateDistance(i, time) > distance {
			possibleWins ++
		} else if possibleWins > 1 {
			break
		}
	}

	return fmt.Sprint(possibleWins)
}
