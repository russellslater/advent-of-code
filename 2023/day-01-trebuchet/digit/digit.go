package digit

var SpelledDigits = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

type Digit struct {
	Idx   int
	Value int
}

var EmptyDigit = Digit{-1, -1}

func (d Digit) IsEmpty() bool {
	return d.Idx == -1
}

func (d Digit) HasLargerIndex(x Digit) bool {
	return d.Idx > x.Idx
}
