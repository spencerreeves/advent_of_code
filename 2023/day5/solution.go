package day5

import (
	"container/list"
	"fmt"
	"github.com/spencerreeves/advent_of_code/2023/input"
	"strconv"
	"strings"
)

const InputPath = "./2023/day5/input.txt"

type Range struct {
	LowerBound int
	Size       int
}

func NewRange(lower, size int) Range {
	return Range{
		LowerBound: lower,
		Size:       size,
	}
}

func (r Range) In(target int) bool {
	return target >= r.LowerBound && target <= r.LowerBound+r.Size
}

// ToIndex will get the index of the target in the range. For example, a range of 50 to 51, ToIndex(50) would return 0.
// If the target is not in the range, -1 is returned
func (r Range) ToIndex(target int) int {
	if !r.In(target) {
		return -1
	}

	return target - r.LowerBound
}

// FromIndex provides the value of the index. For example, a range of 50 to 51, FromIndex(0) would return 50.
func (r Range) FromIndex(index int) int {
	return index + r.LowerBound
}

type LookupTable struct {
	Destinations []Range
	Name         string
	Sources      []Range
}

func NewLookupTable(name string, input []string) LookupTable {
	destinations, sources := parseRanges(input)

	return LookupTable{
		Destinations: destinations,
		Name:         name,
		Sources:      sources,
	}
}

func (l LookupTable) SourceToDestination(source int) int {
	for index, sourceTable := range l.Sources {
		if sourceTable.In(source) {
			return l.Destinations[index].FromIndex(sourceTable.ToIndex(source))
		}
	}

	return source
}

// parseRange returns the destinations and sources
func parseRanges(input []string) ([]Range, []Range) {
	var destinations, sources []Range

	for _, line := range input {
		values := strings.Split(line, " ")
		if len(values) != 3 {
			panic(fmt.Sprintf("invalid range input: %v", line))
		}

		// First value is the destination
		destination, err := strconv.Atoi(values[0])
		if err != nil {
			panic(fmt.Sprintf("invalid destination[%v]: %v", destination, line))
		}

		// Second value is the source
		source, err := strconv.Atoi(values[1])
		if err != nil {
			panic(fmt.Sprintf("invalid source[%v]: %v", source, line))
		}

		// Third value is the range size
		rangeSize, err := strconv.Atoi(values[2])
		if err != nil {
			panic(fmt.Sprintf("invalid range[%v]: %v", rangeSize, line))
		}

		destinations = append(destinations, NewRange(destination, rangeSize))
		sources = append(sources, NewRange(source, rangeSize))
	}

	return destinations, sources
}

//func D5P1(input []string) int {
//	// Create lookup tables
//	tables := []LookupTable{}
//
//	var tableName string
//	var tableInput []string
//	for _, line := range input[2:] {
//		// Ignore newlines
//		if len(line) == 0 {
//			continue
//		}
//
//		// New table
//		if strings.HasSuffix(line, " map:") {
//			if tableName != "" {
//				tables = append(tables, NewLookupTable(tableName, tableInput))
//			}
//
//			tableName = line[:len(line)-5]
//			tableInput = []string{}
//		} else {
//			tableInput = append(tableInput, line)
//		}
//	}
//
//	// Add the last table
//	tables = append(tables, NewLookupTable(tableName, tableInput))
//
//	lowestDistance := -1
//	for _, seedInput := range strings.Split(input[0][7:], " ") {
//		seed, err := strconv.Atoi(seedInput)
//		if err != nil {
//			panic(fmt.Errorf("invalid seed %v: %w", seedInput, err))
//		}
//
//		source := seed
//		for _, table := range tables {
//			source = table.SourceToDestination(source)
//		}
//
//		if lowestDistance == -1 || source < lowestDistance {
//			lowestDistance = source
//		}
//	}
//
//	return lowestDistance
//}

