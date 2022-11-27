package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Day       = "13"
	InputFile = "./2021/_" + Day + "/input.txt"
)

type Fold struct {
	Axis   string
	Offset int
}

func NewFold(input string) (*Fold, error) {
	parsed := strings.Split(input, "=")
	if len(parsed) != 2 {
		return nil, errors.New("fold requires an axis")
	}

	offset, err := strconv.Atoi(parsed[1])
	if err != nil {
		return nil, errors.New("invalid fold offset: " + parsed[1])
	}

	var axis string
	if strings.Contains(parsed[0], "x") {
		axis = "x"
	} else if strings.Contains(parsed[0], "y") {
		axis = "y"
	} else {
		return nil, errors.New("invalid fold axis: " + parsed[0])
	}

	return &Fold{Axis: axis, Offset: offset}, nil
}

type Point struct {
	X int
	Y int
}

func (p Point) ToString() string {
	return fmt.Sprintf("%v,%v", p.X, p.Y)
}

func NewPoint(input string) (*Point, error) {
	parsed := strings.Split(input, ",")
	if len(parsed) != 2 {
		return nil, errors.New("point can only include one comma")
	}

	x, err := strconv.Atoi(parsed[0])
	if err != nil {
		return nil, errors.New("point has invalid X: " + parsed[0])
	}

	y, err := strconv.Atoi(parsed[1])
	if err != nil {
		return nil, errors.New("point has invalid Y: " + parsed[1])
	}

	return &Point{X: x, Y: y}, nil
}

// Transform will return a new point, rotating the Point across the axis. For (vertical) x-axis folds, the point is
// rotated to the left. For (horizontal) y-axis folds, the point is rotated upward.
func (p Point) Transform(fold *Fold) *Point {
	if fold.Axis == "x" {
		if fold.Offset >= p.X {
			return &Point{X: p.X, Y: p.Y}
		}

		return &Point{X: fold.Offset - (p.X - fold.Offset), Y: p.Y}
	}

	if fold.Axis == "y" {
		if fold.Offset >= p.Y {
			return &Point{X: p.X, Y: p.Y}
		}

		return &Point{X: p.X, Y: fold.Offset - (p.Y - fold.Offset)}
	}

	return nil
}

func p1(inputs []string) (int, error) {
	var points []*Point
	var folds []*Fold
	for _, input := range inputs {
		if input == "" {
			continue
		}

		point, err1 := NewPoint(input)
		if err1 == nil {
			points = append(points, point)
			continue
		}

		fold, err2 := NewFold(input)
		if err2 != nil {
			return 0, errors.Wrap(err1, err2.Error())
		}

		folds = append(folds, fold)
	}

	for _, fold := range folds {
		grid := map[string]*Point{}
		var transformedPoints []*Point

		for _, point := range points {
			transformed := point.Transform(fold)
			if _, exists := grid[transformed.ToString()]; !exists {
				grid[transformed.ToString()] = transformed
				transformedPoints = append(transformedPoints, transformed)
			}
		}

		points = transformedPoints
		break
	}

	return len(points), nil
}

func p2(inputs []string) (int, error) {
	var points []*Point
	var folds []*Fold
	for _, input := range inputs {
		if input == "" {
			continue
		}

		point, err1 := NewPoint(input)
		if err1 == nil {
			points = append(points, point)
			continue
		}

		fold, err2 := NewFold(input)
		if err2 != nil {
			return 0, errors.Wrap(err1, err2.Error())
		}

		folds = append(folds, fold)
	}

	for index, fold := range folds {
		grid := map[string]*Point{}
		var transformedPoints []*Point
		maxX, maxY := 0, 0

		for _, point := range points {
			transformed := point.Transform(fold)
			if _, exists := grid[transformed.ToString()]; !exists {
				grid[transformed.ToString()] = transformed
				transformedPoints = append(transformedPoints, transformed)

				if transformed.X > maxX {
					maxX = transformed.X
				}
				if transformed.Y > maxY {
					maxY = transformed.Y
				}
			}
		}

		points = transformedPoints

		if index+1 == len(folds) {
			out := make([][]string, maxX+1)
			for _, point := range points {
				if len(out[point.X]) == 0 {
					out[point.X] = make([]string, maxY+1)
				}

				out[point.X][point.Y] = "#"
			}

			for y, _ := range out[0] {
				for x, _ := range out {
					if out[x] == nil {
						print("|")
						continue
					}

					c := out[x][y]
					if c == "" {
						c = "."
					}
					print(c)
				}
				print("\n")
			}
		}
	}

	return len(points), nil
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
