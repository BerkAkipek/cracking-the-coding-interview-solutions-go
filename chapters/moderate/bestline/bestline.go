package bestline

import "math"

/*
Best Line: Given a two-dimensional graph with points on it, find a line which passes the most number of points.
*/
type Point struct {
	X, Y float64
}

func slope(a, b Point) float64 {
	if b.X == a.X {
		return math.Inf(1)
	}
	return (b.Y - a.Y) / (b.X - a.X)
}

type Line struct {
	M, B float64
}

func BestLine(graph []Point) (Line, int) {
	var bestLine Line
	bestCount := 0

	for i := range graph {
		slopes := make(map[Line]int)
		for j := i + 1; j < len(graph); j++ {
			m := slope(graph[i], graph[j])
			var b float64
			if math.IsInf(m, 1) {
				b = graph[i].X
			} else {
				b = graph[i].Y - m*graph[i].X
			}
			m = math.Round(m*1e9) / 1e9
			b = math.Round(b*1e9) / 1e9

			line := Line{M: m, B: b}
			slopes[line]++

			if slopes[line] > bestCount {
				bestCount = slopes[line]
				bestLine = line
			}
		}
	}

	return bestLine, bestCount + 1
}
