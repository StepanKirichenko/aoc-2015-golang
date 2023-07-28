package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

type box struct {
	l, w, h int
}

func main() {
	f, err := os.OpenFile("day02/input.txt", os.O_RDONLY, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	boxes := parseInput(scanner)

    fmt.Println("Part One:", partOne(boxes))
    fmt.Println("Part Two:", partTwo(boxes))
}

func parseInput(scanner *bufio.Scanner) []box {
	boxes := make([]box, 0)

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(vals[0])
		w, _ := strconv.Atoi(vals[1])
		h, _ := strconv.Atoi(vals[2])
		b := box{l, w, h}
		boxes = append(boxes, b)
	}

	return boxes
}

func calc(boxes []box, formula func(box) int) int {
    sum := 0
    for _, b := range boxes {
        sum += formula(b)
    }
    return sum
}

func partOneFormula(b box) int {
	side1 := b.l * b.w
	side2 := b.w * b.h
	side3 := b.h * b.l
    minSide := min(side1, side2, side3)
    return 2 * (side1 + side2 + side3) + minSide
}

func partOne(boxes []box) int {
    return calc(boxes, partOneFormula)
}

func partTwoFormula(b box) int {
	side1 := b.l + b.w
	side2 := b.w + b.h
	side3 := b.h + b.l
    minSide := min(side1, side2, side3)
    volume := b.l * b.w * b.h
    return 2 * minSide + volume
}

func partTwo(boxes []box) int {
    return calc(boxes, partTwoFormula)
}

func min(args ...int) int {
    if len(args) == 0 {
        return 0
    }
    m := args[0]
    for _, arg := range args {
        if arg < m {
            m = arg
        }
    }
    return m
}
