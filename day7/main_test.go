package main

import (
	"testing"
)

var part1Example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
var part1Expected = `6440`

var part2Example = part1Example
var part2Expected = `5905`

func Test_Part1(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		actual := Part1(part1Example)
		if actual != part1Expected {
			t.Errorf("Expected: %v, Actual: %v", part1Expected, actual)
		}
	})
}

func Test_Part2(t *testing.T) {
	t.Run("Part 2", func(t *testing.T) {
		actual := Part2(part2Example)
		if actual != part2Expected {
			t.Errorf("Expected: %v, Actual: %v", part2Expected, actual)
		}
	})
}
