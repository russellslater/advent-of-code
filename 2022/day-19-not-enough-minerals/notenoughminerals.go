package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-19-not-enough-minerals/input.txt"
	blueprints := getTransformedInput(filename)

	start := time.Now()

	totalQualityLevel := calcTotalQualityLevel(blueprints)
	fmt.Printf("Part One Answer: %v\n", totalQualityLevel)

	elapsed := time.Since(start)
	fmt.Printf("Execution Time: %s\n", elapsed) // ~1.7s

	start = time.Now()

	geodeProduct := calcGeodeProduct(blueprints, 3)
	fmt.Printf("Part Two Answer: %v\n", geodeProduct)

	elapsed = time.Since(start)
	fmt.Printf("Execution Time: %s\n", elapsed) // ~4.0s
}

type Blueprint struct {
	ID            int
	OreRobot      Cost
	ClayRobot     Cost
	ObsidianRobot Cost
	GeodeRobot    Cost
	MaxOre        int
}

type Cost struct {
	Ore      int
	Clay     int
	Obsidian int
}

type State struct {
	TimeRemaining int
	Ore           int
	Clay          int
	Obsidian      int
	Geode         int
	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int
}

func getTransformedInput(filename string) []Blueprint {
	blueprints := []Blueprint{}
	for _, line := range util.LoadInput(filename) {
		bp := strings.Split(line, ":")[0]
		costs := strings.Split(line, ":")[1]

		b := Blueprint{}
		fmt.Sscanf(bp, "Blueprint %d", &b.ID)

		b.OreRobot = Cost{}
		b.ClayRobot = Cost{}
		b.ObsidianRobot = Cost{}
		b.GeodeRobot = Cost{}
		fmt.Sscanf(costs, " Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &b.OreRobot.Ore, &b.ClayRobot.Ore, &b.ObsidianRobot.Ore, &b.ObsidianRobot.Clay, &b.GeodeRobot.Ore, &b.GeodeRobot.Obsidian)

		b.MaxOre = util.Max(util.Max(util.Max(b.OreRobot.Ore, b.ClayRobot.Ore), b.ObsidianRobot.Ore), b.GeodeRobot.Ore)

		blueprints = append(blueprints, b)
	}
	return blueprints
}

func calcTotalQualityLevel(blueprints []Blueprint) int {
	totalQualityLevel := 0
	for _, bp := range blueprints {
		best := searchForBestResult(bp, 24)
		totalQualityLevel += best * bp.ID
	}
	return totalQualityLevel
}

func calcGeodeProduct(blueprints []Blueprint, top int) int {
	product := 1
	for _, bp := range blueprints[:top] {
		best := searchForBestResult(bp, 32)
		product *= best
	}
	return product
}

func searchForBestResult(bp Blueprint, timeRemaining int) int {
	bestResult := 0
	var queue []State
	visited := make(map[State]bool, 0)

	start := State{
		TimeRemaining: timeRemaining,
		Ore:           0,
		Clay:          0,
		Obsidian:      0,
		Geode:         0,
		OreRobot:      1,
		ClayRobot:     0,
		ObsidianRobot: 0,
		GeodeRobot:    0,
	}

	queue = append(queue, start)

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		bestResult = util.Max(bestResult, currState.Geode)

		// Out of time!
		if currState.TimeRemaining == 0 {
			continue
		}

		// Throw away unnecessary materials
		// Minimum of (1) the current balance and (2) the maximum required to spend on the most expensive robot for the remaining time (factoring in collection)
		currState.Ore = util.Min(currState.Ore, (currState.TimeRemaining*bp.MaxOre)-(currState.OreRobot*(currState.TimeRemaining-1)))
		currState.Clay = util.Min(currState.Clay, (currState.TimeRemaining*bp.ObsidianRobot.Clay)-(currState.ClayRobot*(currState.TimeRemaining-1)))
		currState.Obsidian = util.Min(currState.Obsidian, (currState.TimeRemaining*bp.GeodeRobot.Obsidian)-(currState.ObsidianRobot*(currState.TimeRemaining-1)))

		// Already seen?
		if visited[currState] {
			continue
		}
		visited[currState] = true

		queue = append(queue, nextPossibleStates(currState, bp)...)
	}

	return bestResult
}

