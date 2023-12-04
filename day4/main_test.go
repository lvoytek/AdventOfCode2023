package main

import (
	"testing"
)

var part1Example = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
var part1Expected = `13`

var part2Example = part1Example
var part2Expected = `30`

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
