package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"math"

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

type AlmanacConversion struct {
	sourceStart int
	destStart int
	convertRange int
}

type AlmanacMap struct {
	sourceName string
	destName string
	conversions []AlmanacConversion
}

type SeedRange struct {
	startSeed int
	seedRange int
}

func extractMap(lines []string) AlmanacMap {
	re := regexp.MustCompile(`(\b\w+\b)-to-(\b\w+\b)`)
	match := re.FindStringSubmatch(lines[0])

	var conversions []AlmanacConversion

	for _, line := range(lines[1:]) {
		vals := util.GetAllIntsInString(line)
		conversions = append(conversions, AlmanacConversion{
			destStart: vals[0],
			sourceStart: vals[1],
			convertRange: vals[2],
		})
	}

	return AlmanacMap {
		sourceName: match[1],
		destName: match[2],
		conversions: conversions,
	}

}

func expandRange(seedRange SeedRange, conversions []AlmanacConversion) []SeedRange {
	if len(conversions) == 0 {
		return []SeedRange { seedRange }
	}

	conversion := conversions[0]

	if conversion.sourceStart <= seedRange.startSeed + seedRange.seedRange &&
	seedRange.startSeed <= conversion.sourceStart + conversion.convertRange {

		convertedStart := util.MaxInt(conversion.sourceStart, seedRange.startSeed)
		convertedRange := util.MinInt(conversion.sourceStart + conversion.convertRange, seedRange.startSeed + seedRange.seedRange) - convertedStart

		returnRanges := []SeedRange{ SeedRange {
			startSeed: convertedStart + conversion.destStart - conversion.sourceStart,
			seedRange: convertedRange,
		}}

		if seedRange.startSeed < convertedStart {
			returnRanges = append(returnRanges, expandRange(SeedRange {
				startSeed: seedRange.startSeed,
				seedRange: convertedStart - seedRange.startSeed,
			}, conversions[1:])...)
		}

		if seedRange.startSeed + seedRange.seedRange > convertedStart + convertedRange {
			returnRanges = append(returnRanges, expandRange(SeedRange {
				startSeed: convertedStart + convertedRange,
				seedRange: seedRange.startSeed + seedRange.seedRange - convertedStart - convertedRange,
			}, conversions[1:])...)
		}

		return returnRanges
	} else {
		return expandRange(seedRange, conversions[1:])
	}
}

func Part1(input string) string {
	lineChunks := util.ListOfLineChunks(input)
	seeds := util.GetAllIntsInString(lineChunks[0][0])

	var maps []AlmanacMap

	for _, chunk := range(lineChunks[1:]) {
		maps = append(maps, extractMap(chunk))
	}

	lowestLocation := math.MaxInt

	for _, seed := range(seeds) {
		for _, almanacMap := range(maps) {
			for _, conversion := range(almanacMap.conversions) {
				if seed >= conversion.sourceStart && seed < conversion.sourceStart + conversion.convertRange {
					seed += conversion.destStart - conversion.sourceStart
					break
				}
			}
		}

		if seed < lowestLocation {
			lowestLocation = seed
		}
	}

	return fmt.Sprint(lowestLocation)
}

func Part2(input string) string {
	lineChunks := util.ListOfLineChunks(input)
	var seedRanges []SeedRange
	seedInfo := util.GetAllIntsInString(lineChunks[0][0])

	for i := 0; i < len(seedInfo) -1; i += 2 {
		seedRanges = append(seedRanges, SeedRange {
			startSeed: seedInfo[i],
			seedRange: seedInfo[i+1],
		})
	}

	var maps []AlmanacMap

	for _, chunk := range(lineChunks[1:]) {
		maps = append(maps, extractMap(chunk))
	}

	for _, almanacMap := range(maps) {
		var newSeedRanges []SeedRange
		for _, seedRange := range(seedRanges) {
			newSeedRanges = append(newSeedRanges, expandRange(seedRange, almanacMap.conversions)...)
		}
		seedRanges = newSeedRanges

		var noDupRanges []SeedRange
		for _, seedRange := range(seedRanges) {
			isDuplicate := false
			for noDupIndex, noDupRange := range(noDupRanges) {
				if seedRange.startSeed == noDupRange.startSeed {
					isDuplicate = true
					noDupRanges[noDupIndex].seedRange = util.MaxInt(seedRange.seedRange, noDupRange.seedRange)
					break
				}
			}

			if !isDuplicate {
				noDupRanges = append(noDupRanges, seedRange)
			}
		}
		seedRanges = noDupRanges
	}

	lowestLocation := math.MaxInt

	for _, seedRange := range(seedRanges) {
		if seedRange.startSeed < lowestLocation {
			lowestLocation = seedRange.startSeed
		}
	}

	return fmt.Sprint(lowestLocation)
}
