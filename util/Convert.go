package util

import (
	"regexp"
	"strconv"
)

/* Get the first positive integer in a string*/
func GetFirstIntInString(input string) int {
	intRe := regexp.MustCompile(`(\d+)`)
	result, _ := strconv.Atoi(intRe.FindString(input))
	return result
}