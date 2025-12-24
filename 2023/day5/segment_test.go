package day5_test

import (
	"github.com/spencerreeves/advent_of_code/2023/day5"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSegment2(t *testing.T) {
	assert.Equal(t, day5.Segment{}, day5.NewSegment2(0, 0, 0))
}

func TestCopy(t *testing.T) {
	assert.Equal(t, day5.Segment{}.Copy(), day5.NewSegment2(0, 0, 0))
}

func TestSort(t *testing.T) {

	type testCase struct {
		name    string
		segment day5.Segment
		target  day5.Segment
		out     []day5.Segment
	}

	testCases := []testCase{
		{
			name:    "Pre sorted",
			segment: day5.NewSegment2(0, 1, 0),
			target:  day5.NewSegment2(2, 2, 0),
			out: []day5.Segment{
				{0, 1, 0},
				{2, 2, 0},
			},
		},
		{
			name:    "Out of order",
			segment: day5.NewSegment2(2, 2, 0),
			target:  day5.NewSegment2(0, 1, 0),
			out: []day5.Segment{
				{0, 1, 0},
				{2, 2, 0},
			},
		},
		{
			name:    "Single element segments",
			segment: day5.NewSegment2(0, 0, 0),
			target:  day5.NewSegment2(1, 1, 0),
			out: []day5.Segment{
				{0, 0, 0},
				{1, 1, 0},
			},
		},
		{
			name:    "Intersects",
			segment: day5.NewSegment2(0, 2, 0),
			target:  day5.NewSegment2(1, 2, 0),
			out:     []day5.Segment{},
		},
		{
			name:    "Copy",
			segment: day5.NewSegment2(0, 1, 0),
			target:  day5.NewSegment2(0, 1, 0),
			out:     []day5.Segment{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			out := test.segment.Sort(test.target)
			assert.Equal(t, test.out, out)
		})
	}
}

func TestIntersects(t *testing.T) {

	type testCase struct {
		name    string
		segment day5.Segment
		target  day5.Segment
		out     bool
	}

	testCases := []testCase{
		{
			name:    "Overlap",
			segment: day5.NewSegment2(1, 1, 0),
			target:  day5.NewSegment2(0, 2, 0),
			out:     true,
		},
		{
			name:    "Starts and ends in segment",
			segment: day5.NewSegment2(1, 2, 0),
			target:  day5.NewSegment2(1, 1, 0),
			out:     true,
		},
		{
			name:    "Starts out and ends in segment",
			segment: day5.NewSegment2(1, 2, 0),
			target:  day5.NewSegment2(0, 1, 0),
			out:     true,
		},
		{
			name:    "Starts in and ends out segment",
			segment: day5.NewSegment2(1, 2, 0),
			target:  day5.NewSegment2(1, 3, 0),
			out:     true,
		},
		{
			name:    "Same segment",
			segment: day5.NewSegment2(1, 2, 0),
			target:  day5.NewSegment2(1, 2, 0),
			out:     true,
		},
		{
			name:    "Same point",
			segment: day5.NewSegment2(1, 1, 0),
			target:  day5.NewSegment2(1, 1, 0),
			out:     true,
		},
		{
			name:    "Before segment, no overlap",
			segment: day5.NewSegment2(2, 3, 0),
			target:  day5.NewSegment2(0, 1, 0),
			out:     false,
		},
		{
			name:    "After segment, no overlap",
			segment: day5.NewSegment2(0, 1, 0),
			target:  day5.NewSegment2(2, 3, 0),
			out:     false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			out := test.segment.Intersects(test.target)
			assert.Equal(t, test.out, out)
		})
	}
}

func TestFlatten(t *testing.T) {

	type testCase struct {
		name    string
		segment day5.Segment
		target  day5.Segment
		out     []day5.Segment
	}

	testCases := []testCase{
		{
			name:    "No intersection",
			segment: day5.NewSegment2(0, 1, 0),
			target:  day5.NewSegment2(2, 3, 0),
			out: []day5.Segment{
				{0, 1, 0},
				{2, 3, 0},
			},
		},
		{
			name:    "No intersection - reverse",
			segment: day5.NewSegment2(2, 3, 0),
			target:  day5.NewSegment2(0, 1, 0),
			out: []day5.Segment{
				{0, 1, 0},
				{2, 3, 0},
			},
		},
		{
			name:    "Full intersection",
			segment: day5.NewSegment2(1, 2, 1),
			target:  day5.NewSegment2(0, 3, 2),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 2, 3},
				{3, 3, 2},
			},
		},
		{
			name:    "Full intersection - reverse",
			segment: day5.NewSegment2(0, 3, 2),
			target:  day5.NewSegment2(1, 2, 1),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 2, 3},
				{3, 3, 2},
			},
		},
		{
			name:    "Lower intersection",
			segment: day5.NewSegment2(1, 3, 1),
			target:  day5.NewSegment2(0, 2, 2),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 2, 3},
				{3, 3, 1},
			},
		},
		{
			name:    "Lower intersection - reverse",
			segment: day5.NewSegment2(0, 2, 2),
			target:  day5.NewSegment2(1, 3, 1),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 2, 3},
				{3, 3, 1},
			},
		},
		{
			name:    "Lower intersection - bound",
			segment: day5.NewSegment2(0, 1, 2),
			target:  day5.NewSegment2(1, 3, 1),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 1, 3},
				{2, 3, 1},
			},
		},
		{
			name:    "Lower intersection - bound and reverse",
			segment: day5.NewSegment2(1, 3, 1),
			target:  day5.NewSegment2(0, 1, 2),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 1, 3},
				{2, 3, 1},
			},
		},
		{
			name:    "Upper intersection",
			segment: day5.NewSegment2(0, 2, 1),
			target:  day5.NewSegment2(1, 3, 2),
			out: []day5.Segment{
				{0, 0, 1},
				{1, 2, 3},
				{3, 3, 2},
			},
		},
		{
			name:    "Upper intersection - reverse",
			segment: day5.NewSegment2(1, 3, 2),
			target:  day5.NewSegment2(0, 2, 1),
			out: []day5.Segment{
				{0, 0, 1},
				{1, 2, 3},
				{3, 3, 2},
			},
		},
		{
			name:    "Upper intersection - bound",
			segment: day5.NewSegment2(0, 1, 2),
			target:  day5.NewSegment2(1, 2, 1),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 1, 3},
				{2, 2, 1},
			},
		},
		{
			name:    "Lower intersection - bound and reverse",
			segment: day5.NewSegment2(1, 2, 1),
			target:  day5.NewSegment2(0, 1, 2),
			out: []day5.Segment{
				{0, 0, 2},
				{1, 1, 3},
				{2, 2, 1},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			out := test.segment.Flatten(test.target)
			assert.Equal(t, test.out, out)
		})
	}
}
