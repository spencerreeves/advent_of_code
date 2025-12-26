package main

import (
	"fmt"
	"strconv"
	"strings"
)

func P2GetFirstStar(in []string) int {
	bounds, err := GetBounds(in)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}

	var invalidProductIDs []int
	for _, b := range bounds {
		invalidProductIDs = append(GetMirrorRepeats(b.Lower, b.Upper), invalidProductIDs...)
	}

	answer := 0
	for _, productID := range invalidProductIDs {
		answer += productID
	}

	return answer
}

func P2GetSecondStar(in []string) int {
	bounds, err := GetBounds(in)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0
	}

	var invalidProductIDs []int
	for _, b := range bounds {
		for i := b.Lower; i <= b.Upper; i++ {
			invalid := IsInvalidID(i)
			if invalid {
				invalidProductIDs = append(invalidProductIDs, i)
			}
		}
	}

	answer := 0
	for _, productID := range invalidProductIDs {
		answer += productID
	}

	return answer
}

type Bounds struct {
	Lower int
	Upper int
}

func GetBounds(in []string) ([]Bounds, error) {
	var bounds []Bounds

	for _, rr := range in {
		b := strings.Split(strings.TrimSpace(rr), "-")
		if len(b) != 2 {
			return nil, fmt.Errorf("invalid bounds for problem 2, bounds: %v", rr)
		}

		lower, err := strconv.Atoi(b[0])
		if err != nil {
			return nil, fmt.Errorf("invalid lower bounds for problem 2, bounds: %v", bounds[0])
		}

		upper, err := strconv.Atoi(b[1])
		if err != nil {
			fmt.Printf("Invalid upper bounds for problem 2 part 1. Bounds: %v\n", bounds[1])
		}

		bounds = append(bounds, Bounds{lower, upper})
	}

	return bounds, nil
}

// GetMirrorRepeats Rules:
//  1. A repeated number must have an even number of digits (middle digit cannot be used on both sides)
//  2. Repeats are only checked within the range provided
func GetMirrorRepeats(lower, upper int) []int {
	// Deconstruct lower bound into the first segment
	segment := strconv.Itoa(lower)
	if len(segment) > 1 {
		segment = segment[:len(segment)/2]
	}

	var repeats []int
	for {
		repeat, err := strconv.Atoi(segment + segment)
		if err != nil {
			fmt.Printf("Error deconstructing segment in problem 2. Segment: %v, Err: %v\n", segment, err)
			return nil
		}

		if repeat <= upper {
			if repeat >= lower {
				repeats = append(repeats, repeat)
			}
		} else {
			break
		}

		t, err := strconv.Atoi(segment)
		if err != nil {
			fmt.Printf("Error constructing segment in problem 2. Segment: %v, Err: %v\n", segment, err)
		}

		segment = strconv.Itoa(t + 1)
	}

	return repeats
}

// IsValidPattern determines if the given ID is invalid based on the following criteria
// 1. ID must be made of ONLY 1 repeating pattern
// 2. Pattern must repeat at least twice
func IsValidPattern(id, pattern string) bool {
	if len(pattern) > len(id) || len(id)%len(pattern) != 0 {
		return false
	}

	for i := 0; i+len(pattern) <= len(id); i += len(pattern) {
		chunk := id[i : i+len(pattern)]
		if chunk != pattern {
			return false
		}
	}

	return true
}

// IsInvalidID Starts with the first character as the pattern and checks if the pattern repeats, short-circuiting when
// the pattern does not repeat. Add the next character from the input and repeat the previous step until we reach the
// halfway point of the input.
func IsInvalidID(id int) bool {
	if id < 10 {
		return false
	}

	idStr := strconv.Itoa(id)

	var pattern string
	for i := 0; i < len(idStr)/2; i++ {
		pattern += string(idStr[i])

		if IsValidPattern(idStr, pattern) {
			return true
		}
	}

	return false
}
