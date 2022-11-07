package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

const (
	Day       = "12"
	InputFile = "./2021/_" + Day + "/input.txt"
)

type Cave struct {
	ID          string
	Connections []*Cave
	isSmall     bool
}

func NewCave(input string) *Cave {
	return &Cave{
		ID:      input,
		isSmall: strings.ToLower(input) == input && input != "start" && input != "end",
	}
}

func (c Cave) IsStart() bool {
	return c.ID == "start"
}

func (c Cave) IsEnd() bool {
	return c.ID == "end"
}

func ParseCaves(inputs []string) (map[string]*Cave, error) {
	caves := map[string]*Cave{}
	for _, connection := range inputs {
		raw := strings.Split(connection, "-")
		if len(raw) != 2 {
			return caves, errors.Errorf("invalid input: %v", connection)
		}

		c1 := NewCave(raw[0])
		if _, exists := caves[c1.ID]; !exists {
			caves[c1.ID] = c1
		} else {
			c1 = caves[c1.ID]
		}

		c2 := NewCave(raw[1])
		if _, exists := caves[c2.ID]; !exists {
			caves[c2.ID] = c2
		} else {
			c2 = caves[c2.ID]
		}

		c1.Connections = append(c1.Connections, c2)
		c2.Connections = append(c2.Connections, c1)
	}

	return caves, nil
}

func CanVisitOnce(cave *Cave, prev []*Cave) bool {
	if cave == nil || cave.IsStart() {
		return false
	}

	if !cave.isSmall {
		return true
	}

	for _, c := range prev {
		if c.ID == cave.ID {
			return false
		}
	}

	return true
}

func CanVisitTwice(cave *Cave, prev []*Cave) bool {
	if cave == nil || cave.IsStart() {
		return false
	}

	if !cave.isSmall {
		return true
	}

	hasVisitedSmallCaveTwice := false
	network := map[string]int{}
	for _, c := range append(prev, cave) {
		if _, exists := network[c.ID]; !exists {
			network[c.ID] = 0
		}
		network[c.ID]++

		if c.isSmall && network[c.ID] > 2 {
			return false
		}

		if c.isSmall && network[c.ID] == 2 {
			if hasVisitedSmallCaveTwice {
				return false
			}
			hasVisitedSmallCaveTwice = true
		}
	}

	return true
}

func explode(curr *Cave, prev []*Cave, canVisit func(cave *Cave, prev []*Cave) bool) [][]*Cave {
	if curr.IsEnd() {
		return [][]*Cave{append(prev, curr)}
	}

	var paths [][]*Cave
	for _, cave := range curr.Connections {
		if canVisit(cave, append(prev, curr)) {
			paths = append(paths, explode(cave, append(prev, curr), canVisit)...)
		}
	}

	return paths
}

func p1(inputs []string) (int, error) {
	caves, err := ParseCaves(inputs)
	if err != nil {
		return 0, err
	}

	var paths [][]*Cave
	for _, cave := range caves["start"].Connections {
		paths = append(paths, explode(cave, []*Cave{caves["start"]}, CanVisitOnce)...)
	}

	cnt := 0
	for _, path := range paths {
		for _, cave := range path {
			if cave.isSmall {
				cnt++
				break
			}
		}
	}

	return cnt, nil
}

func p2(inputs []string) (int, error) {
	caves, err := ParseCaves(inputs)
	if err != nil {
		return 0, err
	}

	var paths [][]*Cave
	for _, cave := range caves["start"].Connections {
		paths = append(paths, explode(cave, []*Cave{caves["start"]}, CanVisitTwice)...)
	}

	return len(paths), nil
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
