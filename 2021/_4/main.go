package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile = "./2021/_4/input.txt"

type Board struct {
	Keys     map[int]*Coord
	Marked   [][]bool
	Finished bool
}

type Coord struct {
	Row    int
	Col    int
	Value  int
	Marked bool
}

func NewBoard(input []string) (*Board, error) {
	b := Board{
		Keys:   make(map[int]*Coord, 25),
		Marked: [][]bool{make([]bool, 5), make([]bool, 5), make([]bool, 5), make([]bool, 5), make([]bool, 5)},
	}

	for row, line := range input {
		entries := strings.Split(line, " ")

		for col, i := 0, 0; i < len(entries); i++ {
			if entries[i] == "" {
				continue
			}

			num, err := strconv.Atoi(entries[i])
			if err != nil {
				return nil, errors.New("invalid number")
			}

			b.Keys[num] = &Coord{
				Row:    row,
				Col:    col,
				Value:  num,
				Marked: false,
			}

			col++
		}
	}

	return &b, nil
}

// Mark s off the number from the board and returns whether this board has won.
func (b *Board) Mark(num int) bool {
	if v, ok := b.Keys[num]; ok {
		v.Marked = true
		b.Marked[v.Row][v.Col] = true

		b.Finished = (b.Marked[v.Row][0] && b.Marked[v.Row][1] && b.Marked[v.Row][2] && b.Marked[v.Row][3] && b.Marked[v.Row][4]) ||
			(b.Marked[0][v.Col] && b.Marked[1][v.Col] && b.Marked[2][v.Col] && b.Marked[3][v.Col] && b.Marked[4][v.Col])
		return b.Finished
	}

	return false
}

func (b *Board) GetUnmarkedSum() (num int) {
	for _, v := range b.Keys {
		if !v.Marked {
			num += v.Value
		}
	}

	return num
}

func p1(inputs []string) (ans int, err error) {
	// first line are the numbers called
	nums := make([]int, 100)
	for index, val := range strings.Split(inputs[0], ",") {
		nums[index], err = strconv.Atoi(val)
		if err != nil {
			return 0, errors.Wrap(err, "invalid number")
		}
	}

	// second line is thrown away
	// iteratively add new bing boards
	var boards []*Board
	for index := 2; index < len(inputs); index += 6 {
		b, err := NewBoard(inputs[index : index+5])
		if err != nil {
			return 0, errors.Wrap(err, "invalid board")
		}

		boards = append(boards, b)
	}

	for _, num := range nums {
		for _, b := range boards {
			if b.Mark(num) {
				return num * b.GetUnmarkedSum(), nil
			}
		}
	}

	return 0, errors.New("No winner")
}

func p2(inputs []string) (ans int, err error) {
	// first line are the numbers called
	nums := make([]int, 100)
	for index, val := range strings.Split(inputs[0], ",") {
		nums[index], err = strconv.Atoi(val)
		if err != nil {
			return 0, errors.Wrap(err, "invalid number")
		}
	}

	// second line is thrown away
	// iteratively add new bing boards
	var boards []*Board
	for index := 2; index < len(inputs); index += 6 {
		b, err := NewBoard(inputs[index : index+5])
		if err != nil {
			return 0, errors.Wrap(err, "invalid board")
		}

		boards = append(boards, b)
	}

	remainingBoards := len(boards)
	for _, num := range nums {
		for _, b := range boards {
			if b.Finished {
				continue
			}
			if b.Mark(num) {
				remainingBoards--
			}

			if remainingBoards == 0 {
				return num * b.GetUnmarkedSum(), nil
			}
		}
	}

	return 0, errors.New("No winner")
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day 4, part 1", err)
	}
	log.Printf("Day 4, Part 1: %v", ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day 4, part 2", err)
	}

	log.Printf("Day 4, Part 2: %v", ans2)
}
