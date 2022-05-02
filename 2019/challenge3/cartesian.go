package challenge3

import (
	"errors"
	"math"
)

type Point struct{
	x int
	y int
}

type Line struct{
	p1 Point
	p2 Point
	steps int
}

func axisIntersection(start1, end1, start2, end2 int) (int, error) {
	s1 := math.Min(float64(start1), float64(end1))
	e1 := math.Max(float64(start1), float64(end1))
	s2 := math.Min(float64(start2), float64(end2))
	e2 := math.Max(float64(start2), float64(end2))

	// Lines are co-linear. Return the first point they intersect
	if s1 != e1 && s2 != e2 {
		return int(math.Max(s1, s2)), nil
	}

	// Segment 1 is a single point
	if s1 == e1 && s1 >= s2 && s1 <= e2 {
		return int(s1), nil
	}

	// Segment 2 is a single point
	if s2 == e2 && s2 >= s1 && s2 <= e1 {
		return int(s2), nil
	}

	return -1, errors.New("no intersection")
}

// If the lines intersect, returns true and the intersection point, else false and nil
func DoLinesIntersect(l1, l2 Line) (Point, int, error) {
	xAxisIntersection, errX := axisIntersection(l1.p1.x, l1.p2.x, l2.p1.x, l2.p2.x)
	yAxisIntersection, errY := axisIntersection(l1.p1.y, l1.p2.y, l2.p1.y, l2.p2.y)

	if errX != nil || errY != nil {
		return Point{0,0}, 0, errors.New("no intersection")
	}

	if xAxisIntersection == 0 && yAxisIntersection == 0 {
		return Point{0,0}, 0, errors.New("no intersection")
	}

	xSub := Abs(l1.p2.x - xAxisIntersection) + Abs(l2.p2.x - xAxisIntersection)
	ySub := Abs(l1.p2.y - yAxisIntersection) + Abs(l2.p2.y - yAxisIntersection)

	return Point{xAxisIntersection, yAxisIntersection}, l1.steps + l2.steps - xSub - ySub, nil
}