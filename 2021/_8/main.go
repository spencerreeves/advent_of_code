package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Day       = "8"
	InputFile = "./2021/_" + Day + "/input.txt"
)

// diff returns all unique characters in s1, but not s2
func diff(s1, s2 string) string {
	var value string
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) && !strings.ContainsRune(value, r) {
			value += string(r)
		}
	}

	return value
}

// occurrences a map of the number of times a character occurs between s1 and s2
func occurrences(s string) map[int]string {
	cnt := map[rune]int{}
	for _, r := range s {
		if _, ok := cnt[r]; !ok {
			cnt[r] = 0
		}
		cnt[r]++
	}

	value := map[int]string{}
	for k, v := range cnt {
		if _, ok := value[v]; !ok {
			value[v] = ""
		}
		value[v] += string(k)
	}

	return value
}

func first(s string) rune {
	for _, r := range s {
		return r
	}

	return -1
}

type Display struct {
	Mask map[rune]rune
}

func NewDisplay(input string) (*Display, error) {
	inputs, mask := map[int][]string{}, map[rune]rune{'a': 'a', 'b': 'b', 'c': 'c', 'd': 'd', 'e': 'e', 'f': 'f', 'g': 'g'}
	for _, i := range strings.Split(input, " ") {
		if _, ok := inputs[len(i)]; !ok {
			inputs[len(i)] = []string{}
		}

		inputs[len(i)] = append(inputs[len(i)], i)
	}

	// Validate
	if len(inputs[2]) != 1 || len(inputs[3]) != 1 || len(inputs[4]) != 1 || len(inputs[5]) != 3 || len(inputs[6]) != 3 || len(inputs[7]) != 1 {
		return nil, errors.New("invalid input")
	}

	// Segment missing between 2 and 3 length displays is segment A
	eOrb := occurrences(strings.Join(inputs[5], ""))[1]
	bOrd := diff(inputs[4][0], inputs[2][0])
	f := occurrences(strings.ReplaceAll(input, " ", ""))[9]

	mask['a'] = first(diff(inputs[3][0], inputs[2][0]))
	mask['b'] = first(occurrences(eOrb + bOrd)[2])
	mask['c'] = first(diff(inputs[2][0], f))
	mask['e'] = first(diff(eOrb, inputs[4][0]))
	mask['d'] = first(diff(diff(inputs[4][0], inputs[2][0]), string(mask['b'])))
	mask['f'] = first(f)
	mask['g'] = first(diff(inputs[7][0], eOrb+bOrd+inputs[3][0]))

	return &Display{
		Mask: mask,
	}, nil
}

func (d Display) ToInt(segment string) int {
	switch len(segment) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5:
		if strings.ContainsRune(segment, d.Mask['e']) {
			return 2
		}
		if strings.ContainsRune(segment, d.Mask['b']) {
			return 5
		}
		return 3
	case 6:
		if !strings.ContainsRune(segment, d.Mask['d']) {
			return 0
		}
		if !strings.ContainsRune(segment, d.Mask['c']) {
			return 6
		}
		return 9
	case 7:
		return 8
	default:
		return -1
	}
}

func p1(inputs []string) (int, error) {
	cnt := 0
	for _, input := range inputs {
		partial := strings.Split(input, " | ")
		if len(partial) != 2 {
			return 0, errors.New("invalid input")
		}

		d, err := NewDisplay(partial[0])
		if err != nil {
			return 0, errors.Wrap(err, "unable to create display")
		}

		for _, display := range strings.Split(partial[1], " ") {
			if d.ToInt(display) == 1 || d.ToInt(display) == 4 || d.ToInt(display) == 7 || d.ToInt(display) == 8 {
				cnt++
			}
		}
	}
	return cnt, nil
}

func p2(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		partial := strings.Split(input, " | ")
		if len(partial) != 2 {
			return 0, errors.New("invalid input")
		}

		d, err := NewDisplay(partial[0])
		if err != nil {
			return 0, errors.Wrap(err, "unable to create display")
		}

		digits := ""
		for _, display := range strings.Split(partial[1], " ") {
			digits += strconv.Itoa(d.ToInt(display))
		}

		i, _ := strconv.Atoi(digits)
		sum += i
	}
	return sum, nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day %v, part 1", err, Day)
	}
	log.Printf("Day %v, Part 1: %v", Day, ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day %v, part 2", err, Day)
	}

	log.Printf("Day %v, Part 2: %v", Day, ans2)
}
