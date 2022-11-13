package util

import (
	"bufio"
	"log"
	"os"
)

func LoadInput(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer f.Close()

	input := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
