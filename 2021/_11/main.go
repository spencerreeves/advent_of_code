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
	Day       = "11"
	InputFile = "./2021/_" + Day + "/input.txt"
)

type Octo struct {
	Row    int
	Col    int
	Energy int
}

// Step increases the energy level of the octopus
func (o *Octo) Step() {
	if o.Energy == 9 {
		o.Energy = 0
	} else {
		o.Energy++
	}
}

func (o *Octo) Energize() bool {
	if o.Energy == 9 {
		o.Energy = 0
		return true
	} else if o.Energy > 0 {
		o.Energy++
	}

	return false
}

type Grid struct {
	octopie [][]*Octo
}

func NewGrid(inputs []string) *Grid {
	pies := make([][]*Octo, len(inputs))
	for row, line := range inputs {
		pies[row] = make([]*Octo, len(line))
		for col, loc := range line {
			val, err := strconv.Atoi(string(loc))
			if err != nil {
				return nil
			}

			pies[row][col] = &Octo{Row: row, Col: col, Energy: val}
		}
	}

	return &Grid{
		octopie: pies,
	}
}

func (g Grid) Step() int {
	var cnt int
	var flashed []*Octo
	for _, pies := range g.octopie {
		for _, octo := range pies {
			if octo.Step(); octo.Energy == 0 {
				flashed = append(flashed, octo)
				cnt++
			}
		}
	}

	for len(flashed) > 0 {
		for _, octo := range g.GetSurroundingOctopie(flashed[0].Row, flashed[0].Col) {
			if octo.Energize() {
				flashed = append(flashed, octo)
				cnt++
			}
		}

		flashed = flashed[1:]
	}

	return cnt
}

func (g Grid) Get(row, col int) *Octo {
	if !g.isValidCoord(row, col) {
		return nil
	}
	return g.octopie[row][col]
}

func (g Grid) isValidCoord(row, col int) bool {
	return row >= 0 && col >= 0 && row < len(g.octopie) && col < len(g.octopie[0])
}

func (g Grid) GetSurroundingOctopie(row, col int) []*Octo {
	var pies []*Octo
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if g.isValidCoord(r, c) && g.Get(row, col) != g.Get(r, c) {
				pies = append(pies, g.Get(r, c))
			}
		}
	}

	return pies
}

func (g Grid) ToString() string {
	str := ""
	for _, pies := range g.octopie {
		for _, octo := range pies {
			if octo.Energy == 0 {
				str += fmt.Sprintf("%v%v%v", "\u001B[36m", octo.Energy, "\u001B[0m")
			} else {
				str += fmt.Sprintf("%v", octo.Energy)
			}
		}
		str += "\n"
	}

	return str
}

func p1(inputs []string) (int, error) {
	grid := NewGrid(inputs)
	if grid == nil {
		return 0, errors.New("invalid input")
	}

	flashed := 0
	for i := 0; i < 100; i++ {
		flashed += grid.Step()
		// println(grid.ToString())
	}

	return flashed, nil
}

func p2(inputs []string) (int, error) {
	grid := NewGrid(inputs)
	if grid == nil {
		return 0, errors.New("invalid input")
	}

	var step int
	for step = 1; grid.Step() != 100; step++ {
	}

	return step, nil
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
