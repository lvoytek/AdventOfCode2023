package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

type Card struct {
	number int
	multiplier int
	winning []int
	yours []int
}

func extractCard(line string) Card {
	cardNumber := util.GetFirstIntInString(line)

	winningInts := util.GetAllIntsInString(line[strings.Index(line, ":"):strings.Index(line, "|")])
	yourInts := util.GetAllIntsInString(line[strings.Index(line, "|"):])

	return Card {
		number: cardNumber,
		multiplier: 1,
		winning: winningInts,
		yours: yourInts,
	}
}

func getNumMatches(card Card) int {
	numMatches := 0

	for _, yours := range(card.yours) {
		hasMatch := false
		for _, winning := range(card.winning) {
			if yours == winning {
				hasMatch = true
				break
			}
		}

		if hasMatch {
			numMatches ++
		}
	}

	return numMatches
}

func Part1(input string) string {
	lines := util.ListOfLines(input)

	sumOfWinnings := 0

	for _, line := range(lines) {
		nextCard := extractCard(line)
		numMatches := getNumMatches(nextCard)

		if numMatches > 0 {
			sumOfWinnings += int(math.Pow(2, float64(numMatches -1)))
		}
	}

	return fmt.Sprint(sumOfWinnings)
}

func Part2(input string) string {
	lines := util.ListOfLines(input)
	var cards []Card

	for _, line := range(lines) {
		cards = append(cards, extractCard(line))
	}

	totalCards := 0

	for cardIndex, card := range(cards) {
		numMatches := getNumMatches(card)

		for i := cardIndex + 1; i <= cardIndex + numMatches && i < len(cards); i++ {
			cards[i].multiplier += card.multiplier
		}

		totalCards += card.multiplier
	}

	return fmt.Sprint(totalCards)
}
