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

// A node with a given name and its left and right connections
type Node struct {
	name string
	left string
	right string
}

/* Extract a node and its connections from a line of input */
func extractNode(line string) Node {
	re := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	return Node {
		name: matches[0][1],
		left: matches[0][2],
		right: matches[0][3],
	}
}

/* Find # of steps to get from AAA to ZZZ with given instructions */
func Part1(input string) string {
	lines := util.ListOfLineChunks(input)
	instructions := lines[0][0]

	// Create a map of node names to nodes in data set
	nodes := make(map[string]Node)

	for _, line := range(lines[1]) {
		newNode := extractNode(line)
		nodes[newNode.name] = newNode
	}

	// Count instructions while following path from AAA to ZZZ
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

/* Find # of steps to get from all **A nodes to **Z's all at the same time */
func Part2(input string) string {
	lines := util.ListOfLineChunks(input)
	instructions := lines[0][0]

	// Create a map of node names to nodes in data set
	nodes := make(map[string]Node)

	for _, line := range(lines[1]) {
		newNode := extractNode(line)
		nodes[newNode.name] = newNode
	}

	// Find all starting nodes
	var currentNodes []string
	for k, _ := range(nodes) {
		if k[len(k) - 1] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}

	// Find min number of instructions to get from **A to **Z for each start
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

	// Find when all nodes are at **Z using least common multiple
	minInstructionCount := 1

	for _, instructionCount := range(instructionCounts) {
		minInstructionCount = util.Lcm(minInstructionCount, instructionCount)
	}

	return fmt.Sprint(minInstructionCount)
}
