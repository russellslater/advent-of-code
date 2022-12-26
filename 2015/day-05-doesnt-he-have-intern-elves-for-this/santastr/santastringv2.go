package santastr

type SantaStringV2 struct {
	str string
}

func NewSantaStringV2(str string) SantaString {
	return SantaStringV2{str: str}
}

func (s SantaStringV2) String() string {
	return s.str
}

func (s SantaStringV2) IsNice() bool {
	// Contains a pair of any two letters that appears at least twice in the string
	repeats := map[string][]int{}
	for i := 0; i < len(s.str)-1; i++ {
		key := string(s.str[i]) + string(s.str[i+1])
		repeats[key] = append(repeats[key], i)
	}

	isRepeatPair := false
	for _, indices := range repeats {
		if len(indices) >= 2 {
			if indices[len(indices)-1]-indices[0] > 1 {
				isRepeatPair = true
			}
		}
	}

	if !isRepeatPair {
		return false
	}

	// Contains at least one letter which repeats with exactly one letter between them
	isRepeatSandwich := false
	for i := 0; i < len(s.str)-2; i++ {
		if s.str[i] == s.str[i+2] {
			isRepeatSandwich = true
		}
	}

	return isRepeatSandwich
}

func (s SantaStringV2) IsNaughty() bool {
	return !s.IsNice()
}
