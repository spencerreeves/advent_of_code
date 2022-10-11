package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile = "./2021/_6/input.txt"

func p1(inputs []string) (int, error) {
	inputs = strings.Split(inputs[0], ",")
	fishes := make([]*int, len(inputs))
	for i, v := range inputs {
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.New("invalid fish number")
		}

		fishes[i] = &num
	}

	var newFishes []*int
	for i := 0; i < 80; i++ {
		for _, fish := range fishes {
			if *fish == 0 {
				f := 8
				*fish = 6
				newFishes = append(newFishes, &f)
			} else {
				*fish -= 1
			}
		}

		fishes = append(fishes, newFishes...)
		newFishes = []*int{}
	}

	return len(fishes), nil
}

func p2(inputs []string) (int, error) {
	inputs = strings.Split(inputs[0], ",")
	fishes := make([]int, 10)
	for _, v := range inputs {
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.New("invalid fish number")
		}

		fishes[num] += 1
	}

	for i := 0; i < 256; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				fishes[9] = fishes[0]
				fishes[7] += fishes[0]
			} else {
				fishes[j-1] = fishes[j]
			}
		}
		fishes[9] = 0
	}

	total := 0
	for _, v := range fishes {
		total += v
	}

	return total, nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day 6, part 1", err)
	}
	log.Printf("Day 6, Part 1: %v", ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day 6, part 2", err)
	}

	log.Printf("Day 6, Part 2: %v", ans2)
}
