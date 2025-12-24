package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetFirstStar(input []string) int {
	// Create a liked list to represent our dial
	dial := NewDial(99, 50)
	answer := 0

	for _, cmd := range input {
		stepCnt, err := strconv.Atoi(cmd[1:])
		if err != nil {
			fmt.Printf("Encountered error parsing problem 1's input. Command: %v, Error: %v", cmd, err)
			return 0
		}

		if strings.HasPrefix(cmd, "L") {
			dial.Turn(-1 * stepCnt)
		} else {
			dial.Turn(stepCnt)
		}

		if dial.CurrNode.Value == 0 {
			answer++
		}
	}

	return answer
}

func GetSecondStar(input []string) int {
	// Create a liked list to represent our dial
	dial := NewDial(99, 50)
	answer := 0

	for _, cmd := range input {
		stepCnt, err := strconv.Atoi(cmd[1:])
		if err != nil {
			fmt.Printf("Encountered error parsing problem 1's input. Command: %v, Error: %v", cmd, err)
			return 0
		}

		if strings.HasPrefix(cmd, "L") {
			answer += (stepCnt + (100-dial.CurrNode.Value)%100) / 100 // size of the dial
			stepCnt = -1 * stepCnt
		} else {
			answer += (stepCnt + dial.CurrNode.Value) / 100 // size of the dial
		}

		dial.Turn(stepCnt)
	}

	return answer
}

type Dial struct {
	CurrNode Node
}

type Node struct {
	Value    int
	PrevNode *Node
	NextNode *Node
}

func NewDial(size int, startAt int) *Dial {
	if size <= 0 {
		return nil
	}

	// Set size to size + 1 to accommodate the 0
	size = size + 1

	if startAt < 0 {
		startAt = (-1 * startAt) % size
	}

	var firstNode, prevNode, startNode *Node

	for i := range size {
		node := &Node{
			Value:    i,
			PrevNode: prevNode,
		}

		if i == 0 {
			firstNode = node
		}

		if i == startAt {
			startNode = node
		}

		if prevNode != nil {
			prevNode.NextNode = node
		}

		prevNode = node
	}

	prevNode.NextNode = firstNode
	firstNode.PrevNode = prevNode

	return &Dial{
		CurrNode: *startNode,
	}
}

func (d *Dial) Turn(in int) {
	steps := in
	if steps < 0 {
		steps = -1 * steps
	}

	node := d.CurrNode
	for _ = range steps {
		if in < 0 {
			node = *node.PrevNode
		} else {
			node = *node.NextNode
		}
	}

	d.CurrNode = node
}
