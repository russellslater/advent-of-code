package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-10-cathode-ray-tube/device"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-10-cathode-ray-tube/input.txt"
	instructions := loadInstructions(filename)

	device := device.NewDevice()
	signalStrengthSum := device.Run(instructions)
	fmt.Printf("Part One Answer: %d\n", signalStrengthSum)

	fmt.Printf("Part Two Answer:\n")
	device.Display()
}

func loadInstructions(filename string) []*device.Instruction {
	instructions := []*device.Instruction{}
	for _, line := range util.LoadInput(filename) {
		ins := &device.Instruction{Cycles: 1}
		switch line {
		case "noop":
			// do nothing
		default:
			ins.Cycles = 2
			fmt.Sscanf(line, "addx %d", &ins.AddValue)
		}
		instructions = append(instructions, ins)
	}
	return instructions
}
