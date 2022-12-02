package rkpapsiz

type Round struct {
	Opponent rune
	You      rune
}

func (r *Round) key() string {
	return string([]rune{r.Opponent, r.You})
}

func (r *Round) Score(s StrategyScorer) int {
	return s(r)
}
