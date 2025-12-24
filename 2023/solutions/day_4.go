package solutions

import (
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Winners map[int]bool
	Actual  map[int]bool
}

func parseNums(in string) map[int]bool {
	out := map[int]bool{}
	value := ""
	for _, c := range in {
		segment := string(c)

		if strings.Contains("0123456789", segment) {
			value += segment
			continue
		}

		if len(value) > 0 {
			num, _ := strconv.Atoi(value)
			out[num] = true
			value = ""
		}
	}

	if len(value) > 0 {
		num, _ := strconv.Atoi(value)
		out[num] = true
	}

	return out
}

func parseCard(in string) Card {
	split := strings.Split(in[9:], "|")

	return Card{
		Winners: parseNums(split[0]),
		Actual:  parseNums(split[1]),
	}
}

func D4P1(input []string) int {
	cards := []Card{}
	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	totalValue := 0
	for _, card := range cards {
		matches := []int{}
		for myNum, _ := range card.Actual {
			for winningNum, _ := range card.Winners {
				if myNum == winningNum {
					matches = append(matches, myNum)
				}
			}
		}

		if len(matches) > 0 {
			totalValue += int(math.Pow(2, float64(len(matches)-1)))
		}
	}

	return totalValue
}

func D4P2(input []string) int {
	cards := []Card{}
	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	counts := map[int]int{}
	for index, card := range cards {
		counts[index] += 1

		matches := 0
		for myNum, _ := range card.Actual {
			for winningNum, _ := range card.Winners {
				if myNum == winningNum {
					matches += 1
				}
			}
		}

		for i := index + 1; i <= index+matches; i++ {
			counts[i] += counts[index]
		}
	}

	total := 0
	for _, val := range counts {
		total += val
	}

	return total
}
