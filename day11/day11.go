package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day11/input.txt")
	if err != nil {
		panic(err)
	}
	input := bytes.TrimSpace(file)

	partOneSolution := solveDay11(input)
	fmt.Printf("Part One: %s\n", partOneSolution)
	partTwoSolution := solveDay11(partOneSolution)
	fmt.Printf("Part Two: %s\n", partTwoSolution)
}

func solveDay11(password []byte) []byte {
	newPassword := make([]byte, len(password))
	copy(newPassword, password)

	incrementPassword(newPassword)
	for !checkAllConds(newPassword) {
		incrementPassword(newPassword)
	}

	return newPassword
}

const A_BYTE byte = byte('a')
const I_BYTE byte = byte('i')
const O_BYTE byte = byte('o')
const L_BYTE byte = byte('l')
const Z_BYTE byte = byte('z')

func incrementPassword(password []byte) {
	var carry byte = 1
	for i := len(password) - 1; i >= 0 && carry > 0; i -= 1 {
		oldChar := password[i]
		newChar := oldChar + carry
		carry = 0
		if newChar > Z_BYTE {
			newChar -= 26
			carry = 1
		}
		password[i] = newChar
	}
}

func checkAllConds(password []byte) bool {
	return checkCond1(password) && checkCond2(password) && checkCond3(password)
}

func checkCond1(password []byte) bool {
	for i := 2; i < len(password); i += 1 {
		if password[i-2] == password[i-1]-1 && password[i-1] == password[i]-1 {
			return true
		}
	}
	return false
}

func checkCond2(password []byte) bool {
	for _, c := range password {
		if c == I_BYTE || c == O_BYTE || c == L_BYTE {
			return false
		}
	}
	return true
}

func checkCond3(password []byte) bool {
	doubles := 0
	for i := 1; i < len(password); i += 1 {
		if password[i-1] == password[i] {
			doubles += 1
			i += 1
		}
	}
	return doubles >= 2
}
