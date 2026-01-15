package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in := ReadAll("./2025/p1_input.txt", "\n")
	fmt.Printf("Problem 1: %12d | %16d\n", P1GetFirstStar(in), P1GetSecondStar(in))

	in = ReadAll("./2025/p2_input.txt", ",")
	fmt.Printf("Problem 2: %12d | %16d\n", P2GetFirstStar(in), P2GetSecondStar(in))

	in = ReadAll("./2025/p3_input.txt", "\n")
	fmt.Printf("Problem 3: %12d | %16d\n", P3GetFirstStar(in), P3GetSecondStar(in))

	in = ReadAll("./2025/p4_input.txt", "\n")
	fmt.Printf("Problem 4: %12d | %16d\n", P4GetFirstStar(in), P4GetSecondStar(in))

	in = ReadAll("./2025/p5_input.txt", "\n")
	fmt.Printf("Problem 5: %12d | %16d\n", P5GetFirstStar(in), P5GetSecondStar(in))
}

func ReadAll(path string, split string) []string {
	dd, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("invalid path %v: %w", path, err))
	}

	return strings.Split(string(dd), split)
}
