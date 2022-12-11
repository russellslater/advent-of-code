package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/russellslater/advent-of-code/2022/day-11-monkey-in-the-middle/cheeky"
)

func main() {
	filename := "./2022/day-11-monkey-in-the-middle/input.txt"
	troop := buildTroop(filename)

	manageWorry := func(value int) int { return value / 3 }
	level := calcMonkeyBusiness(troop, manageWorry, 20)
	fmt.Printf("Part One Answer: %v\n", level)

	troop = buildTroop(filename)
	lcm := troop.LowestCommonMultiple()
	manageWorry = func(value int) int { return value % lcm }
	level = calcMonkeyBusiness(troop, manageWorry, 10_000)
	fmt.Printf("Part Two Answer: %v\n", level)
}

func multiply(multiplier int) func(int) int {
	return func(num int) int {
		return num * multiplier
	}
}

func square() func(int) int {
	return func(num int) int {
		return num * num
	}
}

func add(addend int) func(int) int {
	return func(num int) int {
		return num + addend
	}
}

func buildTroop(filename string) cheeky.Troop {
	troop := cheeky.Troop{}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var m *cheeky.Monkey

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		switch {
		case strings.HasPrefix(line, "Monkey"):
			m = cheeky.NewMonkey()
			troop = append(troop, m)
		case strings.HasPrefix(line, "Starting items"):
			for _, item := range strings.Split(strings.Split(line, ": ")[1], ", ") {
				num, _ := strconv.Atoi(item)
				m.AddItem(num)
			}
		case strings.HasPrefix(line, "Operation"):
			operation := strings.Split(line, ": ")[1]
			if strings.HasPrefix(operation, "new = old * old") {
				m.Operation = square()
			} else if strings.HasPrefix(operation, "new = old * ") {
				multiplier, _ := strconv.Atoi(strings.Split(operation, "new = old * ")[1])
				m.Operation = multiply(multiplier)
			} else if strings.HasPrefix(operation, "new = old + ") {
				addend, _ := strconv.Atoi(strings.Split(operation, "new = old + ")[1])
				m.Operation = add(addend)
			}
		case strings.HasPrefix(line, "Test"):
			divisor, _ := strconv.Atoi(strings.Split(line, "by ")[1])
			m.Divisor = divisor
		case strings.HasPrefix(line, "If"):
			monkey, _ := strconv.Atoi(strings.Split(s.Text(), "monkey ")[1])
			m.AddReceiver(monkey)
		}
	}

	return troop
}

func calcMonkeyBusiness(troop []*cheeky.Monkey, manageWorry func(int) int, roundTotal int) int {
	highest := 0
	secondHighest := 0

	for round := 1; round <= roundTotal; round++ {
		for _, m := range troop {
			m.InspectItems(troop, manageWorry)

			if round < roundTotal {
				continue
			}

			if m.InspectionCount > highest {
				secondHighest = highest
				highest = m.InspectionCount
			} else if m.InspectionCount > secondHighest {
				secondHighest = m.InspectionCount
			}
		}
	}

	return highest * secondHighest
}
