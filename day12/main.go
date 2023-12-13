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

type DamageMap struct {
	damageMap string
	damageSizes []int
}

func extractDamage(line string) DamageMap {
	splitLine := strings.Split(line, " ")
	return DamageMap {
		damageMap: splitLine[0],
		damageSizes: util.GetAllIntsInString(splitLine[1]),
	}
}

func getDamageMapRegexString(damageMap DamageMap) string {
	damageRegexStr := `^[\.\?]*`

	for i, item := range(damageMap.damageSizes) {
		damageRegexStr += `#{` + fmt.Sprint(item) + "}"

		if i < len(damageMap.damageSizes) - 1 {
			damageRegexStr += `[\.\?]+`
		} else {
			damageRegexStr += `[\.\?]*$`
		}
	}

	return damageRegexStr
}

func findMatchesForString(damageMap []rune, re *regexp.Regexp) int {

	if re.MatchString(string(damageMap)) {
		return 1
	}

	totalMatches := 0

	for i, item := range(damageMap) {
		if item == '?' {
			damageMap[i] = '#'
			totalMatches += findMatchesForString(damageMap, re)
			damageMap[i] = '.'
			totalMatches += findMatchesForString(damageMap, re)
			damageMap[i] = '?'
			break
		}
	}

	return totalMatches
}

func findNumMatchesRegex(damageMap DamageMap) int {
	return findMatchesForString([]rune(damageMap.damageMap), regexp.MustCompile(getDamageMapRegexString(damageMap)))
}

func countSplitMatches(splitDamages []string, damageSizes []int) int {
	// All remaining ? and . are working, therefore all . is only solution
	if len(damageSizes) == 0 {
		return 1
	}

	// Nowhere to put remaining damage, no solution
	if len(splitDamages) == 0 {
		return 0
	}

	nextSplit := splitDamages[0]

	// Upcoming damage size is too big for next damage, skip or die
	if damageSizes[0] > len(nextSplit) {
		if strings.ContainsRune(nextSplit, '#') {
			return 0
		} else {
			return countSplitMatches(splitDamages[1:], damageSizes)
		}
	}

	// Damage size matches damage exactly, skip past first of both
	if damageSizes[0] == len(nextSplit) && strings.ContainsRune(nextSplit, '#') {
		return countSplitMatches(splitDamages[1:], damageSizes[1:])
	}


	// First character is #, therefore the next damage amount must be matched
	if nextSplit[0] == '#' {
		// Next item must be a ? to split next damage
		if nextSplit[damageSizes[0]] != '?' {
			return 0
		}

		splitDamages[0] = nextSplit[damageSizes[0] + 1:]
		splitCount := countSplitMatches(splitDamages, damageSizes[1:])
		splitDamages[0] = nextSplit
		return splitCount

	// First character is ?, recurse with it as a . and a #
	} else {
		splitDamages[0] = nextSplit[1:]
		totalCount := countSplitMatches(splitDamages, damageSizes)
		splitDamages[0] = "#" + nextSplit[1:]
		totalCount += countSplitMatches(splitDamages, damageSizes)
		splitDamages[0] = nextSplit

		return totalCount
	}
}

func findNumMatchesSplit(damageMap DamageMap) int64 {
	splitDamages := strings.Split(damageMap.damageMap, ".")
	var cleanSplitDamages []string

	for _, splitDamage := range(splitDamages) {
		if len(splitDamage) > 0 {
			cleanSplitDamages = append(cleanSplitDamages, splitDamage)
		}
	}

	return int64(countSplitMatches(cleanSplitDamages, damageMap.damageSizes))
}

func unfoldMap(damageMap DamageMap) DamageMap {
	newDamageMap := damageMap.damageMap
	newDamageSizes := damageMap.damageSizes

	for i := 0; i < 4; i++ {
		newDamageMap += "?" + damageMap.damageMap
		newDamageSizes = append(newDamageSizes, damageMap.damageSizes...)
	}

	return DamageMap {
		damageMap: newDamageMap,
		damageSizes: newDamageSizes,
	}
}

func Part1(input string) string {
	lines := util.ListOfLines(input)

	totalMatches := 0

	for _, line := range(lines) {
		damageMap := extractDamage(line)
		totalMatches += findNumMatchesRegex(damageMap)
	}

	return fmt.Sprint(totalMatches)
}

func Part2(input string) string {
	lines := util.ListOfLines(input)

	var totalMatches int64 = 0

	for _, line := range(lines) {
		damageMap := extractDamage(line)
		damageMap = unfoldMap(damageMap)
		totalMatches += findNumMatchesSplit(damageMap)
	}

	return fmt.Sprint(totalMatches)
}
