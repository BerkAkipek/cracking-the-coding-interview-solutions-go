package intersecton

import "math"

/*
Intersection: Given two straight line segments (represented as a start point and an end point),
compute the point of intersection, if any.

The orientations of the pairs (p1, p2, p3) and (p1, p2, p4) are different,
and

The orientations of the pairs (p3, p4, p1) and (p3, p4, p2) are different.
*/

func PointIntersection(a, b [][]float64) []float64 {
	t, u := validation(
		a[0][0], a[0][1],
		a[1][0], a[1][1],
		b[0][0], b[0][1],
		b[1][0], b[1][1],
	)

	if t == -1.0 && u == -1.0 {
		return nil
	}

	xi := a[0][0] + t*(a[1][0]-a[0][0])
	yi := a[0][1] + t*(a[1][1]-a[0][1])
	return []float64{xi, yi}
}

func validation(
	x1, y1, x2, y2,
	x3, y3, x4, y4 float64,
) (float64, float64) {

	denoT := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	const eps = 1e-9
	if math.Abs(denoT) < eps {
		return -1.0, -1.0
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / denoT
	u := ((x1-x3)*(y1-y2) - (y1-y3)*(x1-x2)) / denoT

	if 0 <= t && t <= 1 && 0 <= u && u <= 1 {
		return t, u
	}
	return -1.0, -1.0
}
