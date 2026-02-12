package main

import (
	"strings"
)

func P7GetFirstStar(in []string) int {
	splits, _ := ParseManifold(in)
	return splits
}

func P7GetSecondStar(in []string) int {
	_, timelines := ParseManifold(in)
	return timelines
}

type Beam struct {
	Timelines int
}

func ParseManifold(in []string) (int, int) {
	beamsIn, beamsOut := map[int]Beam{}, map[int]Beam{}
	splitCnt := 0

	// First line will be the starting line
	beamsIn[strings.Index(in[0], "S")] = Beam{1}

	for _, line := range in[1:] {
		for idx, beam := range beamsIn {
			if line[idx] == '.' {
				if b, ok := beamsOut[idx]; ok {
					beam.Timelines += b.Timelines
				}
				beamsOut[idx] = beam
			} else {
				SafeInsert(line, beamsOut, idx-1, beam)
				SafeInsert(line, beamsOut, idx+1, beam)
				splitCnt++
			}
		}
		beamsIn = beamsOut
		beamsOut = make(map[int]Beam)

		// For debugging
		//for idx, beam := range beamsIn {
		//	line = ReplaceAtIdx(line, strconv.Itoa(beam.Timelines), idx)
		//}
		//println(line)
	}

	timelines := 0
	for _, beam := range beamsIn {
		timelines += beam.Timelines
	}

	return splitCnt, timelines
}

func ReplaceAtIdx(str string, replacement string, index int) string {
	newString := str[:index] + replacement
	if index >= len(str) {
		return newString
	}

	return newString + str[index+1:]
}

func SafeInsert(line string, beamsOut map[int]Beam, idx int, beam Beam) {
	if idx < 0 {
		return
	}

	if idx >= len(line) {
		return
	}

	if _, ok := beamsOut[idx]; ok {
		beam.Timelines += beamsOut[idx].Timelines
	}

	beamsOut[idx] = beam
}
