package main

import (
	"log"
	"os"
	"strings"
)

const (
	Day       = "10"
	InputFile = "./2021/_" + Day + "/input.txt"
)

func isOpen(r rune) bool {
	return r == '[' || r == '{' || r == '(' || r == '<'
}

func asPoints(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}

func matchesOpen(open, close rune) bool {
	return open == '[' && close == ']' || open == '{' && close == '}' || open == '(' && close == ')' || open == '<' && close == '>'
}

func p1(inputs []string) (int, error) {
	score := 0
	for _, line := range inputs {
		tmp := line
		for strings.Contains(tmp, "[]") || strings.Contains(tmp, "()") || strings.Contains(tmp, "{}") || strings.Contains(tmp, "<>") {
			tmp = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tmp, "[]", ""), "()", ""), "{}", ""), "<>", "")
		}

		for _, r := range tmp {
			if !isOpen(r) {
				score += asPoints(r)
				break
			}
		}
	}

	return score, nil
}

func p2(inputs []string) (int, error) {
	var scores []int
	for _, line := range inputs {
		tmp := line
		for strings.Contains(tmp, "[]") || strings.Contains(tmp, "()") || strings.Contains(tmp, "{}") || strings.Contains(tmp, "<>") {
			tmp = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(tmp, "[]", ""), "()", ""), "{}", ""), "<>", "")
		}

		points := 0
		for i := len(tmp) - 1; i >= 0; i-- {
			r := rune(tmp[i])
			if !isOpen(r) {
				points = -1
				break
			}
			points = points*5 + strings.IndexRune(" ([{<", r)
		}

		if points != -1 {
			index := -1
			for i, score := range scores {
				if points > score {
					index = i
					break
				}
			}

			if index != -1 {
				scores = append(scores[:index], append([]int{points}, scores[index:]...)...)
			} else {
				scores = append(scores, points)
			}

		}
	}

	return scores[int(len(scores)/2)], nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day %v, part 1", err, Day)
	}
	log.Printf("Day %v, Part 1: %v", Day, ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day %v, part 2", err, Day)
	}

	log.Printf("Day %v, Part 2: %v", Day, ans2)
}
