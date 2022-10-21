package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Day       = "9"
	InputFile = "./2021/_" + Day + "/input.txt"
)

func p1(inputs []string) (int, error) {
	var err error
	heightmap := make([][]int, len(inputs))
	for row, line := range inputs {
		heightmap[row] = make([]int, len(line))
		for col, loc := range line {
			heightmap[row][col], err = strconv.Atoi(string(loc))
			if err != nil {
				return 0, errors.Wrap(err, "invalid input")
			}
		}
	}

	risk := 0
	for row, line := range heightmap {
		for col, height := range line {
			isLow := true
			if row != 0 {
				isLow = isLow && height < heightmap[row-1][col]
			}

			if row != len(heightmap)-1 {
				isLow = isLow && height < heightmap[row+1][col]
			}

			if col != 0 {
				isLow = isLow && height < heightmap[row][col-1]
			}

			if col != len(line)-1 {
				isLow = isLow && height < heightmap[row][col+1]
			}

			if isLow {
				risk += 1 + height
			}
		}
	}

	return risk, nil
}

type Coord struct {
	Col   int
	Row   int
	Value int
}

func p2(inputs []string) (int, error) {
	var err error
	heightmap := make([][]int, len(inputs))
	for row, line := range inputs {
		heightmap[row] = make([]int, len(line))
		for col, loc := range line {
			heightmap[row][col], err = strconv.Atoi(string(loc))
			if err != nil {
				return 0, errors.Wrap(err, "invalid input")
			}
		}
	}

	//maps a coordinate the basin it is in
	processed := map[Coord]int{}

	// List of coordinates in each basin
	var basins [][]Coord
	for row, line := range heightmap {
		for col, height := range line {
			coord := Coord{Row: row, Col: col, Value: height}
			if _, ok := processed[coord]; !ok {
				var currBasin, coordsToProcess []Coord
				coordsToProcess = append(coordsToProcess, coord)

				for len(coordsToProcess) != 0 {
					c := coordsToProcess[0]
					coordsToProcess = coordsToProcess[1:]
					if _, ok := processed[c]; ok {
						continue
					}

					if c.Value != 9 {
						processed[c] = len(basins)
						currBasin = append(currBasin, c)
						for _, v := range []Coord{
							{Row: c.Row - 1, Col: c.Col},
							{Row: c.Row + 1, Col: c.Col},
							{Row: c.Row, Col: c.Col - 1},
							{Row: c.Row, Col: c.Col + 1}} {
							if v.Row < 0 || v.Row >= len(heightmap) || v.Col < 0 || v.Col >= len(line) {
								continue
								// omit row
							}

							v.Value = heightmap[v.Row][v.Col]
							if _, ok := processed[v]; !ok {
								coordsToProcess = append(coordsToProcess, v)
							}
						}
					} else {
						processed[c] = -1
					}
				}

				if len(currBasin) > 0 {
					index := len(basins)
					for i, _ := range basins {
						if len(currBasin) > len(basins[i]) {
							index = i
							break
						}
					}

					for _, c := range currBasin {
						processed[c] = index
					}

					basins = append(basins[:index], append([][]Coord{currBasin}, basins[index:]...)...)
				}
			}
		}
	}

	return len(basins[0]) * len(basins[1]) * len(basins[2]), nil
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
