package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.ReadFile("day01/input.txt")
	if err != nil {
		panic(err)
	}

    fmt.Println("Part One:", partOne(f))
    fmt.Println("Part Two:", partTwo(f))
}

const (
	UP   = byte('(')
	DOWN = byte(')')
)

func partOne(input []byte) int {
	floor := 0

	for _, dir := range input {
		switch dir {
		case UP:
			floor += 1
		case DOWN:
			floor -= 1
		}
	}

	return floor
}

func partTwo(input []byte) int {
    floor := 0

    for i, dir := range input {
        switch dir {
        case UP:
            floor += 1
        case DOWN:
            floor -= 1
            if floor < 0 {
                return i + 1
            }
        }
    }

    return 0
}
