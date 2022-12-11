package main

import (
	"fmt"
)

func main() {
	filename := "./2022/day-11-monkey-in-the-middle/input.txt"
	troop := buildTroop(filename)

	manageWorry := func(value int) int { return value / 3 }
	level := calcMonkeyBusiness(troop, manageWorry, 20)
	fmt.Printf("Part One Answer: %v\n", level)

	troop = buildTroop(filename)
	manageWorry = func(value int) int { return value % troop.LowestCommonMultiple() }
	level = calcMonkeyBusiness(troop, manageWorry, 10_000)
	fmt.Printf("Part Two Answer: %v\n", level)
}

type Monkey struct {
	divisor         int
	operation       func(int) int
	items           []int
	monkeys         []*Monkey
	inspectionCount int
}

func NewMonkey(items []int, operation func(int) int, divisor int) *Monkey {
	return &Monkey{
		items:     items,
		operation: operation,
		divisor:   divisor,
		monkeys:   []*Monkey{},
	}
}

func (m *Monkey) AddMonkey(monkey ...*Monkey) {
	m.monkeys = append(m.monkeys, monkey...)
}

func (m *Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) InspectItems(worryLevel func(int) int) {
	for len(m.items) > 0 {
		m.SendItem(worryLevel)
	}
}

func (m *Monkey) SendItem(worryLevel func(int) int) {
	if len(m.items) == 0 {
		return
	}
	item := m.items[0]
	m.items = m.items[1:]

	value := m.operation(item)
	m.inspectionCount++

	value = worryLevel(value)

	if value%m.divisor == 0 {
		m.monkeys[0].AddItem(value)
	} else {
		m.monkeys[1].AddItem(value)
	}
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

type Troop []*Monkey

func (t Troop) LowestCommonMultiple() int {
	lcm := 1
	for _, monkey := range t {
		lcm *= monkey.divisor
	}
	return lcm
}

func buildTroop(filename string) Troop {
	troop := Troop{}

	troop = append(troop, NewMonkey([]int{57}, multiply(13), 11))
	troop = append(troop, NewMonkey([]int{58, 93, 88, 81, 72, 73, 65}, add(2), 7))
	troop = append(troop, NewMonkey([]int{65, 95}, add(6), 13))
	troop = append(troop, NewMonkey([]int{58, 80, 81, 83}, square(), 5))
	troop = append(troop, NewMonkey([]int{58, 89, 90, 96, 55}, add(3), 3))
	troop = append(troop, NewMonkey([]int{66, 73, 87, 58, 62, 67}, multiply(7), 17))
	troop = append(troop, NewMonkey([]int{85, 55, 89}, add(4), 2))
	troop = append(troop, NewMonkey([]int{73, 80, 54, 94, 90, 52, 69, 58}, add(7), 19))

	troop[0].AddMonkey(troop[3], troop[2])
	troop[1].AddMonkey(troop[6], troop[7])
	troop[2].AddMonkey(troop[3], troop[5])
	troop[3].AddMonkey(troop[4], troop[5])
	troop[4].AddMonkey(troop[1], troop[7])
	troop[5].AddMonkey(troop[4], troop[1])
	troop[6].AddMonkey(troop[2], troop[0])
	troop[7].AddMonkey(troop[6], troop[0])

	return troop
}

func calcMonkeyBusiness(troop []*Monkey, manageWorry func(int) int, roundTotal int) int {
	highest := 0
	secondHighest := 0

	for round := 1; round <= roundTotal; round++ {
		for _, m := range troop {
			m.InspectItems(manageWorry)

			if round < roundTotal {
				continue
			}

			if m.inspectionCount > highest {
				secondHighest = highest
				highest = m.inspectionCount
			} else if m.inspectionCount > secondHighest {
				secondHighest = m.inspectionCount
			}
		}
	}

	return highest * secondHighest
}
