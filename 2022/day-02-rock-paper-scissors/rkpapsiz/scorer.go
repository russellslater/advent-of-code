package rkpapsiz

const (
	R int = 1
	P int = 2
	S int = 3

	W int = 6
	D int = 3
	L int = 0
)

type StrategyScorer func(r *Round) int

func PlayInResponseScorer(r *Round) int {
	return map[string]int{
		"AX": D + R, // rock == rock
		"AY": W + P, // paper > rock
		"AZ": L + S, // scissors < rock
		"BX": L + R, // rock < paper
		"BY": D + P, // paper == paper
		"BZ": W + S, // scissors > paper
		"CX": W + R, // rock > scissors
		"CY": L + P, // paper < scissors
		"CZ": D + S, // scissors == scissors
	}[r.key()]
}

func EndRoundAsRequiredScorer(r *Round) int {
	return map[string]int{
		"AX": L + S, // L with scissor
		"AY": D + R, // D with rock
		"AZ": W + P, // W with paper
		"BX": L + R, // L with rock
		"BY": D + P, // D with paper
		"BZ": W + S, // W with scissors
		"CX": L + P, // L with paper
		"CY": D + S, // D with scissors
		"CZ": W + R, // W with rock
	}[r.key()]
}
