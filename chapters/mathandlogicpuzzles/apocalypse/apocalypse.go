package apocalypse

import "math/rand"

/*
The Apocalypse: In the new post-apocalyptic world, the world queen is desperately concerned
about the birth rate. Therefore, she decrees that all families should ensure that they have one girl or
else they face massive fines. If all families abide by this policy that is, they have continue to have
children until they have one girl, at which point they immediately stop-what will the gender ratio
of the new generation be? (Assume that the odds of someone having a boy or a girl on any given
pregnancy is equal.) Solve this out logically and then write a computer simulation of it
*/

func PopulationExperiment(numFamilies int) float64 {
	possible := []bool{true, false}
	girlCount := 0
	boyCount := 0
	for range numFamilies {
		for {
			res := possible[rand.Intn(len(possible))]
			if res {
				girlCount++
				break
			}
			boyCount++
		}
	}

	return (float64(girlCount) / (float64(girlCount) + float64(boyCount)))
}
