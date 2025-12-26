package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in := ReadAll("./2025/p1_input.txt", "\n")
	fmt.Printf("Problem 1: %12d | %12d\n", P1GetFirstStar(in), P1GetSecondStar(in))

	in = ReadAll("./2025/p2_input.txt", ",")
	fmt.Printf("Problem 2: %12d | %12d\n", P2GetFirstStar(in), P2GetSecondStar(in))
}

func ReadAll(path string, split string) []string {
	dd, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("invalid path %v: %w", path, err))
	}

	return strings.Split(string(dd), split)
}
