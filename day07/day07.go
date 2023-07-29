package main

import (
	"fmt"
    "os"
	"strconv"
	"strings"
)

func main() {
    f, err := os.ReadFile("day07/input.txt")
    if err != nil {
        panic(err)
    }
    input := strings.TrimSpace(string(f))
    lines := strings.Split(input, "\n")
    wires := make(map[string]wire)
    for _, line := range lines {
        name, w := parseLine(line)
        wires[name] = w
    }
    fmt.Println("Part One:", partOne(wires))
    fmt.Println("Part Two:", partTwo(wires))
}

func partOne(wires map[string]wire) uint16 {
    cache := make(map[string]uint16)
    a := calcWire("a", wires, cache)
    return a
}

func partTwo(wires map[string]wire) uint16 {
    cache := make(map[string]uint16)
    a := calcWire("a", wires, cache)
    cache = make(map[string]uint16)
    cache["b"] = a
    a = calcWire("a", wires, cache)
    return a
}

type operation int
type operandType int

const (
	WIRE   operation = 0
	NOT              = 1
	AND              = 2
	OR               = 3
	LSHIFT           = 4
	RSHIFT           = 5
)

const (
	VALUEOP operandType = 0
	WIREOP              = 1
)

type operand struct {
	opType    operandType
	intValue  uint16
	wireValue string
}

type wire struct {
	op   operation
	l, r operand
}

func calcWire(name string, wires map[string]wire, cache map[string]uint16) uint16 {
	cachedValue, exists := cache[name]
	if exists {
		return cachedValue
	}

	w := wires[name]
	var value uint16
	if w.op == WIRE {
		value = calcOperand(w.l, wires, cache)
	} else if w.op == NOT {
		value = calcOperand(w.l, wires, cache)
		value = ^value
	} else {
        l := calcOperand(w.l, wires, cache)
        r := calcOperand(w.r, wires, cache)
		switch w.op {
		case AND:
            value = l & r
        case OR:
            value = l | r
        case LSHIFT:
            value = l << r
        case RSHIFT:
            value = l >> r
		}
	}

	cache[name] = value
	return value
}

func calcOperand(op operand, wires map[string]wire, cache map[string]uint16) uint16 {
	if op.opType == VALUEOP {
		return op.intValue
	}
	cachedValue, exists := cache[op.wireValue]
	if exists {
		return cachedValue
	}
	value := calcWire(op.wireValue, wires, cache)
	cache[op.wireValue] = value
	return value
}

func parseLine(line string) (string, wire) {
	fields := strings.Fields(line)
	tokenCount := len(fields)
	wireName := fields[tokenCount-1]
	var op operation
	var l, r operand
	if fields[0] == "NOT" {
		op = NOT
		l = parseOperand(fields[1])
	} else if tokenCount == 5 {
		l = parseOperand(fields[0])
		r = parseOperand(fields[2])
		switch fields[1] {
		case "AND":
			op = AND
		case "OR":
			op = OR
		case "LSHIFT":
			op = LSHIFT
		case "RSHIFT":
			op = RSHIFT
		}
	} else {
		op = WIRE
		l = parseOperand(fields[0])
	}
	return wireName, wire{op, l, r}
}

func parseOperand(token string) operand {
	value, err := strconv.Atoi(token)
	if err != nil {
		return operand{opType: WIREOP, wireValue: token}
	} else {
		return operand{opType: VALUEOP, intValue: uint16(value)}
	}
}
