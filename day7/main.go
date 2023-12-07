package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

type Hand struct {
	cards string
	bid int
	strength int
}

func getCardValue(card rune, joker bool) int {
	if card == 'A' {
		return 14
	} else if card == 'K' {
		return 13
	} else if card == 'Q' {
		return 12
	} else if card == 'J' {
		if joker {
			return 1
		}

		return 11
	} else if card == 'T' {
		return 10
	}

	return int(card) - '0'
}

func getHandStrength(cards string, joker bool) int {
	cardCounts := make(map[rune]int)

	for _, card := range(cards) {
		cardCounts[card] ++
	}

	maxOccurance := 0
	var maxKey rune

	for key, cardCount := range(cardCounts) {
		if cardCount > maxOccurance && key != 'J' {
			maxOccurance = cardCount
			maxKey = key
		}
	}

	if joker {
		cardCounts[maxKey] += cardCounts['J']
		maxOccurance += cardCounts['J']
		cardCounts['J'] = 0
	}

	twos := 0

	for _, cardCount := range(cardCounts) {
		if cardCount == 2 {
			twos ++
		}
	}

	strength := maxOccurance

	if maxOccurance > 3 || (maxOccurance == 3 && twos > 0) {
		strength += 2
	} else if maxOccurance == 3 || twos > 1 {
		strength += 1
	}

	return strength
}

func runPart(input string, joker bool) string {
	lines := util.ListOfLines(input)

	var hands []Hand

	for _, line := range(lines) {
		splitLine := strings.Split(line, " ")
		strength := getHandStrength(splitLine[0], joker)
		bid := util.GetFirstIntInString(splitLine[1])

		hands = append(hands, Hand {
			cards: splitLine[0],
			bid: bid,
			strength: strength,
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength == hands[j].strength {
			for cardNum := 0; cardNum < len(hands[i].cards); cardNum ++ {
				if hands[i].cards[cardNum] != hands[j].cards[cardNum] {
					return getCardValue(([]rune (hands[i].cards))[cardNum], joker) < getCardValue(([]rune (hands[j].cards))[cardNum], joker)
				}
			}
		}

		return hands[i].strength < hands[j].strength
	})

	totalWinnings := 0

	for i, hand := range(hands) {
		totalWinnings += (i+1) * hand.bid
	}

	return fmt.Sprint(totalWinnings)
}

func Part1(input string) string {
	return runPart(input, false)
}

func Part2(input string) string {
	return runPart(input, true)
}
