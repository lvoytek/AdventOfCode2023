package util

import (
	"regexp"
	"strconv"
)

/* Get the first positive integer in a string*/
func GetFirstIntInString(input string) int {
	intRe := regexp.MustCompile(`(-?\d+)`)
	result, _ := strconv.Atoi(intRe.FindString(input))
	return result
}

/* Get all positive integers in a string */
func GetAllIntsInString(input string) []int {
	intRe := regexp.MustCompile(`(-?\d+)`)
	resultStrings := intRe.FindAllString(input, -1)

	var result []int

	for _, resultString := range(resultStrings) {
		nextResult, _ := strconv.Atoi(resultString)
		result = append(result, nextResult)
	}

	return result
}