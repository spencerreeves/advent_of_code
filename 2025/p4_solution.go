package main

import "iter"

func P4GetFirstStar(in []string) int {
	grid := NewGrid(in)

	accessibleBlocks := 0
	for block := range grid.All() {
		if grid.IsAccessible(block) {
			accessibleBlocks++
		}
	}

	return accessibleBlocks
}

func P4GetSecondStar(in []string) int {
	grid := NewGrid(in)

	removedBlocks := 0
	for {
		var blocksToRemove []Block

		for block := range grid.All() {
			if grid.IsAccessible(block) {
				blocksToRemove = append(blocksToRemove, block)
			}
		}

		if len(blocksToRemove) == 0 {
			break
		}

		removedBlocks += len(blocksToRemove)
		for _, block := range blocksToRemove {
			grid.RemovePaper(block)
		}
	}

	return removedBlocks
}

const PaperRoll = '@'
const AccessibilityThreshold = 4

type Coord struct {
	X, Y int
}

type Block struct {
	Coords     Coord
	IsOccupied bool
}

type Grid struct {
	data [][]Block
}

func NewGrid(in []string) *Grid {
	data := make([][]Block, len(in))

	for y := range in {
		for x := range in[y] {
			data[y] = append(data[y], Block{
				Coords:     Coord{x, y},
				IsOccupied: in[y][x] == PaperRoll,
			})
		}
	}

	return &Grid{data: data}
}

func (g *Grid) IsAccessible(block Block) bool {
	if !block.IsOccupied {
		return false
	}

	adjacentBlocks := []Coord{
		// Top Row
		{block.Coords.X - 1, block.Coords.Y - 1},
		{block.Coords.X, block.Coords.Y - 1},
		{block.Coords.X + 1, block.Coords.Y - 1},
		// Same Row
		{block.Coords.X - 1, block.Coords.Y},
		{block.Coords.X + 1, block.Coords.Y},
		// Bottom Row
		{block.Coords.X - 1, block.Coords.Y + 1},
		{block.Coords.X, block.Coords.Y + 1},
		{block.Coords.X + 1, block.Coords.Y + 1},
	}

	occupiedAdjacentBlocks := 0
	for _, b := range adjacentBlocks {
		if g.IsValidCoordinate(b.X, b.Y) && g.data[b.Y][b.X].IsOccupied {
			occupiedAdjacentBlocks++
		}
	}

	if occupiedAdjacentBlocks < AccessibilityThreshold {
		return true
	}

	return false
}

func (g *Grid) RemovePaper(block Block) {
	g.data[block.Coords.Y][block.Coords.X].IsOccupied = false
}

func (g *Grid) IsValidCoordinate(x, y int) bool {
	if y >= 0 && len(g.data) > y && x >= 0 && len(g.data[y]) > x {
		return true
	}

	return false
}

func (g *Grid) All() iter.Seq[Block] {
	return func(yield func(Block) bool) {
		for y := range g.data {
			for x := range g.data[y] {
				if !yield(g.data[y][x]) {
					return
				}
			}
		}
	}
}
