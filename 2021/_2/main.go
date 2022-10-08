package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile = "./2021/_2/input.txt"

func p1(inputs []string) (int, error) {
	var hor, dep int
	for _, v := range inputs {
		var token, rawAmt string
		if parsed := strings.Split(v, " "); len(parsed) != 2 {
			return 0, errors.New("invalid input")
		} else {
			token = parsed[0]
			rawAmt = parsed[1]
		}

		amt, err := strconv.Atoi(rawAmt)
		if err != nil {
			return 0, errors.Wrap(err, "amt conversion")
		}

		switch token {
		case "forward":
			hor += amt
		case "down":
			dep += amt
		case "up":
			dep -= amt
		default:
			return 0, errors.New("invalid action")
		}
	}

	return hor * dep, nil
}

func p2(inputs []string) (int, error) {
	var hor, dep, aim int
	for _, v := range inputs {
		var token, rawAmt string
		if parsed := strings.Split(v, " "); len(parsed) != 2 {
			return 0, errors.New("invalid input")
		} else {
			token = parsed[0]
			rawAmt = parsed[1]
		}

		amt, err := strconv.Atoi(rawAmt)
		if err != nil {
			return 0, errors.Wrap(err, "amt conversion")
		}

		switch token {
		case "forward":
			hor += amt
			dep += aim * amt
		case "down":
			aim += amt
		case "up":
			aim -= amt
		default:
			return 0, errors.New("invalid action")
		}
	}

	return hor * dep, nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day 2, part 1", err)
	}
	log.Printf("Day 2, Part 1: %v", ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day 2, part 2", err)
	}

	log.Printf("Day 2, Part 2: %v", ans2)
}
