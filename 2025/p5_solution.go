package main

import (
	"fmt"
	"strconv"
	"strings"
)

func P5GetFirstStar(in []string) int {
	ranges, ingredientIDs := ParseDB(in)

	freshIngredients := 0
	for _, id := range ingredientIDs {
		for _, rng := range ranges {
			if id >= rng.Low && id <= rng.High {
				freshIngredients++
				break
			}
		}
	}

	return freshIngredients
}

func P5GetSecondStar(in []string) int {
	ranges, _ := ParseDB(in)

	if len(ranges) == 0 {
		return 0
	}

	ll := &P5Node{
		Value: ranges[0],
	}

	for _, rng := range ranges[1:] {
		ll.Add(&P5Node{Value: rng})
		ll = ll.GetFirst()
	}

	freshIngredients := 0
	for n := ll; n != nil; n = n.Next {
		freshIngredients += (n.Value.High - n.Value.Low) + 1 // Add one since the ranges are inclusive
	}

	return freshIngredients
}

type Range struct {
	Low, High int
}

func ParseDB(in []string) ([]Range, []int) {
	var ranges []Range
	var ids []int

	for _, line := range in {
		if strings.Contains(line, "-") {
			segments := strings.Split(line, "-")
			if len(segments) != 2 {
				fmt.Println("Error parsing database:", line)
				return nil, nil
			}

			low, err := strconv.Atoi(segments[0])
			if err != nil {
				fmt.Println("Error parsing low:", line)
				return nil, nil
			}

			high, err := strconv.Atoi(segments[1])
			if err != nil {
				fmt.Println("Error parsing high:", line)
				return nil, nil
			}

			ranges = append(ranges, Range{Low: low, High: high})

			continue
		}

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil
		}

		ids = append(ids, id)
	}

	return ranges, ids
}

type P5Node struct {
	Prev, Next *P5Node
	Value      Range
}

func (n *P5Node) GetFirst() *P5Node {
	if n.Prev != nil {
		return n.Prev.GetFirst()
	}

	return n
}

func (n *P5Node) Add(o *P5Node) {
	// Do nothing, node is null
	if o == nil {
		return
	}

	// Since the nodes are sorted, if this node is strickly lower than the current range
	// add it to just before this node
	if o.Value.High < n.Value.Low {
		if n.Prev == nil {
			o.Next = n
			n.Prev = o
		} else {
			n.Prev.Next = o
			o.Prev = n.Prev
			o.Next = n
			n.Prev = o
		}

		return
	}

	// When strickly higher than the current range, add this to the next node.
	if o.Value.Low > n.Value.High {
		if n.Next == nil {
			n.Next = o
			o.Prev = n
		} else {
			n.Next.Add(o)
		}

		return
	}

	// If we got to this point, then the ranges overlap. Therefore, we combine them
	// remove the current node, and add the new range to the next node
	low := n.Value.Low
	if o.Value.Low < low {
		low = o.Value.Low
	}

	high := n.Value.High
	if o.Value.High > high {
		high = o.Value.High
	}

	// If we are at the end of the list, overwrite the current node
	if n.Next == nil {
		n.Value = Range{Low: low, High: high}
		return
	}

	n.Next.Prev = n.Prev
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}

	n.Next.Add(&P5Node{
		Value: Range{Low: low, High: high},
	})
}
