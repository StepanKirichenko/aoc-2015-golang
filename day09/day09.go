package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("day09/input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(string(f))
	lines := strings.Split(input, "\n")
	paths, locCount := parseInput(lines)
	fmt.Println("Parth One:", partOne(paths, locCount))
	fmt.Println("Parth Two:", partTwo(paths, locCount))
}

type pair struct {
	a, b int
}

func partOne(paths map[pair]int, locCount int) int {
    route := make([]int, locCount)
    for i := range route {
        route[i] = i
    }
	min := calcPathLength(route, paths)
    
    updateMin := func (route []int) {
        dist := calcPathLength(route, paths)
        if dist < min {
            min = dist
        }
    }

    forEachPermutation(route, updateMin)

    return min
}

func partTwo(paths map[pair]int, locCount int) int {
    route := make([]int, locCount)
    for i := range route {
        route[i] = i
    }
    max := 0
    
    updateMin := func (route []int) {
        dist := calcPathLength(route, paths)
        if dist > max {
            max = dist
        }
    }

    forEachPermutation(route, updateMin)

    return max
}

func calcPathLength(route []int, paths map[pair]int) int {
	dist := 0
	lenght := len(route)
	for i := 0; i < lenght-1; i++ {
		dist += paths[pair{route[i], route[i+1]}]
	}
    return dist
}

func forEachPermutation(arr []int, fn func([]int)) {
    n := len(arr)
    c := make([]int, n)
    fn(arr)
    i := 1
    for i < n {
        if c[i] < i {
            if i%2 == 0 {
                arr[0], arr[i] = arr[i], arr[0]
            } else {
                arr[c[i]], arr[i] = arr[i], arr[c[i]]
            }
            c[i] += 1
            i = 1
            fn(arr)
        } else {
            c[i] = 0
            i += 1
        }
    }
}

func parseInput(lines []string) (map[pair]int, int) {
	paths := make(map[pair]int)
	locs := make(map[string]int)
	locCount := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		from, to := fields[0], fields[2]
		a, exists := locs[from]
		if !exists {
			a = locCount
			locCount += 1
			locs[from] = a
		}
		b, exists := locs[to]
		if !exists {
			b = locCount
			locCount += 1
			locs[to] = b
		}
		dist, _ := strconv.Atoi(fields[4])
		paths[pair{a, b}] = dist
		paths[pair{b, a}] = dist
	}
	return paths, locCount
}
