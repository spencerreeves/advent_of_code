package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in := ReadAll("./2025/p1_input.txt")
	fmt.Printf("Problem 1: %10d | %10d\n", GetFirstStar(in), GetSecondStar(in))
}

func ReadAll(path string) []string {
	dd, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("invalid path %v: %w", path, err))
	}

	return strings.Split(string(dd), "\n")
}
