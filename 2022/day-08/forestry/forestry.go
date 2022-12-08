package forestry

type Forest []string

func (f Forest) heightAt(x int, y int) int {
	return int(f[y][x] - '0')
}

func CountVisibleArbores(forest Forest) int {
	visible := make([][]bool, len(forest))
	for i := range visible {
		visible[i] = make([]bool, len(forest[0]))
	}

	for y := 0; y < len(forest); y++ {
		viewTrees(forest, visible, y, 0, 1, 0)                 // from left
		viewTrees(forest, visible, y, len(forest[0])-1, -1, 0) // from right
	}

	for x := 0; x < len(forest[0]); x++ {
		viewTrees(forest, visible, 0, x, 0, 1)              // from top
		viewTrees(forest, visible, len(forest)-1, x, 0, -1) // from bottom
	}

	count := 0
	for i := 0; i < len(visible); i++ {
		for j := 0; j < len(visible[i]); j++ {
			if visible[i][j] {
				count++
			}
		}
	}

	return count
}

func viewTrees(forest Forest, visible [][]bool, y int, x int, dx int, dy int) {
	maxHeight := -1
	for x >= 0 && x < len(forest[0]) && y >= 0 && y < len(forest) {
		if forest.heightAt(x, y) > maxHeight {
			if !visible[y][x] {
				visible[y][x] = true
			}
			maxHeight = forest.heightAt(x, y)
		}
		x += dx
		y += dy
	}
}

func MaxScenicScore(forest Forest) int {
	maxScore := 0
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[0]); x++ {
			up := calcScenicScoreInDirection(forest, x, y, 0, -1)
			down := calcScenicScoreInDirection(forest, x, y, 0, 1)
			left := calcScenicScoreInDirection(forest, x, y, -1, 0)
			right := calcScenicScoreInDirection(forest, x, y, 1, 0)

			score := up * down * left * right

			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func calcScenicScoreInDirection(forest Forest, x int, y int, dx int, dy int) int {
	score := 0
	height := forest.heightAt(x, y)

	x += dx
	y += dy

	for x >= 0 && x < len(forest[0]) && y >= 0 && y < len(forest) {
		score++
		if forest.heightAt(x, y) >= height {
			break
		}
		x += dx
		y += dy
	}

	return score
}