// parseLine returns three integers; the destination, source, and range size
func parseLine(in string) (int, int, int) {
	values := strings.Split(in, " ")
	if len(values) != 3 {
		panic(fmt.Sprintf("invalid range input: %v", in))
	}

	// First value is the destination
	destination, err := strconv.Atoi(values[0])
	if err != nil {
		panic(fmt.Sprintf("invalid destination[%v]: %v", destination, in))
	}

	// Second value is the source
	source, err := strconv.Atoi(values[1])
	if err != nil {
		panic(fmt.Sprintf("invalid source[%v]: %v", source, in))
	}

	// Third value is the range size
	rangeSize, err := strconv.Atoi(values[2])
	if err != nil {
		panic(fmt.Sprintf("invalid range[%v]: %v", rangeSize, in))
	}

	return destination, source, rangeSize
}

func parseSegments(input []string) Segment {
	start := Segment{LowerBound: 0}

	for _, line := range input {
		// Skip newlines and lines that contain mapping names
		if len(line) == 0 || strings.HasSuffix(line, " map:") {
			continue
		}

		destination, source, rangeSize := parseRange(line)

		start.Insert(NewSegment(source, destination, rangeSize))
	}

	return start
}

func last(segments []Segment) Segment {
	if len(segments) == 0 {
		return Segment{}
	}

	return segments[len(segments)-1]
}

type Node struct {
	Next  *Node
	Prev  *Node
	Value Segment
}

func NewNode(value Segment) *Node {
	return &Node{
		Value: value,
	}
}

type LinkedList struct {
	First *Node
	Last  *Node
	Count int
}

func (l *LinkedList) Append(node *Node) {
	if l.Count == 0 {
		l.First = node
		l.Last = node
		l.Count++

		return
	}

	l.Last.Next = node
	node.Prev = l.Last

	l.Last = node
}

func (l *LinkedList) Empty() bool {
	return l.Count == 0
}

func (l *LinkedList) InsertBefore(loc *Node, newNode *Node) {
	if loc.Prev != nil {
		loc.Prev.Next = newNode
	}

	if loc == l.First {
		l.First = newNode
	}

	loc.Prev = newNode
	newNode.Next = loc
	newNode.Prev = loc.Prev
	l.Count++
}

