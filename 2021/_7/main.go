package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Day       = "7"
	InputFile = "./2021/_" + Day + "/input.txt"
)

func p1(inputs []string) (int, error) {
	inputs = strings.Split(inputs[0], ",")
	var crabs []int
	for _, v := range inputs {
		pos, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.Wrap(err, "invalid crab position")
		}

		if len(crabs) <= pos {
			crabs = append(crabs, make([]int, pos-len(crabs)+1)...)
		}

		crabs[pos]++
	}

	moveRight, totalCrabs := make([]int, len(crabs)), crabs[0]
	for i := 1; i < len(crabs); i++ {
		moveRight[i] = totalCrabs + moveRight[i-1]
		totalCrabs = totalCrabs + crabs[i]
	}

	moveLeft, totalCrabs := make([]int, len(crabs)), crabs[len(crabs)-1]
	for i := len(crabs) - 2; i >= 0; i-- {
		moveLeft[i] = totalCrabs + moveLeft[i+1]
		totalCrabs = totalCrabs + crabs[i]
	}

	effort, minEffort := make([]int, len(crabs)), moveRight[0]+moveLeft[0]
	for i, _ := range crabs {
		effort[i] = moveRight[i] + moveLeft[i]
		if effort[i] < minEffort {
			minEffort = effort[i]
		}
	}

	return minEffort, nil
}

func p2(inputs []string) (int, error) {
	inputs = strings.Split(inputs[0], ",")
	var crabs []int
	for _, v := range inputs {
		pos, err := strconv.Atoi(v)
		if err != nil {
			return 0, errors.Wrap(err, "invalid crab position")
		}

		if len(crabs) <= pos {
			crabs = append(crabs, make([]int, pos-len(crabs)+1)...)
		}

		crabs[pos]++
	}

	moveRight, weight, totalCrabs := make([]int, len(crabs)), crabs[0], crabs[0]
	for i := 1; i < len(crabs); i++ {
		moveRight[i] = weight + moveRight[i-1]
		totalCrabs = totalCrabs + crabs[i]
		weight = weight + totalCrabs
	}

	moveLeft, weight, totalCrabs := make([]int, len(crabs)), crabs[len(crabs)-1], crabs[len(crabs)-1]
	for i := len(crabs) - 2; i >= 0; i-- {
		moveLeft[i] = weight + moveLeft[i+1]
		totalCrabs = totalCrabs + crabs[i]
		weight = weight + totalCrabs
	}

	effort, minEffort := make([]int, len(crabs)), moveRight[0]+moveLeft[0]
	for i, _ := range crabs {
		effort[i] = moveRight[i] + moveLeft[i]
		if effort[i] < minEffort {
			minEffort = effort[i]
		}
	}

	return minEffort, nil
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
