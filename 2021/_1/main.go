package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile = "./2021/_1/input.txt"

func p1(inputs []string) (int, error) {
	var prev *int
	var inc int
	for _, v := range inputs {
		val, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.Wrap(err, "invalid input")
		}

		if prev != nil && val > *prev {
			inc++
		}

		prev = &val
	}

	return inc, nil
}

func p2(inputs []string) (int, error) {
	var i1, i2, i3 int
	var err error
	if i1, err = strconv.Atoi(inputs[0]); err != nil {
		return 0, errors.Wrap(err, "invalid input i1")
	}
	if i2, err = strconv.Atoi(inputs[1]); err != nil {
		return 0, errors.Wrap(err, "invalid input i2")
	}
	if i3, err = strconv.Atoi(inputs[2]); err != nil {
		return 0, errors.Wrap(err, "invalid input i3")
	}

	var val, inc int
	for _, v := range inputs[3:] {
		val, err = strconv.Atoi(v)
		if err != nil {
			return 0, errors.Wrap(err, "invalid input")
		}

		if i1+i2+i3 < i2+i3+val {
			inc++
		}

		i1 = i2
		i2 = i3
		i3 = val
	}

	return inc, nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day 1, part 1", err)
	}
	log.Printf("Day 1, Part 1: %v", ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day 1, part 2", err)
	}

	log.Printf("Day 1, Part 2: %v", ans2)
}
