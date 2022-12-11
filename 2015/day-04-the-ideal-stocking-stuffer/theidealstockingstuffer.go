package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-04-the-ideal-stocking-stuffer/input.txt"
	inputLines := getTransformedInput(filename)

	num := lowestNumberForHashStart(inputLines, "00000")
	fmt.Printf("Part One Answer: %d\n", num)

	num = lowestNumberForHashStart(inputLines, "000000")
	fmt.Printf("Part Two Answer: %d\n", num)
}

func getTransformedInput(filename string) string {
	return util.LoadInput(filename)[0]
}

func lowestNumberForHashStart(key string, start string) int {
	num := 1
	for {
		hash := mD5HashHex(key + fmt.Sprint(num))
		if hash[:len(start)] == start {
			break
		}
		num++
	}
	return num
}

func mD5HashHex(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