// 5 decisions that can be taken:
// 1) Wait and collect
// 2) Make ore robot
// 3) Make clay robot
// 4) Make obsidian robot
// 5) Make geode robot
func nextPossibleStates(state State, bp Blueprint) []State {
	decisions := []State{}

	// Make ore robot?
	// Have enough ore and not already producing max ore required?
	if state.Ore >= bp.OreRobot.Ore && state.OreRobot < bp.MaxOre {
		decisions = append(decisions, State{
			TimeRemaining: state.TimeRemaining - 1,
			Ore:           state.Ore - bp.OreRobot.Ore + state.OreRobot,
			Clay:          state.Clay + state.ClayRobot,
			Obsidian:      state.Obsidian + state.ObsidianRobot,
			Geode:         state.Geode + state.GeodeRobot,
			OreRobot:      state.OreRobot + 1,
			ClayRobot:     state.ClayRobot,
			ObsidianRobot: state.ObsidianRobot,
			GeodeRobot:    state.GeodeRobot,
		})
	}

	// Make clay robot?
	// Have enough ore and not already producing max clay required?
	if state.Ore >= bp.ClayRobot.Ore && state.ClayRobot < bp.ObsidianRobot.Clay {
		decisions = append(decisions, State{
			TimeRemaining: state.TimeRemaining - 1,
			Ore:           state.Ore - bp.ClayRobot.Ore + state.OreRobot,
			Clay:          state.Clay + state.ClayRobot,
			Obsidian:      state.Obsidian + state.ObsidianRobot,
			Geode:         state.Geode + state.GeodeRobot,
			OreRobot:      state.OreRobot,
			ClayRobot:     state.ClayRobot + 1,
			ObsidianRobot: state.ObsidianRobot,
			GeodeRobot:    state.GeodeRobot,
		})
	}

	// Make obsidian robot?
	// Have enough obsidian and not already producing max obsidian required?
	if state.Ore >= bp.ObsidianRobot.Ore && state.Clay >= bp.ObsidianRobot.Clay &&
		state.ObsidianRobot < bp.GeodeRobot.Obsidian {
		decisions = append(decisions, State{
			TimeRemaining: state.TimeRemaining - 1,
			Ore:           state.Ore - bp.ObsidianRobot.Ore + state.OreRobot,
			Clay:          state.Clay - bp.ObsidianRobot.Clay + state.ClayRobot,
			Obsidian:      state.Obsidian + state.ObsidianRobot,
			Geode:         state.Geode + state.GeodeRobot,
			OreRobot:      state.OreRobot,
			ClayRobot:     state.ClayRobot,
			ObsidianRobot: state.ObsidianRobot + 1,
			GeodeRobot:    state.GeodeRobot,
		})
	}

	// Make geode robot?
	if state.Ore >= bp.GeodeRobot.Ore && state.Obsidian >= bp.GeodeRobot.Obsidian {
		decisions = append(decisions, State{
			TimeRemaining: state.TimeRemaining - 1,
			Ore:           state.Ore - bp.GeodeRobot.Ore + state.OreRobot,
			Clay:          state.Clay + state.ClayRobot,
			Obsidian:      state.Obsidian - bp.GeodeRobot.Obsidian + state.ObsidianRobot,
			Geode:         state.Geode + state.GeodeRobot,
			OreRobot:      state.OreRobot,
			ClayRobot:     state.ClayRobot,
			ObsidianRobot: state.ObsidianRobot,
			GeodeRobot:    state.GeodeRobot + 1,
		})
	}

	// Wait and collect
	decisions = append(decisions, State{
		TimeRemaining: state.TimeRemaining - 1,
		Ore:           state.Ore + state.OreRobot,
		Clay:          state.Clay + state.ClayRobot,
		Obsidian:      state.Obsidian + state.ObsidianRobot,
		Geode:         state.Geode + state.GeodeRobot,
		OreRobot:      state.OreRobot,
		ClayRobot:     state.ClayRobot,
		ObsidianRobot: state.ObsidianRobot,
		GeodeRobot:    state.GeodeRobot,
	})

	return decisions
}
