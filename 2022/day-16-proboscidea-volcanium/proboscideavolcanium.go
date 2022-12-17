package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-16-proboscidea-volcanium/input.txt"
	valves := getTransformedInput(filename)

	output := solveProblem(valves)
	fmt.Printf("Solution output: %v\n", output)
}

type Valve struct {
	Name            string
	FlowRate        int
	ConnectedValves []*Valve
}

func (v *Valve) String() string {
	cvs := []string{}
	for _, cv := range v.ConnectedValves {
		cvs = append(cvs, cv.Name)
	}
	return fmt.Sprintf("Valve %s | Flow Rate: %-2d | Connected Valves: %s", v.Name, v.FlowRate, strings.Join(cvs, ", "))
}

func getTransformedInput(filename string) []*Valve {
	valves := []*Valve{}
	connectedValves := map[*Valve][]string{}

	for _, line := range util.LoadInput(filename) {
		v := &Valve{}
		v.ConnectedValves = []*Valve{}
		fmt.Sscanf(line, "Valve %s has flow rate=%d;", &v.Name, &v.FlowRate)

		tunnels := strings.Split(strings.NewReplacer(" tunnels lead to valves ", "", " tunnel leads to valve ", "").Replace(strings.Split(line, ";")[1]), ", ")
		connectedValves[v] = tunnels

		valves = append(valves, v)
	}

	// TODO: Optimize
	for valve, tunnels := range connectedValves {
		for _, tunnel := range tunnels {
			for _, v := range valves {
				if v.Name == tunnel {
					valve.ConnectedValves = append(valve.ConnectedValves, v)
					break
				}
			}
		}
	}

	return valves
}

func solveProblem(valves []*Valve) int {
	for _, v := range valves {
		fmt.Println(v)
	}

	return 0
}
