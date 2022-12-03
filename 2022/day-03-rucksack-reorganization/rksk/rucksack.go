package rksk

type Rucksack struct {
	Items string
}

func (r *Rucksack) ReoccuringItem() Item {
	mid := len(r.Items) / 2
	firstHalf := r.Items[:mid]
	secondHalf := r.Items[mid:]

	itemsFound := map[rune]bool{}

	for _, item := range firstHalf {
		itemsFound[item] = true
	}

	for _, item := range secondHalf {
		if _, ok := itemsFound[item]; ok {
			return Item(item)
		}
	}

	return 0
}
