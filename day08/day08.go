package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	f, err := os.ReadFile("day08/input.txt")
	if err != nil {
		panic(err)
	}
	f = bytes.TrimSpace(f)
	lines := bytes.Split(f, []byte("\n"))
	fmt.Println("Part One:", partOne(lines))
	fmt.Println("Part Two:", partTwo(lines))
}

func partOne(lines [][]byte) int {
	totalDiff := 0
	for _, line := range lines {
		totalDiff += diffInMemory(line)
	}
	return totalDiff
}

func partTwo(lines [][]byte) int {
	totalDiff := 0
	for _, line := range lines {
		totalDiff += diffEncoded(line)
	}
	return totalDiff
}

const (
	BACKSLASH byte = byte('\\')
	X              = byte('x')
	QUOTE          = byte('"')
)

func diffInMemory(literal []byte) int {
	length := len(literal)
	inMemory := 0
	for i := 0; i < length; i++ {
		if literal[i] == BACKSLASH {
			i += 1
			if literal[i] == X {
				i += 2
			}
		}
		inMemory += 1
	}
	inMemory -= 2
	return length - inMemory
}

func diffEncoded(literal []byte) int {
	lenght := len(literal)
	encoded := 0
	for _, char := range literal {
        encoded += 1
        if char == BACKSLASH || char == QUOTE {
            encoded += 1
        }
	}
    encoded += 2
    return encoded - lenght
}
