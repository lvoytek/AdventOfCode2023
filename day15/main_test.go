package main

import (
	"testing"
)

var part1Example = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
var part1Expected = `1320`

var part2Example = part1Example
var part2Expected = `145`

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
