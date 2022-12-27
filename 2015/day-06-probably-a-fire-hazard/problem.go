package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/twod"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-06-probably-a-fire-hazard/input.txt"
	instructions := getTransformedInput(filename)

	count := countLitLights(instructions)
	fmt.Printf("Part One Answer: %d\n", count)

	count = calcTotalBrightness(instructions)
	fmt.Printf("Part One Answer: %d\n", count)
}

func getTransformedInput(filename string) []Instruction {
	instructions := []Instruction{}
	for _, line := range util.LoadInput(filename) {
		ins := Instruction{}

		switch {
		case strings.HasPrefix(line, "turn on"):
			ins.Action = On
			line = strings.TrimPrefix(line, "turn on ")
		case strings.HasPrefix(line, "turn off"):
			ins.Action = Off
			line = strings.TrimPrefix(line, "turn off ")
		case strings.HasPrefix(line, "toggle"):
			ins.Action = Toggle
			line = strings.TrimPrefix(line, "toggle ")
		}

		positions := strings.Split(line, " through ")
		ins.TopLeft = twod.ParsePosition(positions[0])
		ins.BottomRight = twod.ParsePosition(positions[1])

		instructions = append(instructions, ins)
	}
	return instructions
}

type Instruction struct {
	Action      Action
	TopLeft     twod.Position
	BottomRight twod.Position
}

type Action int

const (
	On Action = iota
	Off
	Toggle
)

func countLitLights(instructions []Instruction) int {
	maxX, maxY := 999, 999

	lights := make([][]bool, maxX+1)
	for i := range lights {
		lights[i] = make([]bool, maxY+1)
	}

	count := 0
	for _, ins := range instructions {
		for x := ins.TopLeft.X; x <= ins.BottomRight.X; x++ {
			for y := ins.TopLeft.Y; y <= ins.BottomRight.Y; y++ {
				prevState := lights[x][y]
				switch ins.Action {
				case On:
					lights[x][y] = true
				case Off:
					lights[x][y] = false
				case Toggle:
					lights[x][y] = !lights[x][y]
				}
				if prevState != lights[x][y] {
					if lights[x][y] {
						count++
					} else {
						count--
					}
				}
			}
		}
	}

	return count
}

func calcTotalBrightness(instructions []Instruction) int {
	maxX, maxY := 999, 999

	lights := make([][]int, maxX+1)
	for i := range lights {
		lights[i] = make([]int, maxY+1)
	}

	brightness := 0

	for _, ins := range instructions {
		for x := ins.TopLeft.X; x <= ins.BottomRight.X; x++ {
			for y := ins.TopLeft.Y; y <= ins.BottomRight.Y; y++ {
				prevState := lights[x][y]
				switch ins.Action {
				case On:
					lights[x][y] += 1
				case Off:
					if lights[x][y] > 0 {
						lights[x][y] -= 1
					}
				case Toggle:
					lights[x][y] += 2
				}
				brightness += (prevState - lights[x][y]) * -1
			}
		}
	}

	return brightness
}