func (l *LinkedList) Remove(node *Node) {
	if l.Count <= 1 {
		l.First = nil
		l.Last = nil
		l.Count = 0

		return
	}

	if node == l.First {
		l.First = l.First.Next
		return
	}

	if node == l.Last {
		l.Last = l.Last.Prev
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func toSegment2(a any) Segment {
	return a.(Segment)
}

func toString(elem *list.Element) string {
	segment := toSegment2(elem.Value)
	out := fmt.Sprintf("[%2d - %2d | %3d]", segment.Lower, segment.Upper, segment.Translation)

	if elem.Next() == nil {
		return out
	}

	return out + "\n" + toString(elem.Next())
}

// parseSegments flattens the various maps into a single map for a O(1) lookup time complexity.
// To do this, we create a custom struct (Segment) to hold our upper and lower bounds and our value translation.
// The process to flatten the maps is:
//  1. Parse the input into a Segment. The translation is the difference between the source and destination
//  2. Get intersecting Segments
//  3. For each intersection with the new segment
//  3. Break up the new segment to match each intersecting segment
//  4. For each broken up segment, compute a new lower bound, upper bound, and translation
//  5. Add the segment to the flattened map.
func parseSegements(in []string) {
	lookupTable := list.List{}

	for _, line := range in {
		// Skip newlines and lines that contain mapping names
		// We discard map information because we are guaranteed no intersections inter-map
		if len(line) == 0 || strings.HasSuffix(line, " map:") {
			if strings.HasSuffix(line, " map:") {
				fmt.Println(line)
			}

			// TODO: Remove debug
			if len(line) == 0 {
				fmt.Println(toString(lookupTable.Front()) + "\n")
			}
			continue
		}

		// 1. Parse the input into a Segment. The translation is the difference between the source and destination
		destination, source, rangeSize := parseLine(line)
		newSegment := Segment{
			Lower: source,
			Upper: source+rangeSize-1,
			Translation: destination-source,
		}

		// 2. Get intersecting Segments
		var intersections []Segment
		for segment := lookupTable.Front(); segment != nil; segment = segment.Next() {
			if 
		}

		if segments.Len() == 0 {
			segments.PushBack(newSegment)
			continue
		}

		for currNode := segments.Front(); currNode != nil; currNode = currNode.Next() {
			segment := toSegment2(currNode.Value)

			if !segment.Intersects(newSegment) {
				if segment.Lower < newSegment.Lower {
					continue
				}

				// Insert and exit
				segments.InsertBefore(newSegment, currNode)
				break
			}

			// For each node that it intersects, apply that node to this node

			// Don't include last element because it will be added later
			flattenedSegments := segment.Flatten(newSegment)
			for _, part := range flattenedSegments[:len(flattenedSegments)-1] {
				segments.InsertBefore(part.Copy(), currNode)
			}

			newSegment = last(flattenedSegments)
			currNode = currNode.Prev()
			segments.Remove(currNode.Next())
		}
	}

}

func Part1() int {
	in := input.ReadAll(InputPath)

	var segments list.List

	for _, line := range in[2:] {
		// Skip newlines and lines that contain mapping names
		if len(line) == 0 || strings.HasSuffix(line, " map:") {
			if strings.HasSuffix(line, " map:") {
				fmt.Println(line)
			}
			if len(line) == 0 {
				fmt.Println(toString(segments.Front()) + "\n")
			}
			continue
		}

		destination, source, rangeSize := parseRange(line)
		newSegment := NewSegment2(source, source+rangeSize-1, destination-source)

		if segments.Len() == 0 {
			segments.PushBack(newSegment)
			continue
		}

		for currNode := segments.Front(); currNode != nil; currNode = currNode.Next() {
			segment := toSegment2(currNode.Value)

			if !segment.Intersects(newSegment) {
				if segment.Lower < newSegment.Lower {
					continue
				}

				// Insert and exit
				segments.InsertBefore(newSegment, currNode)
				break
			}

			// For each node that it intersects, apply that node to this node

			// Don't include last element because it will be added later
			flattenedSegments := segment.Flatten(newSegment)
			for _, part := range flattenedSegments[:len(flattenedSegments)-1] {
				segments.InsertBefore(part.Copy(), currNode)
			}

			newSegment = last(flattenedSegments)
			currNode = currNode.Prev()
			segments.Remove(currNode.Next())
		}
	}

	fmt.Println(toString(segments.Front()))

	return segments.Len()
}

func Part2() int {
	in := input.ReadAll(InputPath)

	return len(in)
}

//
//func D5P2(input []string) int {
//	// Create lookup tables
//	tables := []LookupTable{}
//
//	var tableName string
//	var tableInput []string
//	for _, line := range input[2:] {
//		// Ignore newlines
//		if len(line) == 0 {
//			continue
//		}
//
//		// New table
//		if strings.HasSuffix(line, " map:") {
//			if tableName != "" {
//				tables = append(tables, NewLookupTable(tableName, tableInput))
//			}
//
//			tableName = line[:len(line)-5]
//			tableInput = []string{}
//		} else {
//			tableInput = append(tableInput, line)
//		}
//	}
//
//	// Add the last table
//	tables = append(tables, NewLookupTable(tableName, tableInput))
//
//	// Seeds come in pair of seed start index and seed range
//	lowestDistance := -1
//	seeds := strings.Split(input[0][7:], " ")
//	for i := 0; i < len(seeds); i += 2 {
//		seedStart, err := strconv.Atoi(seeds[i])
//		if err != nil {
//			panic(fmt.Errorf("invalid seed start %v: %w", seeds[i], err))
//		}
//
//		seedRange, err := strconv.Atoi(seeds[i+1])
//		if err != nil {
//			panic(fmt.Errorf("invalid seed range  %v: %w", seeds[i+1], err))
//		}
//
//		for seed := seedStart; seed <= seedStart+seedRange; seed++ {
//			source := seed
//			for _, table := range tables {
//				source = table.SourceToDestination(source)
//			}
//
//			if lowestDistance == -1 || source < lowestDistance {
//				lowestDistance = source
//			}
//		}
//	}
//
//	return lowestDistance
//}
