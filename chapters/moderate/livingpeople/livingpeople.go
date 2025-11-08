package livingpeople

/*
Living People: Given a list of people with their birth and death years, implement a method to
compute the year with the most number of people alive. You may assume that all people were born
between 1900 and 2000 (inclusive). If a person was alive during any portion of that year, they should
be included in that year's count. For example. Person (birth = 1908, death = 1909) is included in the
counts for both 1908 and 1909
*/

type People struct {
	Birth int
	Death int
}

func MostPeopleAlive(people []*People, interval [2]int) int {
	if len(people) == 0 {
		return interval[0]
	}

	if interval[0] > interval[1] {
		panic("invalid interval")
	}

	delta := make(map[int]int)

	for _, p := range people {
		delta[p.Birth]++
		delta[p.Death+1]--
	}

	population := 0
	maxAlive := 0
	maxYear := 0
	for year := interval[0]; year <= interval[1]; year++ {
		population += delta[year]
		if population > maxAlive {
			maxAlive = population
			maxYear = year
		}
	}

	return maxYear
}
