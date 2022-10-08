package main

import (
	"fmt"
	"github.com/spencerreeves/advent_of_code/2020/day1"
	"github.com/spencerreeves/advent_of_code/2020/day2"
	"github.com/spencerreeves/advent_of_code/2020/day3"
	"github.com/spencerreeves/advent_of_code/2020/day4"
	"github.com/spencerreeves/advent_of_code/2020/day5"
	"log"
)

func main() {
	log.Printf("\n\n*****     Advent of Code     *****     \n")
	fmt.Printf("**  Day 1 **\n")
	trackFunc("Day 1, Problem 1", day1.Problem1)
	trackFunc("Day 1, Problem 2", day1.Problem2)

	fmt.Printf("**  Day 2  **\n")
	trackFunc("Day 2, Problem 1", day2.Problem1)
	trackFunc("Day 2, Problem 2", day2.Problem2)

	fmt.Printf("**  Day 3 **\n")
	trackFunc("Day 3, Problem 1", day3.Problem1)
	trackFunc("Day 3, Problem 2", day3.Problem2)

	fmt.Printf("**  Day 4 **\n")
	trackFunc("Day 4, Problem 1", day4.Problem1)
	trackFunc("Day 4, Problem 2", day4.Problem2)

	fmt.Printf("** Day 5 **\n")
	trackFunc("Day 5, Problem 1", day5.Problem1)
	trackFunc("Day 5, Problem 2", day5.Problem2)
}
