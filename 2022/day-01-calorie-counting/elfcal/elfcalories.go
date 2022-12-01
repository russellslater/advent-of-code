package elfcal

import "sort"

type ElfCalories []int

func (e ElfCalories) TotalCalories() int {
	return calorieSum(e)
}

func calorieSum(calorieCounts []int) int {
	sum := 0
	for _, c := range calorieCounts {
		sum += c
	}
	return sum
}

func TopElfTotalCalories(calories []ElfCalories, top int) int {
	calorieCounts := make([]int, len(calories))
	for i, elf := range calories {
		calorieCounts[i] = elf.TotalCalories()
	}
	sort.Ints(calorieCounts)
	return calorieSum(calorieCounts[len(calorieCounts)-top:])
}
