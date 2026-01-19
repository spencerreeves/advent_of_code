package main

import (
	"strconv"
	"strings"
)

type MathOp int8

const (
	Add MathOp = iota
	Multiply
)

type MathProblem struct {
	Args   []int
	Action MathOp
	Answer int
}

func P6GetFirstStar(in []string) int {
	problems := ParseMathProblems(in)

	grandTotal := 0
	for _, problem := range problems {
		answer := 0
		if problem.Action == Multiply {
			answer = 1
		}

		for _, arg := range problem.Args {
			if problem.Action == Add {
				answer += arg
			} else {
				answer *= arg
			}
		}

		problem.Answer = answer
		grandTotal += answer
	}

	return grandTotal
}

func P6GetSecondStar(in []string) int {
	problems := ParseMathProblemsV2(in)

	grandTotal := 0
	for _, problem := range problems {
		answer := 0
		if problem.Action == Multiply {
			answer = 1
		}

		for _, arg := range problem.Args {
			if problem.Action == Add {
				answer += arg
			} else {
				answer *= arg
			}
		}

		problem.Answer = answer
		grandTotal += answer
	}

	return grandTotal
}

func ParseMathProblems(inputs []string) []MathProblem {
	var problems []MathProblem

	for _, row := range inputs {
		for idx, arg := range strings.Fields(row) {
			if arg == "*" {
				problems[idx].Action = Multiply
				continue
			} else if arg == "+" {
				problems[idx].Action = Add
				continue
			}

			value, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}

			if len(problems) < idx+1 {
				problems = append(problems, MathProblem{
					Args: []int{value},
				})
			} else {
				problems[idx].Args = append(problems[idx].Args, value)
			}
		}
	}

	return problems
}

func ParseMathProblemsV2(inputs []string) []MathProblem {
	// Consider the input as a matrix, with the last row being the operators.
	// Operators always start at the first index of the problem, and therefore they are used
	// as the leftmost index.

	if len(inputs) < 1 {
		return nil
	}

	var problems []MathProblem
	problem := MathProblem{}

	for idx := range len(inputs[0]) {
		slice := GetSlice(inputs, idx)

		// Store current problem and prep new problem
		isBreak := true
		for _, char := range slice {
			if char != ' ' {
				isBreak = false
				break
			}
		}

		if isBreak {
			problems = append(problems, problem)
			problem = MathProblem{}
			continue
		}

		// Construct the number
		num := 0
		for _, char := range slice[:len(slice)-1] {
			if char == ' ' {
				continue
			}

			num = num*10 + int(char-'0')
		}
		problem.Args = append(problem.Args, num)

		// Determine the operator
		if slice[len(slice)-1] == '+' {
			problem.Action = Add
		} else if slice[len(slice)-1] == '*' {
			problem.Action = Multiply
		}
	}

	return append(problems, problem)
}

func GetSlice(inputs []string, idx int) []uint8 {
	values := make([]uint8, len(inputs))
	for valIdx, in := range inputs {
		values[valIdx] = in[idx]
	}

	return values
}
