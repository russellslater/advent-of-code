package game

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID     int
	Rounds []Round
}

func (g Game) String() string {
	return fmt.Sprintf("Game %d: %+v [%d]", g.ID, g.Rounds, g.PowerMinCubeSets())
}

func (g Game) IsValid(blueCount int, redCount int, greenCount int) bool {
	for _, round := range g.Rounds {
		if round.BlueCount > blueCount {
			return false
		}
		if round.RedCount > redCount {
			return false
		}
		if round.GreenCount > greenCount {
			return false
		}
	}

	return true
}

func (r Game) PowerMinCubeSets() int {
	maxBlue := 0
	maxRed := 0
	maxGreen := 0

	for _, round := range r.Rounds {
		if round.BlueCount > maxBlue {
			maxBlue = round.BlueCount
		}
		if round.RedCount > maxRed {
			maxRed = round.RedCount
		}
		if round.GreenCount > maxGreen {
			maxGreen = round.GreenCount
		}
	}

	return maxBlue * maxRed * maxGreen
}

type Round struct {
	BlueCount  int
	RedCount   int
	GreenCount int
}

func NewGame(str string) *Game {
	game := Game{}
	game.ID, _ = parseGameID(str)

	str = str[len(fmt.Sprintf("Game %d: ", game.ID)):]

	game.Rounds = parseRounds(str)

	return &game
}

func parseGameID(s string) (int, error) {
	re := regexp.MustCompile(`Game (\d+):`)
	match := re.FindStringSubmatch(s)
	return strconv.Atoi(match[1])
}

func parseRounds(line string) []Round {
	rounds := []Round{}

	for _, str := range strings.Split(line, "; ") {
		round := Round{}

		for _, r := range strings.Split(str, ", ") {
			parts := strings.Split(r, " ")
			count, _ := strconv.Atoi(parts[0])

			switch parts[1] {
			case "red":
				round.RedCount = count
			case "blue":
				round.BlueCount = count
			case "green":
				round.GreenCount = count
			}
		}

		rounds = append(rounds, round)
	}

	return rounds
}
