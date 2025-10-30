package basketball

import "math"

/*
Basketball: You have a basketball hoop and someone says that you can play one of two games.
Game 1: You get one shot to make the hoop.
Game 2: You get three shots and you have to make two of three shots.
If p is the probability of making a particular shot, for which values of p should you pick one game or the other?
*/

func factorial(n int) int {
	result := 1
	for i := 1; i < n+1; i++ {
		result *= i
	}
	return result
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func WinProbability(n, k int, p float64) float64 {
	result := 0.0
	for i := k; k <= n; i++ {
		result += float64(combination(n, k)) * math.Pow(p, float64(k)) * math.Pow(1-p, float64((n-k)))
	}
	return result
}
