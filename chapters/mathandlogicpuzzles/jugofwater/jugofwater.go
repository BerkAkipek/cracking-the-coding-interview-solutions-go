package jugofwater

/*
Jugs of Water: You have a five-quart jug, a three-quart jug, and an unlimited supply of water (but
no measuring cups). How would you come up with exactly four quarts of water? Note that the jugs
are oddly shaped, such that filling up exactly "half" of the jug would be impossible
*/

import (
	"fmt"
)

type S struct{ a, b int }

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func bfs(A, B, T int) {
	if T > A && T > B || T%gcd(A, B) != 0 {
		fmt.Println("No solution")
		return
	}
	q := []S{{0, 0}}
	v := map[S]bool{}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		if s.a == T || s.b == T {
			fmt.Println("Goal:", s)
			return
		}
		if v[s] {
			continue
		}
		v[s] = true
		a, b := s.a, s.b
		next := []S{
			{A, b}, {a, B}, {0, b}, {a, 0},
			{a - min(a, B-b), b + min(a, B-b)},
			{a + min(b, A-a), b - min(b, A-a)},
		}
		q = append(q, next...)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
