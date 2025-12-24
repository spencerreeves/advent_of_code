package solutions

import (
	"slices"
	"strconv"
	"strings"
)

type Part struct {
	Length int
	Serial int
	Symbol string
	X      int
	Y      int
}

func NewPart(value string, xPos, yPos int) Part {
	serial, err := strconv.Atoi(value)

	var symbol string
	if err != nil {
		symbol = value
	}

	return Part{
		Length: len(value),
		Serial: serial,
		Symbol: symbol,
		X:      xPos,
		Y:      yPos,
	}
}

func (p Part) Equals(part Part) bool {
	return p.X == part.X && p.Y == part.Y
}

type Engine struct {
	parts map[int]map[int]*Part
}

func NewEngine() Engine {
	return Engine{
		parts: map[int]map[int]*Part{},
	}
}

func (e Engine) AddPart(part Part) *Part {
	if existingPart := e.GetPart(part.X, part.Y); existingPart != nil {
		return existingPart
	}

	// Create copy of part
	p := part

	for xPos := p.X; xPos < p.X+p.Length; xPos++ {
		if e.parts[p.Y] == nil {
			e.parts[p.Y] = map[int]*Part{}
		}

		e.parts[p.Y][xPos] = &p
	}

	return &p
}

func (e Engine) GetPart(xPos, yPos int) *Part {
	if row, rowExists := e.parts[yPos]; rowExists {
		if part, partExists := row[xPos]; partExists {
			return part
		}
	}

	return nil
}

func (e Engine) ListParts() []*Part {
	parts := []*Part{}
	added := map[*Part]bool{}

	for _, row := range e.parts {
		for _, part := range row {
			if _, exists := added[part]; !exists {
				added[part] = true
				parts = append(parts, part)
			}
		}
	}

	return parts
}

func (e Engine) GetAdjacentParts(part Part) []*Part {
	var parts []*Part
	for yPos := part.Y - 1; yPos <= part.Y+1; yPos++ {
		for xPos := part.X - 1; xPos <= part.X+part.Length; xPos++ {
			if adjacent := e.GetPart(xPos, yPos); adjacent != nil && !adjacent.Equals(part) && !slices.Contains(parts, adjacent) {
				parts = append(parts, e.GetPart(xPos, yPos))
			}
		}
	}

	return parts
}

func parseEngineDiagram(in []string) Engine {
	engine := NewEngine()

	for yPos, line := range in {
		xPos := 0
		value := ""

		for _, c := range line {
			segment := string(c)

			if strings.Contains("0123456789", segment) {
				value += segment
				continue
			}

			if len(value) > 0 {
				engine.AddPart(NewPart(value, xPos, yPos))
				xPos += len(value)
				value = ""
			}

			engine.AddPart(NewPart(segment, xPos, yPos))
			xPos += len(segment)
		}

		if len(value) > 0 {
			engine.AddPart(NewPart(value, xPos, yPos))
			xPos += len(value)
			value = ""
		}
	}

	return engine
}

func D3P1(input []string) int {
	engine := parseEngineDiagram(input)

	partNumberSum := 0
	for _, part := range engine.ListParts() {
		if part.Serial > 0 {
			for _, adjacent := range engine.GetAdjacentParts(*part) {
				if adjacent.Symbol != "" && adjacent.Symbol != "." {
					partNumberSum += part.Serial
					break
				}
			}
		}
	}

	return partNumberSum
}

func D3P2(input []string) int {
	engine := parseEngineDiagram(input)

	gearRatioSum := 0
	for _, part := range engine.ListParts() {
		if part.Symbol == "*" {
			adjacentNumberedParts := 0
			gearRatio := 1

			for _, adjacent := range engine.GetAdjacentParts(*part) {
				if adjacent.Serial > 0 {
					adjacentNumberedParts += 1
					gearRatio *= adjacent.Serial
				}
			}

			if adjacentNumberedParts == 2 {
				gearRatioSum += gearRatio
			}
		}
	}

	return gearRatioSum
}
