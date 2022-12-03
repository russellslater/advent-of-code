package rksk

type Item rune

func (i Item) Score() int {
	if i >= 97 && i <= 122 {
		return int(i) - 96
	} else if i >= 65 && i <= 90 {
		return int(i) - 38
	}
	return 0
}
