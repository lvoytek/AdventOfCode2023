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

// Standard format for a function run as a check in a workflow
type WorkflowCheckFunc func(map[rune]int, rune, int) bool

/* Less than workflow check function */
func lt(rating map[rune]int, category rune, val int) bool {
	return rating[category] < val
}

/* Greater than workflow check function */
func gt(rating map[rune]int, category rune, val int) bool {
	return rating[category] > val
}

/* Always pass workflow check function */
func pass(map[rune]int, rune, int) bool {
	return true
}

// Map workflow check characters to workflow check functions
var CheckFunctions = map[rune]WorkflowCheckFunc {
	'<':lt,
	'>':gt,
	'p':pass,
}

// A single check in a workflow and the workflow name to go to on success
type WorkflowCheck struct {
	checkType rune
	checkVar rune
	checkVal int
	result string
}

// A range of integers (inclusive)
type Range struct {
	min int
	max int
}

// x, m, a, and s rating ranges and the next workflow to test with
type RatingRange struct {
	ranges map[rune]Range
	nextWorkflow string
}

/* Extract all workflows from text input as a map of names to arrays of checks */
func extractWorkflows(lines []string) map[string][]WorkflowCheck {
	workflows := make(map[string][]WorkflowCheck)
	workflowRe := regexp.MustCompile(`(\w)([><])(\d+):(\w+)|(\w+)`)

	for _, workflowString := range(lines) {
		splitString := strings.Split(workflowString, "{")
		workflowName := splitString[0]
		checkStrings := strings.Split(splitString[1][:len(splitString[1])-1], ",")

		var checks []WorkflowCheck

		for _, checkString := range(checkStrings) {
			checkParts := workflowRe.FindStringSubmatch(checkString)[1:]

			checkResult := checkParts[len(checkParts) - 1]
			checkType := 'p'
			checkVal := 0
			checkVar := 'x'

			if len(checkParts[0]) > 0 {
				checkVal = util.GetFirstIntInString(checkParts[2])
				checkVar = rune(checkParts[0][0])
				checkType = rune(checkParts[1][0])
				checkResult = checkParts[3]
			}

			checks = append(checks, WorkflowCheck {
				checkType: checkType,
				checkVar: checkVar,
				checkVal: checkVal,
				result: checkResult,
			})
		}

		workflows[workflowName] = checks
	}

	return workflows
}

/* Extract x, m, a, and s rating data from the input */
func extractRatings(lines []string) (ratings []map[rune]int) {
	for _, line := range(lines) {
		vals := util.GetAllIntsInString(line)

		ratings = append(ratings, map[rune]int{
			'x': vals[0],
			'm': vals[1],
			'a': vals[2],
			's': vals[3],
		})
	}

	return ratings
}

/* Run through the workflows for each rating and return all that pass */
func getAcceptedRatings(ratings []map[rune]int, workflows map[string][]WorkflowCheck) []map[rune]int {
	var acceptedRatings []map[rune]int

	for _, rating := range(ratings) {
		currentWorkflow := workflows["in"]

		for {
			foundPassFail := false
			for _, check := range(currentWorkflow) {
				if CheckFunctions[check.checkType](rating, check.checkVar, check.checkVal) {
					if check.result == "A" {
						acceptedRatings = append(acceptedRatings, rating)
						foundPassFail = true
					} else if check.result == "R" {
						foundPassFail = true
					} else {
						currentWorkflow = workflows[check.result]
					}

					break
				}
			}

			if foundPassFail {
				break
			}
		}
	}

	return acceptedRatings
}

/* Find the number of combinations of accepted ratings based on minimum and maximum rating values */
func findPossibleAcceptances(workflows map[string][]WorkflowCheck, min int, max int) int64{
	ratingRanges := []RatingRange { RatingRange {
			ranges: map[rune]Range {
				'x': {min, max},
				'm': {min, max},
				'a': {min, max},
				's': {min, max},
			},
			nextWorkflow: "in",
		},
	}

	var acceptedRanges []RatingRange

	for len(ratingRanges) > 0 {
		currentRatings := ratingRanges[0]
		ratingRanges = ratingRanges[1:]

		if currentRatings.nextWorkflow == "A" {
			acceptedRanges = append(acceptedRanges, currentRatings)
			continue
		}

		if currentRatings.nextWorkflow == "R" {
			continue
		}

		currentWorkflow := workflows[currentRatings.nextWorkflow]

		for _, check := range(currentWorkflow) {
			if check.checkType == 'p' {
				currentRatings.nextWorkflow = check.result
				ratingRanges = append(ratingRanges, currentRatings)
				break
			} else if check.checkType == '>' {
				if currentRatings.ranges[check.checkVar].min > check.checkVal {
					currentRatings.nextWorkflow = check.result
					ratingRanges = append(ratingRanges, currentRatings)
					break
				} else if currentRatings.ranges[check.checkVar].max > check.checkVal {
					prevMax := currentRatings.ranges[check.checkVar].max
					currentRatings.ranges[check.checkVar] = Range{currentRatings.ranges[check.checkVar].min, check.checkVal}

					newRatingsRanges := make(map[rune]Range)
					for k, v := range currentRatings.ranges {
						if k == check.checkVar {
							newRatingsRanges[k] = Range{check.checkVal + 1, prevMax}
						} else {
							newRatingsRanges[k] = v
						}
					}

					ratingRanges = append(ratingRanges, RatingRange {
						ranges: newRatingsRanges,
						nextWorkflow: check.result,
					})
				}
			} else {
				if currentRatings.ranges[check.checkVar].max < check.checkVal {
					currentRatings.nextWorkflow = check.result
					ratingRanges = append(ratingRanges, currentRatings)
					break
				} else if currentRatings.ranges[check.checkVar].min < check.checkVal {
					prevMin := currentRatings.ranges[check.checkVar].min
					currentRatings.ranges[check.checkVar] = Range{check.checkVal, currentRatings.ranges[check.checkVar].max}

					newRatingsRanges := make(map[rune]Range)
					for k, v := range currentRatings.ranges {
						if k == check.checkVar {
							newRatingsRanges[k] = Range{prevMin, check.checkVal - 1}
						} else {
							newRatingsRanges[k] = v
						}
					}

					ratingRanges = append(ratingRanges, RatingRange {
						ranges: newRatingsRanges,
						nextWorkflow: check.result,
					})
				}
			}
		}
	}

	totalPossible := int64(0)

	for _, ratingRanges := range(acceptedRanges) {
		nextPossible := int64(1)

		for _, ratingRange := range(ratingRanges.ranges) {
			nextPossible *= int64(ratingRange.max-ratingRange.min+1)
		}

		totalPossible += nextPossible
	}

	return totalPossible
}

/* Find the total of all x, m, a, and s ratings for all accepted ratings combined */
func Part1(input string) string {
	sections := util.ListOfLineChunks(input)
	workflowStrings := sections[0]
	ratingsStrings := sections[1]

	workflows := extractWorkflows(workflowStrings)
	ratings := extractRatings(ratingsStrings)

	acceptedRatings := getAcceptedRatings(ratings, workflows)

	acceptedTotal := 0

	for _, rating := range(acceptedRatings) {
		for _, item := range(rating) {
			acceptedTotal += item
		}
	}

	return fmt.Sprint(acceptedTotal)
}

/* Find the number of possible accepted rating combinations where ratings can be 1 to 4000 */
func Part2(input string) string {
	workflowStrings := util.ListOfLineChunks(input)[0]
	workflows := extractWorkflows(workflowStrings)
	return fmt.Sprint(findPossibleAcceptances(workflows, 1, 4000))
}
