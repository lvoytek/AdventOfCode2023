package util

import (
	"strings"
)

/* Return file contents as a list of strings containing each line */
func ListofLines(input string) (lines []string) {
	for _, line := range(strings.Split(input, "\n")) {
		lines = append(lines, line)
	}

	return lines
}

/* Return file contents as a matrix of runes */
func MatrixOfCharacters(input string) (matrix [][]rune) {
	for _, line := range(strings.Split(input, "\n")) {
		characters := []rune(line)
		matrix = append(matrix, characters)
	}

	return matrix
}