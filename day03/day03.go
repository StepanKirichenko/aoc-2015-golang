package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("day03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}

const (
	UP    = byte('^')
	DOWN  = byte('v')
	LEFT  = byte('<')
	RIGHT = byte('>')
)

type house struct {
	x, y int
}

func partOne(input []byte) int {
	visited := make(map[house]int)
	visited[house{0, 0}] = 1

	var curHouse house
    for _, dir := range input {
		switch dir {
		case UP:
			curHouse.y += 1
		case DOWN:
			curHouse.y -= 1
		case LEFT:
			curHouse.x -= 1
		case RIGHT:
			curHouse.x += 1
		}
		visited[curHouse] += 1
	}

	return len(visited)
}

func partTwo(input []byte) int {
	visited := make(map[house]int)
	visited[house{0, 0}] = 2

	curHouses := []house{{0, 0}, {0, 0}}
	curIndex := 0
    for _, dir := range input {
		curHouse := curHouses[curIndex]
		switch dir {
		case UP:
			curHouse.y += 1
		case DOWN:
			curHouse.y -= 1
		case LEFT:
			curHouse.x -= 1
		case RIGHT:
			curHouse.x += 1
		}
		visited[curHouse] += 1
        curHouses[curIndex] = curHouse
        curIndex = 1 - curIndex
	}

	return len(visited)
}
