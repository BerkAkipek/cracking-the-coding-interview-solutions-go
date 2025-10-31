package blueeyedisland

import (
	"fmt"
	"math/rand"
)

/*
Blue-Eyed Island: A bunch of people are living on an island, when a visitor comes with a strange
order: all blue-eyed people must leave the island as soon as possible. There will be a flight out at
8:00 pm every evening. Each person can see everyone else's eye color, but they do not know their
own (nor is anyone allowed to tell them). Additionally, they do not know how many people have
blue eyes, although they do know that at least one person does. How many days will it take the
blue-eyed people to leave?
*/

type Person struct {
	id           int
	eyeColor     string
	visibleBlues int
	isBlue       bool
	departed     bool
}

type Island struct {
	people []*Person
}

func CreateIsland(population int) *Island {
	choices := []string{"blue", "brown"}
	people := []*Person{}
	for i := 0; i < population; i++ {
		color := choices[rand.Intn(len(choices))]
		p := &Person{id: i, eyeColor: color}
		if color == "blue" {
			p.isBlue = true
		}
		people = append(people, p)
	}
	return &Island{people: people}
}

func Game() {
	island := CreateIsland(100)

	numBlue := 0
	for _, p := range island.people {
		if p.eyeColor == "blue" {
			numBlue++
		}
	}
	for _, p := range island.people {
		if p.eyeColor == "blue" {
			p.visibleBlues = numBlue - 1
		} else {
			p.visibleBlues = numBlue
		}
	}

	day := 0
	for {
		day++
		leftToday := 0

		for _, p := range island.people {
			if p.departed {
				continue
			}
			if p.isBlue && day == p.visibleBlues+1 {
				p.departed = true
				leftToday++
			}
		}

		fmt.Printf("Day %d: %d blue-eyed people left.\n", day, leftToday)

		if leftToday == numBlue {
			fmt.Printf("All %d blue-eyed people have left on day %d.\n", numBlue, day)
			break
		}
	}
}
