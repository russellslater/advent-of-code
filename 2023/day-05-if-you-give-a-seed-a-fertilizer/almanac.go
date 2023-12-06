package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

type SeedRange struct {
	SrcStart, Length int64
}

type MappingRange struct {
	DestStart, SrcStart, Length int64
}

var steps = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func main() {
	filename := "./2023/day-05-if-you-give-a-seed-a-fertilizer/input.txt"
	lines := util.LoadInput(filename)

	var seeds []int64           // for Part One
	seedRanges := []SeedRange{} // for Part Two
	mappings := make(map[string][]MappingRange)
	var currentMap string

	for _, line := range lines {
		if strings.Contains(line, "map:") {
			currentMap = strings.TrimSpace(strings.TrimSuffix(strings.Split(line, ":")[0], " map"))
		} else if strings.HasPrefix(line, "seeds:") {
			seedParts := strings.Fields(strings.TrimSpace(strings.TrimPrefix(line, "seeds:")))
			// parse for Part One
			for _, s := range seedParts {
				seed, _ := strconv.ParseInt(s, 10, 64)
				seeds = append(seeds, seed)
			}

			// parse for Part Two
			for i := 0; i < len(seedParts); i += 2 {
				start, _ := strconv.ParseInt(seedParts[i], 10, 64)
				length, _ := strconv.ParseInt(seedParts[i+1], 10, 64)
				seedRanges = append(seedRanges, SeedRange{SrcStart: start, Length: length})
			}

		} else if line != "" {
			parts := strings.Fields(line)
			destStart, _ := strconv.ParseInt(parts[0], 10, 64)
			srcStart, _ := strconv.ParseInt(parts[1], 10, 64)
			length, _ := strconv.ParseInt(parts[2], 10, 64)
			mappings[currentMap] = append(mappings[currentMap], MappingRange{DestStart: destStart, SrcStart: srcStart, Length: length})
		}
	}

	lowestLocation := findLowestLocationNumber(seeds, mappings)
	fmt.Printf("Part One Answer: %d\n", lowestLocation)

	lowestLocation = findLowestLocationNumberForSeedRanges(seedRanges, mappings)
	fmt.Printf("Part Two Answer: %d\n", lowestLocation)
}

func findLowestLocationNumber(seeds []int64, mappings map[string][]MappingRange) int64 {
	lowestLocation := int64(-1)
	for _, seed := range seeds {
		location := calcLocationForSeed(seed, mappings)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func findLowestLocationNumberForSeedRanges(seedRanges []SeedRange, mappings map[string][]MappingRange) int64 {
	numCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numCores)

	results := make(chan int64, numCores)

	// define routine for processing a range of seeds
	processRanges := func(ranges []SeedRange) {
		var localLowest int64 = -1
		for _, seedRange := range ranges {
			for i := int64(0); i < seedRange.Length; i++ {
				seed := seedRange.SrcStart + i
				location := calcLocationForSeed(seed, mappings)
				if localLowest == -1 || location < localLowest {
					localLowest = location
				}
			}
		}
		results <- localLowest
	}

	// divide up the work by core
	for i := 0; i < numCores; i++ {
		start := i * len(seedRanges) / numCores
		end := (i + 1) * len(seedRanges) / numCores
		if start < len(seedRanges) {
			go processRanges(seedRanges[start:end])
		} else {
			results <- -1 // no work to do
		}
	}

	// gather the results
	globalLowest := int64(math.MaxInt64)
	for i := 0; i < numCores; i++ {
		localLowest := <-results
		if localLowest != -1 && localLowest < globalLowest {
			globalLowest = localLowest
		}
	}

	if globalLowest == math.MaxInt64 {
		return -1 // valid location not found
	}

	return globalLowest
}

func calcLocationForSeed(num int64, mappings map[string][]MappingRange) int64 {
	for _, step := range steps {
		num = convert(num, mappings[step])
	}
	return num
}

func convert(number int64, ranges []MappingRange) int64 {
	for _, r := range ranges {
		if number >= r.SrcStart && number < r.SrcStart+r.Length {
			return r.DestStart + (number - r.SrcStart)
		}
	}
	return number // not found
}
