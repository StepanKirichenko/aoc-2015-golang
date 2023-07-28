package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
    "bytes"
)

func main() {
	input, err := os.ReadFile("day04/input.txt")
	if err != nil {
		panic(err)
	}
    key := bytes.TrimSpace(input)
	fmt.Println("Part One:", partOne(key))
	fmt.Println("Part Two:", partTwo(key))
}

func partOne(key []byte) int {
    return findNumberFunc(key, partOnePredicate)
}

func partTwo(key []byte) int {
    return findNumberFunc(key, partTwoPredicate)
}

func partOnePredicate(s string) bool {
    return strings.HasPrefix(s, "00000")
}

func partTwoPredicate(s string) bool {
    return strings.HasPrefix(s, "000000")
}

func findNumberFunc(key []byte, predicate func (string) bool) int {
	md5Hash := md5.New()
	for num := 0; ; num += 1 {
		md5Hash.Reset()
		from := fmt.Sprintf("%s%d", key, num)
		md5Hash.Write([]byte(from))
		sum := md5Hash.Sum(nil)
		sumString := fmt.Sprintf("%x", sum)
        if predicate(sumString) {
            return num
        }
	}
}
