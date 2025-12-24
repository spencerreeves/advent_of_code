package day5

type Segment struct {
	Lower       int // inclusive
	Upper       int // inclusive
	Translation int
}

func NewSegment2(lower, upper, transformation int) Segment {
	return Segment{
		Lower:       lower,
		Upper:       upper,
		Translation: transformation,
	}
}

// Copy returns a deep copy of this segment
func (s Segment) Copy() Segment {
	return NewSegment2(s.Lower, s.Upper, s.Translation)
}

// Sort returns a list of the target and this segment in sorted order
func (s Segment) Sort(target Segment) []Segment {
	if s.Intersects(target) {
		return []Segment{}
	}

	if s.Lower < target.Lower {
		return []Segment{s.Copy(), target.Copy()}
	}

	return []Segment{target.Copy(), s.Copy()}
}

// Flatten takes a segment and combines it with this segment.
// The returned segments include a list of segments that do not intersect
func (s Segment) Flatten(target Segment) []Segment {
	if !s.Intersects(target) {
		return s.Sort(target)
	}

	first, second := s.Copy(), target.Copy()
	if first.Lower > second.Lower {
		first, second = second, first
	}

	// Larger than segment, like [0-3] and [1-2]
	if first.Lower < second.Lower && first.Upper > second.Upper {
		return []Segment{
			NewSegment2(first.Lower, second.Lower-1, first.Translation),
			NewSegment2(second.Lower, second.Upper, first.Translation+second.Translation),
			NewSegment2(second.Upper+1, first.Upper, first.Translation),
		}
	}

	// The same segment, like [0-1] and [0-1]
	if first.Lower == second.Lower && first.Upper == second.Upper {
		first.Translation += second.Translation
		return []Segment{first}
	}

	// Either first fully encloses second, or vice versa.
	// In this situation, there will only be two segments
	// Example [0-2] and [0-1] or [0-1] and [0-2]
	if first.Lower == second.Lower {
		segment := NewSegment2(second.Upper+1, first.Upper, first.Translation)
		if first.Upper < second.Upper {
			segment = NewSegment2(first.Upper+1, second.Upper, second.Translation)
		}

		return []Segment{
			NewSegment2(first.Lower, segment.Lower-1, first.Translation+second.Translation),
			segment,
		}
	}

	// The first segment starts before the second segment;
	// and the second segment extends past the first segment
	// Example [0-2] and [1-3] or [0-10] and [10-20]
	return []Segment{
		NewSegment2(first.Lower, second.Lower-1, first.Translation),
		NewSegment2(second.Lower, first.Upper, first.Translation+second.Translation),
		NewSegment2(first.Upper+1, second.Upper, second.Translation),
	}
}

// Intersects returns true when part of the target segment overlaps with this segment
func (s Segment) Intersects(target Segment) bool {
	return (target.Lower < s.Lower && target.Upper > s.Upper) || // Larger than segment
		(target.Lower >= s.Lower && target.Lower <= s.Upper) || // Starts in segment
		(target.Upper >= s.Lower && target.Upper <= s.Upper) // Ends in segment
}
