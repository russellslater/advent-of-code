package santastr

import "regexp"

type SantaStringV1 struct {
	str string
}

func NewSantaStringV1(str string) SantaString {
	return SantaStringV1{str: str}
}

func (s SantaStringV1) String() string {
	return s.str
}

func (s SantaStringV1) IsNice() bool {
	// Does not contain the strings ab, cd, pq, or xy
	var re = regexp.MustCompile(`ab|cd|pq|xy`)
	if re.MatchString(s.str) {
		return false
	}

	// Contains at least one letter that appears twice in a row
	isRepeatChar := false
	for i := 0; i < len(s.str)-1; i++ {
		if s.str[i] == s.str[i+1] {
			isRepeatChar = true
		}
	}

	if !isRepeatChar {
		return false
	}

	// Contains at least three vowels
	vowelCount := 0
	for _, char := range s.str {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowelCount++
		}
	}

	return vowelCount >= 3
}

func (s SantaStringV1) IsNaughty() bool {
	return !s.IsNice()
}
