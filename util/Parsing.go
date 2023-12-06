package util

import (
	"strings"
)

/* Return file contents as a list of strings containing each line */
func ListOfLines(input string) (lines []string) {
	for _, line := range(strings.Split(input, "\n")) {
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines
}

/* Return a list of lists of lines split by empty space lines */
func ListOfLineChunks(input string) (chunks [][]string) {
	var nextChunk []string

	for _, line := range(strings.Split(input, "\n")) {
		if len(line) > 0 {
			nextChunk = append(nextChunk, line)
		} else {
			if len(nextChunk) > 0 {
				chunks = append(chunks, nextChunk)
				nextChunk = []string{}
			}
		}
	}

	return chunks
}

/* Return file contents as a matrix of runes */
func MatrixOfCharacters(input string) (matrix [][]rune) {
	for _, line := range(strings.Split(input, "\n")) {
		if len(line) > 0 {
			characters := []rune(line)
			matrix = append(matrix, characters)
		}
	}

	return matrix
}

/* Return file contents as a matrix of numeric digits with non-digits replaced */
func MatrixOfDigits(input string, nonDigits int) (matrix [][]int) {
	for _, line := range(strings.Split(input, "\n")) {
		if len(line) == 0 {
			continue
		}

		var lineOfDigits []int
		for _, character := range(line) {
			nextDigit := int(character) - int('0')
			if nextDigit < 0 || nextDigit > 9 {
				nextDigit = nonDigits
			}

			lineOfDigits = append(lineOfDigits, nextDigit)
		}

		matrix = append(matrix, lineOfDigits)
	}

	return matrix
}