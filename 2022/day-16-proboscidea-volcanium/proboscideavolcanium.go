package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-16-proboscidea-volcanium/input.txt"
	valveMap := generateWeightedValveMap(getTransformedInput(filename))

	start := valveMap["AA"]

	maxMinutes := 30
	startMinute := 0

	paths := calcMaxReleasedPressure(start, startMinute, maxMinutes, Path{ReleasedPressure: 0, OpenValves: make([]string, 0)}, valveMap)
	fmt.Printf("Part One Answer: %d\n", paths[0].ReleasedPressure)

	startMinute = 4 // Time required to teach elephant

	// Simply calculate all possible paths and find the max combined pressure where there is no overlap in opened valves
	maxReleasedPressure := 0
	paths = calcMaxReleasedPressure(start, startMinute, maxMinutes, Path{ReleasedPressure: 0, OpenValves: make([]string, 0)}, valveMap)
	for _, humanPath := range paths {
		for _, elephantPath := range paths {
			combinedPressure := humanPath.ReleasedPressure + elephantPath.ReleasedPressure
			if combinedPressure > maxReleasedPressure && !humanPath.HasOpenValvesCollision(elephantPath) {
				maxReleasedPressure = combinedPressure
			}
		}
	}
	fmt.Printf("Part Two Answer: %d\n", maxReleasedPressure)
}

func getTransformedInput(filename string) map[string]Valve {
	valveMap := make(map[string]Valve)
	for _, line := range util.LoadInput(filename) {
		v := Valve{}
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &v.Name, &v.FlowRate)
		neighbours := make(map[string]int)
		for _, v := range strings.Split(strings.NewReplacer(" tunnels lead to valves ", "", " tunnel leads to valve ", "").Replace(strings.Split(line, ";")[1]), ", ") {
			neighbours[v] = 1
		}
		v.Neighbours = neighbours
		valveMap[v.Name] = v
	}
	return valveMap
}

type Valve struct {
	Name       string
	FlowRate   int
	Neighbours map[string]int // Cost in time to get to neighbour
}

type Path struct {
	OpenValves       []string
	ReleasedPressure int
}

func (p Path) HasOpenValvesCollision(other Path) bool {
	for _, v1 := range p.OpenValves {
		for _, v2 := range other.OpenValves {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func generateWeightedValveMap(valveMap map[string]Valve) map[string]Valve {
	// Find shortest path between all valves with non-zero flow rate
	weightedValveMap := make(map[string]Valve)
	for _, v := range valveMap {
		// Always start at AA, even if it has zero flow rate
		if v.Name != "AA" && v.FlowRate == 0 {
			continue
		}
		neighbours := make(map[string]int)
		for _, to := range valveMap {
			if valveMap[to.Name].FlowRate > 0 {
				neighbours[to.Name] = shortestDistance(v, to, valveMap) + 1
			}
		}
		weightedValveMap[v.Name] = Valve{Name: v.Name, FlowRate: v.FlowRate, Neighbours: neighbours}
	}
	return weightedValveMap
}

type shortestDistanceState struct {
	ValveName string
	Steps     int
}

func shortestDistance(start, end Valve, valveMap map[string]Valve) int {
	if start.Name == end.Name {
		return 0
	}

	visited := make(map[string]bool, 0) // Valve name as key
	var queue []shortestDistanceState
	queue = append(queue, shortestDistanceState{start.Name, 0})

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		// Finished yet?
		if currState.ValveName == end.Name {
			return currState.Steps
		}

		if visited[currState.ValveName] {
			continue
		}
		visited[currState.ValveName] = true

		for name, dist := range valveMap[currState.ValveName].Neighbours {
			if visited[name] {
				continue
			}
			queue = append(queue, shortestDistanceState{name, currState.Steps + dist})
		}
	}

	return -1
}

func calcMaxReleasedPressure(pos Valve, minute int, maxMinutes int, currentPath Path, valveMap map[string]Valve) []Path {
	paths := make([]Path, 0)

	// Where to explore next
	nextValveMap := make(map[string]Valve)
	for k, v := range valveMap {
		nextValveMap[k] = v
	}
	delete(nextValveMap, pos.Name)

	for _, valve := range nextValveMap {
		// Out of time!
		if minute+pos.Neighbours[valve.Name] >= maxMinutes {
			continue
		}

		openValves := make([]string, len(currentPath.OpenValves))
		copy(openValves, currentPath.OpenValves)

		path := Path{
			ReleasedPressure: currentPath.ReleasedPressure + valve.FlowRate*(maxMinutes-(minute+pos.Neighbours[valve.Name])),
			OpenValves:       append(openValves, valve.Name),
		}

		// More caves (with valves) to explore?
		if len(nextValveMap) > 1 {
			paths = append(paths, calcMaxReleasedPressure(valve, minute+pos.Neighbours[valve.Name], maxMinutes, path, nextValveMap)...)
		}

		paths = append(paths, path)
	}

	// Sort by released pressure; highest first
	sort.Slice(paths, func(i, j int) bool {
		return paths[i].ReleasedPressure > paths[j].ReleasedPressure
	})

	return paths
}
