package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

const (
	RootMonkey  = "root"
	HumanMonkey = "humn"
	Add         = "+"
	Subtract    = "-"
	Multiply    = "*"
	Divide      = "/"
)

func main() {
	filename := "./2022/day-21-monkey-math/input.txt"
	solver := getTransformedInput(filename)

	rootNum := solver.Solve(RootMonkey)
	fmt.Printf("Part One Answer: %v\n", rootNum)

	num := solver.SolveReversed(HumanMonkey)
	fmt.Printf("Part Two Answer: %v\n", num)
}

type MonkeyMathSolver struct {
	operations map[string]MonkeyOp
	numbers    map[string]int // Cache for stashing pre-calculated and calcuted numbers
}

func NewMonkeyMathSolver(operations map[string]MonkeyOp, numbers map[string]int) *MonkeyMathSolver {
	return &MonkeyMathSolver{operations: operations, numbers: numbers}
}

type MonkeyOp struct {
	LeftMonkey  string
	RightMonkey string
	Operator    string
}

func getTransformedInput(filename string) *MonkeyMathSolver {
	operations := map[string]MonkeyOp{}
	numbers := map[string]int{}

	for _, line := range util.LoadInput(filename) {
		parts := strings.Fields(line)
		monkey := strings.Replace(parts[0], ":", "", -1)
		if len(parts) == 4 {
			operations[monkey] = MonkeyOp{LeftMonkey: parts[1], RightMonkey: parts[3], Operator: parts[2]}
		} else {
			numbers[monkey] = util.MustAtoi(parts[1])
		}
	}
	return NewMonkeyMathSolver(operations, numbers)
}

func (m *MonkeyMathSolver) Solve(targetMonkey string) int {
	if num, ok := m.numbers[targetMonkey]; ok {
		return num
	}

	op := m.operations[targetMonkey]
	left := m.Solve(op.LeftMonkey)
	right := m.Solve(op.RightMonkey)

	switch op.Operator {
	case Add:
		return left + right
	case Subtract:
		return left - right
	case Multiply:
		return left * right
	case Divide:
		return left / right
	}

	return 0
}

func (m *MonkeyMathSolver) SolveReversed(targetMonkey string) int {
	// Find the monkey whose number is the result of the target monkey
	for monkey, op := range m.operations {
		// Is the target monkey the left or right operand?
		if op.LeftMonkey == targetMonkey {
			rightOperand := m.Solve(op.RightMonkey) // Right operand can be calculated simply
			if monkey == RootMonkey {
				return rightOperand // Left and right are equal!
			} else {
				result := m.SolveReversed(monkey) // Go all the way back up to the root monkey

				// Rearrange the equation to get the left operand
				switch op.Operator {
				case Add:
					return result - rightOperand // result = left + right -> left = result - right
				case Subtract:
					return result + rightOperand // result = left - right -> left = result + right
				case Multiply:
					return result / rightOperand // result = left * right -> left = result / right
				case Divide:
					return result * rightOperand // result = left / right -> left = result * right
				}
			}
		} else if op.RightMonkey == targetMonkey {
			leftOperand := m.Solve(op.LeftMonkey) // Left operand can be calculated simply
			if monkey == RootMonkey {
				return leftOperand // Left and right are equal!
			} else {
				result := m.SolveReversed(monkey) // Go all the way back up to the root monkey

				// Rearrange the equation to get the right operand
				switch op.Operator {
				case Add:
					return result - leftOperand // result = left + right -> right = result - left
				case Subtract:
					return leftOperand - result // result = left - right -> right = left - result
				case Multiply:
					return result / leftOperand // result = left * right -> right = result / left
				case Divide:
					return leftOperand / result // result = left / right -> right = left / result
				}
			}
		}
	}

	return 0
}
