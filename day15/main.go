package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

// A lens with a label and focal length
type Lens struct {
	label string
	focalLength int
	deleted bool
}

// A box containing an ordered set of lenses
type Box struct {
	boxNumber int
	lenses []Lens
}

/* Run the provided hashing algorithm of char * 17 mod 256 */
func runHash(stringToHash string) int {
	current := 0

	for _, character := range(stringToHash) {
		current += int(character)
		current *= 17
		current %= 256
	}

	return current
}

/* Find the sum of hashes for each comma separated input */
func Part1(input string) string {
	lines := util.ListOfLines(input)
	stringsToHash := strings.Split(lines[0], ",")

	hashSum := 0

	for _, item := range(stringsToHash) {
		hashSum += runHash(item)
	}

	return fmt.Sprint(hashSum)
}

/* Place lenses in designated boxes then find the total focal power */
func Part2(input string) string {
	lines := util.ListOfLines(input)
	stringsToHash := strings.Split(lines[0], ",")

	boxes := make(map[int]Box)
	equalRe := regexp.MustCompile(`(\w+)=\d+`)

	for _, item := range(stringsToHash) {
		if equalRe.MatchString(item) {
			newLensLabel := strings.Split(item, "=")[0]
			newHash := runHash(newLensLabel)
			newFocalLength := util.GetFirstIntInString(item)

			_, boxExists := boxes[newHash]

			if !boxExists {
				boxes[newHash] = Box {
					boxNumber: newHash,
				}
			}

			foundAndReplaced := false

			for i, lens := range(boxes[newHash].lenses) {
				if lens.label == newLensLabel && !lens.deleted {
					boxes[newHash].lenses[i].focalLength = newFocalLength
					foundAndReplaced = true
				}
			}

			if !foundAndReplaced {
				boxes[newHash] = Box {
					boxNumber: newHash,
					lenses: append(boxes[newHash].lenses, Lens {
						label: newLensLabel,
						focalLength: newFocalLength,
						deleted: false,
					}),
				}
			}
		} else {
			newLensLabel := item[:len(item)-1]
			newHash := runHash(newLensLabel)

			for i, lens := range(boxes[newHash].lenses) {
				if lens.label == newLensLabel && !lens.deleted {
					boxes[newHash].lenses[i].deleted = true
					break
				}
			}

		}
	}

	totalFocusPower := 0

	for _, box := range(boxes) {
		slot := 1
		for _, lens := range(box.lenses) {
			if !lens.deleted {
				totalFocusPower += (box.boxNumber + 1) * (slot) * lens.focalLength
				slot ++
			}
		}
	}

	return fmt.Sprint(totalFocusPower)
}
