package aoc2023

import (
	"github.com/mariomac/gostream/item"
	"strconv"
	"strings"
)

type Game struct {
	ID        int
	MaxColors map[string]int
}

func toGame(line string) Game {
	parts := strings.Split(line, ":")

	fields := strings.Fields(parts[0])

	id, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}

	rounds := strings.Split(parts[1], ";")

	maxColors := make(map[string]int)

	for _, round := range rounds {
		pulls := strings.Split(round, ",")

		for _, pull := range pulls {
			fields := strings.Fields(pull)

			count, err := strconv.Atoi(fields[0])
			if err != nil {
				panic(err)
			}

			if count > maxColors[fields[1]] {
				maxColors[fields[1]] = count
			}
		}
	}

	return Game{
		ID:        id,
		MaxColors: maxColors,
	}
}

func isPossible(numRed, numGreen, numBlue int) func(Game) bool {
	return func(g Game) bool {
		return g.MaxColors["red"] <= numRed &&
			g.MaxColors["green"] <= numGreen &&
			g.MaxColors["blue"] <= numBlue
	}
}

func (g Game) ToID() int {
	return g.ID
}

func SumGames(games []string) int {
	sliceGames := Slice[string](games)

	possibleGames := Map(sliceGames, toGame).
		Filter(isPossible(12, 13, 14))

	sum, ok := Map(possibleGames, Game.ToID).
		Reduce(item.Add[int])
	if !ok {
		panic("no games")
	}

	return sum
}

func toPower(g Game) int {
	return g.MaxColors["red"] * g.MaxColors["green"] * g.MaxColors["blue"]
}

func SumPowerGames(games []string) int {
	stringGames := Slice[string](games)

	sliceGames := Map(stringGames, toGame)

	sum, ok := Map(sliceGames, toPower).
		Reduce(item.Add[int])
	if !ok {
		panic("no games")
	}

	return sum
}
