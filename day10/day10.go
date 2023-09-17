package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day10/input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(file))
	partOneResult := solvePartOne(input)
	partTwoResult := solvePartTwo(input)
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Part one: %d\n", partOneResult)
	fmt.Printf("Part two: %d\n", partTwoResult)
}

func transformString(input string) string {
	result := make([]string, 0)

	var prevRune rune = 0 
	count := 0
	for _, c := range input {
		if c != prevRune {
			if prevRune != 0 {
				result = append(result, strconv.Itoa(count))
				result = append(result, string(prevRune))
			}
			count = 0
		}
		prevRune = c
		count += 1
	}

	result = append(result, strconv.Itoa(count))
	result = append(result, string(prevRune))

	return strings.Join(result, "")
}

func countLengthAfterNTransformations(input string, n int) int {
    for i := 0; i < n; i += 1 {
        input = transformString(input)
    }
    return len(input)
}

func solvePartOne(input string) int {
    return countLengthAfterNTransformations(input, 40)
}

func solvePartTwo(input string) int {
    return countLengthAfterNTransformations(input, 50)
}
