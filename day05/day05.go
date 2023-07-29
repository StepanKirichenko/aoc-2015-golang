package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}
	input := string(f)
	lines := strings.Split(input, "\n")

	fmt.Println("Part One:", partOne(lines))
	fmt.Println("Part Two:", partTwo(lines))
}

func partOne(strs []string) int {
	return countStringsFunc(strs, partOnePredicate)
}

func partTwo(strs []string) int {
    return countStringsFunc(strs, partTwoPredicate)
}

func countStringsFunc(strs []string, predicate func(string) bool) int {
	count := 0
	for _, str := range strs {
		if predicate(str) {
			count += 1
		}
	}
	return count
}

func partOnePredicate(s string) bool {
	return hasThreeVowels(s) && hasDoubleLetter(s) && hasNoForbiddenStrings(s)
}

func partTwoPredicate(s string) bool {
    return hasRepeatingPair(s) && hasRpeatingWithOneBetween(s)
}

func hasThreeVowels(s string) bool {
	vowels := "aeiou"
	vowelCount := 0
	for _, letter := range s {
		for _, vowel := range vowels {
			if letter == vowel {
				vowelCount += 1
				break
			}
		}
		if vowelCount >= 3 {
			return true
		}
	}
	return false
}

func hasDoubleLetter(s string) bool {
	length := len(s)
	for i := 1; i < length; i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}

func hasNoForbiddenStrings(s string) bool {
	forbiddenStrings := [...]string{"ab", "cd", "pq", "xy"}
	for _, forbidden := range forbiddenStrings {
		if strings.Contains(s, forbidden) {
			return false
		}
	}
	return true
}

func hasRepeatingPair(s string) bool {
	length := len(s)
	for i := 1; i < length; i++ {
		pair := s[i-1 : i+1]
		count := strings.Count(s, pair)
		if count >= 2 {
			return true
		}
	}
	return false
}

func hasRpeatingWithOneBetween(s string) bool {
	length := len(s)
	for i := 2; i < length; i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
    return false
}
