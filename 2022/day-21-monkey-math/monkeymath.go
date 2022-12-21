package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-21-monkey-math/input.txt"
	operations, numbers := getTransformedInput(filename)

	rootNum := calcRootNumber(operations, numbers)
	fmt.Printf("Part One Answer: %v\n", rootNum)
}

const (
	Add      = "+"
	Subtract = "-"
	Multiply = "*"
	Divide   = "/"
)

type MonkeyOp struct {
	Monkey1  string
	Monkey2  string
	Operator string
}

func getTransformedInput(filename string) (map[string]MonkeyOp, map[string]int) {
	operations := map[string]MonkeyOp{}
	numbers := map[string]int{}

	for _, line := range util.LoadInput(filename) {
		parts := strings.Fields(line)
		monkey := strings.Replace(parts[0], ":", "", -1)
		if len(parts) == 4 {
			operations[monkey] = MonkeyOp{Monkey1: parts[1], Monkey2: parts[3], Operator: parts[2]}
		} else {
			numbers[monkey] = util.MustAtoi(parts[1])
		}
	}
	return operations, numbers
}

func calcRootNumber(operations map[string]MonkeyOp, numbers map[string]int) int {
	for {
		for monkey, op := range operations {
			if _, ok := numbers[monkey]; ok {
				continue
			}

			if _, ok := numbers[op.Monkey1]; !ok {
				continue
			}

			if _, ok := numbers[op.Monkey2]; !ok {
				continue
			}

			switch op.Operator {
			case Add:
				numbers[monkey] = numbers[op.Monkey1] + numbers[op.Monkey2]
			case Subtract:
				numbers[monkey] = numbers[op.Monkey1] - numbers[op.Monkey2]
			case Multiply:
				numbers[monkey] = numbers[op.Monkey1] * numbers[op.Monkey2]
			case Divide:
				numbers[monkey] = numbers[op.Monkey1] / numbers[op.Monkey2]
			}
		}

		if rootNum, ok := numbers["root"]; ok {
			return rootNum
		}
	}
}
