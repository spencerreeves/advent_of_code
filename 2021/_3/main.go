package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const InputFile = "./2021/_3/input.txt"

func p1(inputs []string) (int64, error) {
	heat := make([]map[bool]int, len(inputs[0]))
	for _, v := range inputs {
		for index, c := range v {
			if heat[index] == nil {
				heat[index] = make(map[bool]int)
			}

			heat[index][c == '0']++
		}
	}

	var gamma, epsilon string
	for _, v := range heat {
		if v[true] > v[false] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	var e, g int64
	var err error
	if g, err = strconv.ParseInt(gamma, 2, 64); err != nil {
		return 0, err
	}
	if e, err = strconv.ParseInt(epsilon, 2, 64); err != nil {
		return 0, err
	}

	return e * g, nil
}

func p2(inputs []string) (int64, error) {
	oxyrngs := inputs
	for sigDig := 0; len(oxyrngs) > 1; sigDig++ {
		ll := make(map[bool][]string, 2)
		for _, v := range oxyrngs {
			ll[v[sigDig] == '0'] = append(ll[v[sigDig] == '0'], v)
		}

		oxyrngs = ll[false]
		if len(ll[true]) > len(ll[false]) {
			oxyrngs = ll[true]
		}
	}

	co2rngs := inputs
	for sigDig := 0; len(co2rngs) > 1; sigDig++ {
		ll := make(map[bool][]string, 2)
		for _, v := range co2rngs {
			ll[v[sigDig] == '0'] = append(ll[v[sigDig] == '0'], v)
		}

		co2rngs = ll[true]
		if len(ll[false]) < len(ll[true]) {
			co2rngs = ll[false]
		}
	}

	var oxy, co2 int64
	var err error
	if oxy, err = strconv.ParseInt(oxyrngs[0], 2, 64); err != nil {
		return 0, err
	}
	if co2, err = strconv.ParseInt(co2rngs[0], 2, 64); err != nil {
		return 0, err
	}

	return oxy * co2, nil
}

func main() {
	dd, err := os.ReadFile(InputFile)
	if err != nil {
		log.Panic(err, "\tinput file")
	}

	input := strings.Split(string(dd), "\n")

	ans1, err := p1(input)
	if err != nil {
		log.Panicf("%v, day 3, part 1", err)
	}
	log.Printf("Day 3, Part 1: %v", ans1)

	ans2, err := p2(input)
	if err != nil {
		log.Panicf("%v, day 3, part 2", err)
	}

	log.Printf("Day 3, Part 2: %v", ans2)
}
