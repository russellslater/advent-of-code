package asspair

type Range struct {
	Start int
	End   int
}

type AssignmentPair struct {
	One Range
	Two Range
}

func (a *AssignmentPair) IsOverlap() bool {
	return rangeOverlaps(a.One, a.Two) || rangeOverlaps(a.Two, a.One)
}

func (a *AssignmentPair) IsContained() bool {
	return rangeContains(a.One, a.Two) || rangeContains(a.Two, a.One)
}

func rangeContains(target Range, other Range) bool {
	return target.Start <= other.Start && target.End >= other.End
}

func rangeOverlaps(target Range, other Range) bool {
	return other.Start >= target.Start && other.Start <= target.End
}
