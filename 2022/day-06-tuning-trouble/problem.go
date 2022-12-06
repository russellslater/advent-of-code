package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-06-tuning-trouble/device"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-06-tuning-trouble/input.txt"
	sig := getTransformedInput(filename)
	fmt.Printf("Part One Answer: %v\n", device.Device{DetectionLimit: 4}.Detect(sig))
	fmt.Printf("Part Two Answer: %v\n", device.Device{DetectionLimit: 14}.Detect(sig))
}

func getTransformedInput(filename string) string {
	return util.LoadInput(filename)[0]
}
