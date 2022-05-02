package challenge3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFile1 = "challenge3/input1.txt"
const inputFile2 = "challenge3/input2.txt"
const inputFile = "challenge3/input.txt"

func panicIfError(err error) bool {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return false
}

func Abs(value int) int {
	if value > 0 {
		return value
	}

	return value * -1
}

func distanceToOrigin(point Point) int {
	return Abs(point.x) + Abs(point.y)
}

func inputToLine(input string, origin Point, steps int) (Line, error) {
	direction := input[0]
	distance, err := strconv.Atoi(input[1:])
	panicIfError(err)

	steps += distance

	switch direction {
	case 'R':
		return Line{origin, Point{origin.x + distance, origin.y}, steps}, nil
	case 'L':
		return Line{origin, Point{origin.x - distance, origin.y}, steps}, nil
	case 'U':
		return Line{origin, Point{origin.x, origin.y + distance}, steps}, nil
	case 'D':
		return Line{origin, Point{origin.x, origin.y - distance}, steps}, nil
	}

	return Line{}, errors.New("bad input")
}

func getWires(inputFile string) [][]Line {
	file, err := os.Open(inputFile)
	panicIfError(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	var wires [][]Line

	for {
		steps := 0
		origin := Point{0, 0}
		input, err := reader.ReadString('\n')
		var wire []Line

		if input != "" {
			for _, direction := range strings.Split(strings.TrimSuffix(input, "\n"), ",") {
				segment, err := inputToLine(direction, origin, steps)
				panicIfError(err)

				wire = append(wire, segment)
				origin = Point{segment.p2.x, segment.p2.y}
				steps = segment.steps
			}

			wires = append(wires, wire)
		}

		if err == io.EOF {
			break
		}

		panicIfError(err)
	}

	return wires
}

func Run() {
	fmt.Println("Starting challenge 3...")

	index := 0
	wires := getWires(inputFile)
	for i, wire := range wires {
		if len(wire) < len(wires[index]) {
			index = i
		}
	 }

	 leastAmountOfSteps := math.MaxInt64
	 closestIntersectionDistance := math.MaxInt64
	 closestIntersection := Point{0,0}

	 for i, wire := range wires {
	 	if i == index {
	 		continue
		}

		for _, line := range wire {
			for _, segment := range wires[index] {
				//fmt.Println(line, "<---->", segment)
				intersection, steps, err := DoLinesIntersect(line, segment)

				if err == nil {
					distance := distanceToOrigin(intersection)
					fmt.Println(steps, distance, intersection)
					if distance < closestIntersectionDistance && steps < leastAmountOfSteps{
						leastAmountOfSteps = steps
						closestIntersectionDistance = distance
						closestIntersection = intersection
					}
				}
			}
		}
	 }

	fmt.Println(closestIntersection, closestIntersectionDistance, leastAmountOfSteps)
}