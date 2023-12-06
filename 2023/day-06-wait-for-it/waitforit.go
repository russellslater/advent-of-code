package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

type RaceData struct {
	Times     []int
	Distances []int
}

func (r RaceData) LongerRaceTime() int {
	return util.MustAtoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r.Times)), ""), "[]"))
}

func (r RaceData) LongerRaceDistance() int {
	return util.MustAtoi(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(r.Distances)), ""), "[]"))
}

func main() {
	filename := "./2023/day-06-wait-for-it/input.txt"
	raceData := parseInput(filename)

	product := 1
	for i := 0; i < len(raceData.Times); i++ {
		wins := waysToWin(raceData.Times[i], raceData.Distances[i])
		product *= wins
	}

	fmt.Printf("Part One Answer: %d\n", product)

	waysToWinLongerRace := waysToWin(raceData.LongerRaceTime(), raceData.LongerRaceDistance())

	fmt.Printf("Part Two Answer: %d\n", waysToWinLongerRace)
}

func parseInput(filename string) RaceData {
	lines := util.LoadInput(filename)

	times := []int{}
	distances := []int{}

	for _, line := range lines {
		if strings.Contains(line, "Time:") {
			for _, t := range strings.Fields(strings.TrimPrefix(line, "Time:")) {
				times = append(times, util.MustAtoi(t))
			}
			fmt.Println(times, len(times))
		} else if strings.Contains(line, "Distance:") {
			for _, d := range strings.Fields(strings.TrimPrefix(line, "Distance:")) {
				distances = append(distances, util.MustAtoi(d))
			}
		}
	}

	return RaceData{Times: times, Distances: distances}
}

func waysToWin(time int, recordDistance int) int {
	wins := 0

	for i := 0; i <= time; i++ {
		// hold button for i seconds and then release
		speed := i
		timeRemaining := time - i
		distance := speed * timeRemaining

		if distance > recordDistance {
			wins += 1
		}
	}

	return wins
}
