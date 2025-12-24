package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

var NumMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func D1P1(input []string) int {
	sum := 0
	for _, line := range input {
		first, last := "", ""
		for _, c := range line {
			val := string(c)
			if _, err := strconv.Atoi(val); err == nil {
				if first == "" {
					first, last = val, val
				} else {
					last = val
				}
			}
		}

		val, _ := strconv.Atoi(fmt.Sprintf("%v%v", first, last))
		sum += val
	}

	return sum
}

func D1P2(input []string) int {
	sum := 0
	for index, line := range input {
		first, last, running := "", "", ""
		for _, c := range line {
			val := string(c)
			if _, err := strconv.Atoi(val); err == nil {
				if first == "" {
					first, last = val, val
				} else {
					last = val
				}
			} else {
				running += val

				for word, strNum := range NumMap {
					if strings.Contains(running, word) {
						if first == "" {
							first, last = strNum, strNum
						} else {
							last = strNum
						}
						running = val
					}
				}
			}
		}

		val, _ := strconv.Atoi(fmt.Sprintf("%v%v", first, last))
		sum += val

		if val != parseValue(line) {
			panic(fmt.Sprintf("Audit failure [index %v]: Expected %v, got %v", index, parseValue(line), val))
		}
	}

	return sum
}

const chars string = "123456789"

func parseValue(input string) int {
	if len(strings.TrimSpace(input)) == 0 {
		return 0
	}
	iFirst := strings.IndexAny(input, chars)
	iLast := strings.LastIndexAny(input, chars)

	first := ""
	last := ""
	if iFirst > -1 {
		first = input[iFirst : iFirst+1]
	}
	if iLast > -1 {
		last = input[iLast : iLast+1]
	}

	spelled := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for k, v := range spelled {
		if i := strings.Index(input, k); i >= 0 && i < iFirst {
			first = v
			iFirst = i
			if last == "" {
				last = v
				iLast = i
			}
		}
		if i := strings.LastIndex(input, k); i >= 0 && i > iLast {
			last = v
			iLast = i
			if first == "" {
				first = v
				iFirst = i
			}
		}
	}

	num, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

	return num
}
