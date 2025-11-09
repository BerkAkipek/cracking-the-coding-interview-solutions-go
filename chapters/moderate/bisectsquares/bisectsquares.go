package bisectsquares

import "math"

/*
Bisect Squares: Given two squares on a two-dimensional plane, find a line that would cut these two
squares in half. Assume that the top and the bottom sides of the square run parallel to the x-axis.
*/

type Point struct {
	X float64
	Y float64
}

type Square struct {
	Center Point
	Side   float64
}

type Segment struct {
	A Point
	B Point
}

func BisectSquares(a, b Square) Segment {
	if math.Abs(a.Center.X-b.Center.X) < 1e-9 {
		return Segment{
			A: Point{a.Center.X, a.Center.Y - a.Side/2},
			B: Point{b.Center.X, b.Center.Y + b.Side/2},
		}
	}

	m := (b.Center.Y - a.Center.Y) / (b.Center.X - a.Center.X)
	c := a.Center.Y - m*a.Center.X

	dir := Point{b.Center.X - a.Center.X, b.Center.Y - a.Center.Y}

	choose := func(center Point, points []Point, dir Point) Point {
		var best Point
		for _, p := range points {
			vx := p.X - center.X
			vy := p.Y - center.Y
			if (vx*dir.X >= 0) && (vy*dir.Y >= 0) {
				best = p
				break
			}
		}
		return best
	}

	pointsA := intersectionsFor(a, m, c)
	pointsB := intersectionsFor(b, m, c)
	A := choose(a.Center, pointsA, dir)
	B := choose(b.Center, pointsB, Point{-dir.X, -dir.Y})

	return Segment{
		A: A,
		B: B,
	}
}

func intersectionsFor(s Square, m, c float64) []Point {
	points := []Point{}

	left := s.Center.X - s.Side/2
	right := s.Center.X + s.Side/2
	bottom := s.Center.Y - s.Side/2
	top := s.Center.Y + s.Side/2

	for _, x := range []float64{left, right} {
		y := m*x + c
		if y >= bottom && y <= top {
			points = append(points, Point{x, y})
		}
	}

	if m != 0 {
		for _, y := range []float64{bottom, top} {
			x := (y - c) / m
			if x >= left && x <= right {
				points = append(points, Point{x, y})
			}
		}
	}

	return points
}
