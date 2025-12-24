package main

import (
	"fmt"
	"github.com/spencerreeves/advent_of_code/2023/day5"
	"github.com/spencerreeves/advent_of_code/2023/input"
	"github.com/spencerreeves/advent_of_code/2023/solutions"
)

func main() {
	in := input.ReadAll(fmt.Sprintf("./2023/inputs/day_%v.txt", 1))
	fmt.Printf("Day 1: %v | %v \n", solutions.D1P1(in), solutions.D1P2(in))

	in = input.ReadAll(fmt.Sprintf("./2023/inputs/day_%v.txt", 2))
	fmt.Printf("Day 2: %v | %v \n", solutions.D2P1(in), solutions.D2P2(in))

	in = input.ReadAll(fmt.Sprintf("./2023/inputs/day_%v.txt", 3))
	fmt.Printf("Day 3: %v | %v \n", solutions.D3P1(in), solutions.D3P2(in))

	in = input.ReadAll(fmt.Sprintf("./2023/inputs/day_%v.txt", 4))
	fmt.Printf("Day 4: %v | %v \n", solutions.D4P1(in), solutions.D4P2(in))

	fmt.Printf("Day 5: %v | %v \n", day5.Part1(), day5.Part2())
}
