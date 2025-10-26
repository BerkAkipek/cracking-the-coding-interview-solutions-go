package animalshelter

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"slices"
)

/*
Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
that type). They cannot select which specific animal they would like. Create the data structures to
maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
and dequeueCat. You may use the built-in Linked list data structure.
*/

type Animal interface {
	MakeNoise() string
}

type Dog struct{ ID *big.Int }

func NewDog() (*Dog, error) {
	id, err := rand.Prime(rand.Reader, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot generate unique id: %w", err)
	}

	dog := &Dog{ID: id}
	return dog, nil
}

type Cat struct{ ID *big.Int }

func NewCat() (*Cat, error) {
	id, err := rand.Prime(rand.Reader, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot generate unique id: %w", err)
	}

	cat := &Cat{ID: id}
	return cat, nil
}

type DogQueue struct {
	dogs []*Dog
}

type CatQueue struct {
	cats []*Cat
}

type Shelter struct {
	dogQueue      *DogQueue
	catQueue      *CatQueue
	adoptionQueue []Animal
}

func NewShelter() *Shelter {
	return &Shelter{
		dogQueue: &DogQueue{dogs: []*Dog{}},
		catQueue: &CatQueue{cats: []*Cat{}},
	}
}

func (d *Dog) MakeNoise() string { return "Bark" }
func (c *Cat) MakeNoise() string { return "Miav" }

func (s *Shelter) Enqueue(a Animal) {
	switch x := a.(type) {
	case *Dog:
		s.dogQueue.dogs = append(s.dogQueue.dogs, x)
	case *Cat:
		s.catQueue.cats = append(s.catQueue.cats, x)
	}
	s.adoptionQueue = append(s.adoptionQueue, a)
}

func (s *Shelter) DequeueAny() (Animal, error) {
	if len(s.adoptionQueue) == 0 {
		return nil, fmt.Errorf("there are no animals in shelter")
	}

	animal := s.adoptionQueue[0]
	switch animal.(type) {
	case *Dog:
		s.dogQueue.dogs = s.dogQueue.dogs[1:]
	case *Cat:
		s.catQueue.cats = s.catQueue.cats[1:]
	}

	s.adoptionQueue = s.adoptionQueue[1:]
	return animal, nil
}

func (s *Shelter) DequeueDog() (*Dog, error) {
	if len(s.dogQueue.dogs) == 0 {
		return nil, fmt.Errorf("there are no dogs in the shelter")
	}

	firstDog := s.dogQueue.dogs[0]
	s.dogQueue.dogs = s.dogQueue.dogs[1:]
	for ind, animus := range s.adoptionQueue {
		if x, ok := animus.(*Dog); ok && x.ID.Cmp(firstDog.ID) == 0 {
			s.adoptionQueue = slices.Delete(s.adoptionQueue, ind, ind+1)
			break
		}
	}

	return firstDog, nil
}

func (s *Shelter) DequeueCat() (*Cat, error) {
	if len(s.catQueue.cats) == 0 {
		return nil, fmt.Errorf("there are no cats in the shelter")
	}

	firstCat := s.catQueue.cats[0]
	s.catQueue.cats = s.catQueue.cats[1:]
	for ind, animus := range s.adoptionQueue {
		if x, ok := animus.(*Cat); ok && x.ID.Cmp(firstCat.ID) == 0 {
			s.adoptionQueue = slices.Delete(s.adoptionQueue, ind, ind+1)
		}
	}

	return firstCat, nil
}
