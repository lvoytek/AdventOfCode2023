package main

import (
	"testing"
)

var part1Example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`
var part1Expected = `136`

var part2Example = part1Example
var part2Expected = `64`

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
