package challenge2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const input = "challenge2/input.txt"

func hasError(err error) bool {
	if err != nil {
		log.Fatal("[Error]: ", err)
		panic(err)
	}

	return false
}

func Run() {
	fmt.Println("Running challenge 2...")
	data, err := ioutil.ReadFile(input)
	hasError(err)

	memory := strings.Split(string(data), ",")

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			memoryCopy := append(memory[:0:0], memory...)
			writeAddress(1, memoryCopy, noun)
			writeAddress(2, memoryCopy, verb)

			if runInstructions(memoryCopy) == 19690720 {
				fmt.Println(noun, verb, 100 * noun + verb)
				return
			}
		}
	}
}

func runInstructions( memory []string) int {
	for pointer := 0; pointer + 4 < len(memory); pointer += 4 {
		if memory[pointer] == "99" {
			return readAddress(0, memory)
		}

		var opCode, addr1, addr2, writeAddr, total int
		getOpCodeData(memory[pointer : pointer + 4], &opCode, &addr1, &addr2, &writeAddr)

		if opCode == 1 {
			total = readAddress(addr1, memory) + readAddress(addr2, memory)
		}
		if opCode == 2 {
			total = readAddress(addr1, memory) * readAddress(addr2, memory)
		}

		writeAddress(writeAddr, memory, total)
	}

	return 0
}

func getOpCodeData(s []string, vars... *int){
	for i, str := range s {
		*vars[i], _ = strconv.Atoi(str)
	}
}

func readAddress(pointer int, memory []string) int {
	data, err := strconv.Atoi(memory[pointer])
	hasError(err)

	return data
}

func writeAddress(pointer int, memory []string, data int) {
	memory[pointer] = strconv.Itoa(data)
}
