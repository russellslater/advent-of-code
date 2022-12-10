package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Command int

const (
	Up      Command = 0
	Down    Command = 1
	Forward Command = 2
)

func DiveWithAim(commands []string) (horizontalPos int, depth int) {
	horizontalPos, depth = 0, 0
	aim := 0

	for i := 0; i < len(commands); i++ {
		cmdStr := commands[i]

		cmd, value := GetCommand(cmdStr)

		switch cmd {
		case Up:
			aim -= value
		case Down:
			aim += value
		case Forward:
			horizontalPos += value
			depth += (aim * value)
		}
	}

	return
}

func Dive(commands []string) (horizontalPos int, depth int) {
	horizontalPos, depth = 0, 0

	for i := 0; i < len(commands); i++ {
		cmdStr := commands[i]

		cmd, value := GetCommand(cmdStr)

		switch cmd {
		case Up:
			depth -= value
		case Down:
			depth += value
		case Forward:
			horizontalPos += value
		}
	}

	return
}

func GetCommand(cmdStr string) (cmd Command, value int) {
	r, _ := regexp.Compile(`(?P<cmd>[a-z]+)\s(?P<value>\d+)`)

	if r.MatchString(cmdStr) {
		match := r.FindStringSubmatch(cmdStr)

		switch match[1] {
		case "up":
			cmd = Up
		case "down":
			cmd = Down
		case "forward":
			cmd = Forward
		}

		value, _ = strconv.Atoi(match[2])
	}

	return cmd, value
}

func main() {
	inputLines := util.LoadInput("./2021/day-02-dive/input.txt")

	hPos, depth := Dive(inputLines)
	product := hPos * depth

	fmt.Printf("Dive - Horizontal Position: %d, Depth: %d, Product: %d\n", hPos, depth, product)

	hPos, depth = DiveWithAim(inputLines)
	product = hPos * depth

	fmt.Printf("Dive with Aim - Horizontal Position: %d, Depth: %d, Product: %d\n", hPos, depth, product)
}
