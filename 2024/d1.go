package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const InputFile = "./2024/d1_input.txt"

type D1Location struct {
	Left  int
	Right int
}

func (d *D1Location) UnmarshalText(bytes []byte) error {
	vals := strings.Split(string(bytes), "   ")
	if len(vals) != 2 {
		return fmt.Errorf("invalid Day 1 input")
	}

	var err error
	if d.Left, err = strconv.Atoi(vals[0]); err != nil {
		return fmt.Errorf("invalid Day 1 input (left): %v", vals[0])
	}

	if d.Right, err = strconv.Atoi(vals[1]); err != nil {
		return fmt.Errorf("invalid Day 1 input (right): %v", vals[1])
	}

	return nil
}

func ReduceInsertSort[T cmp.Ordered, K any](ts []T, k K, reducer func(K) T) []T {
	t := reducer(k)
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func D1P1() string {
	in, err := NewInput[D1Location](InputFile)
	if err != nil {
		return "Err: Invalid input"
	}

	inputs, err := in.All()
	if err != nil {
		return "Err: Failed to read inputs"
	}

	var left []int
	for _, l := range inputs {
		left = ReduceInsertSort(left, l, func(d *D1Location) int { return d.Left })
	}

	var right []int
	for _, l := range inputs {
		right = ReduceInsertSort(right, l, func(d *D1Location) int { return d.Right })
	}

	var ans int
	for idx := range left {
		i := left[idx] - right[idx]
		if i < 0 {
			i = i * -1
		}
		ans += i
	}

	return fmt.Sprint(ans)
}
