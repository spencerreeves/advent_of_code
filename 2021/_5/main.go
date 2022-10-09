package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile = "./2021/_5/input.txt"

type Matrix struct {
	Lines   []*Line
	HeatMap map[int]map[int]int
}

func NewMatrix(inputs []string, withDiag bool) (m *Matrix, err error) {
	lines := make([]*Line, len(inputs))
	for i, v := range inputs {
		lines[i], err = NewLine(v, withDiag)
		if err != nil {
			return nil, err
		}
	}

	heatMap := map[int]map[int]int{}
	for _, l := range lines {
		for _, c := range l.Intersects {
			if _, ok := heatMap[c.X]; !ok {
				heatMap[c.X] = map[int]int{c.Y: 0}
			}

			heatMap[c.X][c.Y] += 1
		}
	}

	return &Matrix{
		Lines:   lines,
		HeatMap: heatMap,
	}, nil
}

func (m *Matrix) GetIntersections() int {
	ans := 0
	for _, cols := range m.HeatMap {
		for _, intersects := range cols {
			if intersects >= 2 {
				ans++
			}
		}
	}

	return ans
}

type Line struct {
	Start      *Coord
	End        *Coord
	Intersects []*Coord
}

func NewLine(input string, withDiag bool) (line *Line, err error) {
	// input form required x1,y1 -> x2,y2
	line = &Line{}
	rawCoords := strings.Split(input, " -> ")
	if len(rawCoords) != 2 {
		return nil, errors.New("invalid line")
	}

	line.Start, err = NewCoord(rawCoords[0])
	if err != nil {
		return nil, errors.New("invalid start")
	}

	line.End, err = NewCoord(rawCoords[1])
	if err != nil {
		return nil, errors.New("invalid start")
	}

	if line.Start.X == line.End.X {
		high, low := line.End.Y, line.Start.Y
		if high < low {
			high, low = line.Start.Y, line.End.Y
		}

		for i := low; i <= high; i++ {
			line.Intersects = append(line.Intersects, &Coord{X: line.Start.X, Y: i})
		}
	} else if line.Start.Y == line.End.Y {
		high, low := line.End.X, line.Start.X
		if high < low {
			high, low = line.Start.X, line.End.X
		}

		for i := low; i <= high; i++ {
			line.Intersects = append(line.Intersects, &Coord{X: i, Y: line.Start.Y})
		}
	} else if withDiag {
		changeX, changeY := 1, 1
		if line.End.X-line.Start.X < 0 {
			changeX = -1
		}
		if line.End.Y-line.Start.Y < 0 {
			changeY = -1
		}

		for x, y := line.Start.X, line.Start.Y; x != line.End.X+changeX && y != line.End.Y+changeY; x, y = x+changeX, y+changeY {
			line.Intersects = append(line.Intersects, &Coord{X: x, Y: y})
		}
	}

	return line, nil
}

type Coord struct {
	X int
	Y int
}

func NewCoord(input string) (*Coord, error) {
	rawCoords := strings.Split(input, ",")
	if len(rawCoords) != 2 {
		return nil, errors.New("Invalid coord")
	}

	x, err := strconv.Atoi(rawCoords[0])
	if err != nil {
		return nil, errors.New("invalid number")
	}

	y, err := strconv.Atoi(rawCoords[1])
	if err != nil {
		return nil, errors.New("invalid number")
	}

	return &Coord{
		X: x,
		Y: y,
	}, nil
}

func p1(inputs []string) (ans int, err error) {
	matrix, err := NewMatrix(inputs, false)
	if err != nil {
		return 0, errors.Wrap(err, "invalid matrix")
	}

	return matrix.GetIntersections(), nil
}

func p2(inputs []string) (ans int, err error) {
	matrix, err := NewMatrix(inputs, true)
	if err != nil {
		return 0, errors.Wrap(err, "invalid matrix")
	}

	return matrix.GetIntersections(), nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day 5, part 1", err)
	}
	log.Printf("Day 5, Part 1: %v", ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day 5, part 2", err)
	}

	log.Printf("Day 5, Part 2: %v", ans2)
}
