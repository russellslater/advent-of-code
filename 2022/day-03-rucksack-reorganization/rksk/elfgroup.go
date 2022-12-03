package rksk

type ElfGroup []*Rucksack

func (g ElfGroup) CommonItem() Item {
	if len(g) <= 1 {
		return 0
	}

	// Unique items in first rucksack
	set := make(map[rune]bool)
	for _, item := range g[0].Items {
		set[item] = true
	}

	// Process remaining rucksacks
	for i := 1; i < len(g); i++ {
		// Unique items in current rucksack and then intersect
		s := make(map[rune]bool)
		for _, r := range g[i].Items {
			s[r] = true
		}
		set = intersection(set, s)
	}

	// Return first occurence of an item found in all rucksacks
	for r := range set {
		return Item(r)
	}

	return 0
}

func intersection(s1, s2 map[rune]bool) map[rune]bool {
	s := make(map[rune]bool)
	for r := range s1 {
		// If items exists second set, add it to the intersection
		if s2[r] {
			s[r] = true
		}
	}
	return s
}

func GroupElves(rucksacks []*Rucksack, size int) []ElfGroup {
	if size == 0 {
		return []ElfGroup{}
	}

	var g []ElfGroup
	for i := 0; i < len(rucksacks); i += size {
		end := i + size
		if end > len(rucksacks) {
			end = len(rucksacks)
		}
		g = append(g, rucksacks[i:end])
	}
	return g
}
