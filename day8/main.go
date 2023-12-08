package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"

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

type Node struct {
	name string
	left string
	right string
}

func extractNode(line string) Node {
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	return Node {
		name: matches[0][1],
		left: matches[0][2],
		right: matches[0][3],
	}
}

func Part1(input string) string {
	lines := util.ListOfLineChunks(input)
	instructions := lines[0][0]

	nodes := make(map[string]Node)

	for _, line := range(lines[1]) {
		newNode := extractNode(line)
		nodes[newNode.name] = newNode
	}

	instructionIndex := 0
	currentNode := "AAA"
	instructionCount := 0

	for {
		if currentNode == "ZZZ" {
			break
		}

		if instructions[instructionIndex] == 'L' {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}

		instructionIndex ++
		instructionCount ++

		if instructionIndex >= len(instructions) {
			instructionIndex = 0
		}
	}

	return fmt.Sprint(instructionCount)
}

func Part2(input string) string {
	lines := util.ListOfLineChunks(input)
	instructions := lines[0][0]

	nodes := make(map[string]Node)

	for _, line := range(lines[1]) {
		newNode := extractNode(line)
		nodes[newNode.name] = newNode
	}

	var currentNodes []string
	for k, _ := range(nodes) {
		if k[len(k) - 1] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}

	var instructionCounts []int

	for _, currentNode := range(currentNodes) {
		instructionIndex := 0
		instructionCount := 0

		for {
			if currentNode[len(currentNode) - 1] == 'Z' {
				break
			}

			if instructions[instructionIndex] == 'L' {
				currentNode = nodes[currentNode].left
			} else {
				currentNode = nodes[currentNode].right
			}

			instructionIndex ++
			instructionCount ++

			if instructionIndex >= len(instructions) {
				instructionIndex = 0
			}
		}

		instructionCounts = append(instructionCounts, instructionCount)
	}

	minInstructionCount := 1

	for _, instructionCount := range(instructionCounts) {
		minInstructionCount = util.Lcm(minInstructionCount, instructionCount)
	}

	return fmt.Sprint(minInstructionCount)
}
