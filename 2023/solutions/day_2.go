package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	ID    int
	Plays []Play
}

type Play struct {
	Blue  int
	Green int
	Red   int
}

func parseInput(in string) Game {
	id, err := strconv.Atoi(in[5:strings.Index(in, ":")])
	if err != nil {
		return Game{}
	}

	var plays []Play
	for _, hand := range strings.Split(in[strings.Index(in, ":")+1:], ";") {
		var play Play

		for _, segment := range strings.Split(hand, ",") {
			// Skip first character since it will always be a space
			split := strings.Split(segment[1:], " ")

			count, err := strconv.Atoi(split[0])
			if err != nil {
				panic(fmt.Errorf("count is not an int[%v]: %w", split[0], err))
			}

			// The Third element is the qualifier
			switch split[1] {
			case "blue":
				play.Blue = count
			case "green":
				play.Green = count
			case "red":
				play.Red = count
			}
		}

		plays = append(plays, play)
	}

	return Game{
		ID:    id,
		Plays: plays,
	}
}

func D2P1(input []string) int {
	var games []Game
	for _, in := range input {
		games = append(games, parseInput(in))
	}

	possibleIdSums := 0
	for _, game := range games {
		possible := true

		for _, play := range game.Plays {
			if play.Blue > 14 || play.Green > 13 || play.Red > 12 {
				possible = false
				break
			}
		}

		if possible {
			possibleIdSums += game.ID
		}
	}

	return possibleIdSums
}

func D2P2(input []string) int {
	var games []Game
	for _, in := range input {
		games = append(games, parseInput(in))
	}

	var powerSums int
	for _, game := range games {
		var b, g, r int
		for _, play := range game.Plays {
			if b <= play.Blue {
				b = play.Blue
			}

			if g <= play.Green {
				g = play.Green
			}

			if r <= play.Red {
				r = play.Red
			}
		}

		powerSums += b * g * r
	}

	return powerSums
}
