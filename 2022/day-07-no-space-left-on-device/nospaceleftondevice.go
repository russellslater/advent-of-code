package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/russellslater/advent-of-code/2022/day-07-no-space-left-on-device/dir"
)

func main() {
	root := parseDirectoryStructure("./2022/day-07-no-space-left-on-device/input.txt")

	fmt.Printf("Part One Answer: %d\n", root.SumWithinSizeLimit(100_000))

	maxSpace := 70_000_000
	targetUnusedSpace := 30_000_000
	requiredSpace := root.TotalSize() - (maxSpace - targetUnusedSpace)

	fmt.Printf("Part Two Answer: %v\n", root.MinSizeOverTarget(requiredSpace))
}

func parseDirectoryStructure(filename string) *dir.Directory {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	root := &dir.Directory{
		Name:     "",
		Children: []*dir.Directory{},
		Size:     0,
	}
	currDir := root

	isCmdAvailable := false
	var tokens []string

	scanner := bufio.NewScanner(file)
	for isCmdAvailable || scanner.Scan() {
		if !isCmdAvailable {
			tokens = strings.Fields(scanner.Text())
		}
		isCmdAvailable = false

		if tokens[1] == "cd" {
			if tokens[2] == ".." {
				currDir = currDir.Parent
			} else if tokens[2] == "/" {
				currDir = root
			} else {
				dir := &dir.Directory{
					Name:     tokens[2],
					Parent:   currDir,
					Children: []*dir.Directory{},
					Size:     0,
				}
				currDir.Children = append(currDir.Children, dir)
				currDir = dir
			}
		} else if tokens[1] == "ls" {
			for scanner.Scan() {
				tokens = strings.Fields(scanner.Text())

				if tokens[0] == "$" {
					isCmdAvailable = true
					break
				}

				if tokens[0] != "dir" {
					num, _ := strconv.Atoi(tokens[0])
					currDir.Size += num
				}
			}
		}
	}

	return root
}
