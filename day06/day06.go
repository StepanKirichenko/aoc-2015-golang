package main

import (
	"fmt"
    "os"
	"strconv"
	"strings"
)

func main() {
    f, err := os.ReadFile("day06/input.txt")
    if err != nil {
        panic(err)
    }
    input := string(f)
    input = strings.TrimSpace(input)
    lines := strings.Split(input, "\n")
    fmt.Println("Part One:", partOne(lines))
    fmt.Println("Part Two:", partTwo(lines))
}

const (
	ON     int = 0
	OFF        = 1
	TOGGLE     = 2
)

type point struct {
    x, y int
}

func partOne(lines []string) int {
    var lights [1000][1000]bool
    for _, line := range lines {
        command, start, end := parseLine(line)
        for y := start.y; y <= end.y; y++ {
            for x := start.x; x <= end.x; x++ {
                switch command {
                case ON:
                    lights[y][x] = true
                case OFF:
                    lights[y][x] = false
                case TOGGLE:
                    lights[y][x] = !lights[y][x]
                }
            }
        }
    }

    count := 0
    for y := 0; y < 1000; y++ {
        for x := 0; x < 1000; x++ {
            if lights[y][x] {
                count += 1
            }
        }
    }
    return count
}

func partTwo(lines []string) int {
    var lights [1000][1000]int
    for _, line := range lines {
        command, start, end := parseLine(line)
        for y := start.y; y <= end.y; y++ {
            for x := start.x; x <= end.x; x++ {
                switch command {
                case ON:
                    lights[y][x] += 1
                case OFF:
                    lights[y][x] -= 1
                    if (lights[y][x] < 0) {
                        lights[y][x] = 0
                    }
                case TOGGLE:
                    lights[y][x] += 2
                }
            }
        }
    }

    brightness := 0
    for y := 0; y < 1000; y++ {
        for x := 0; x < 1000; x++ {
            brightness += lights[y][x]
        }
    }
    return brightness
}

func parseLine(line string) (command int, start, end point) {
    fields := strings.Fields(line)
    var args []string
    if fields[0] == "toggle" {
        command = TOGGLE
        args = fields[1:]
    } else {
        if fields[1] == "on" {
            command = ON
        } else {
            command = OFF
        }
        args = fields[2:]
    }
    startCoords := args[0]
    endCoords := args[2]
    start = parseCoords(startCoords)
    end = parseCoords(endCoords)
    return command, start, end
}

func parseCoords(coords string) point {
    strs := strings.Split(coords, ",")
    x, _ := strconv.Atoi(strs[0])
    y, _ := strconv.Atoi(strs[1])
    return point{x, y}
}
