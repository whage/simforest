package simforest

import (
	"math/rand"
)

const (
	foxCount = 20
	bunnyCount = 80
)

type Environment struct {
	width, height int
}

func (e Environment) Height() int {
	return e.height
}

func (e Environment) Width() int {
	return e.width
}

func CreateEnvironment(width, height int) *Environment {
	return &Environment{
		width,
		height,
	}
}

type Mover interface {
	Pos() Position
	Move()
}

type Mater interface {
	TryToMate(other Animal) Creature
}

type Ager interface {
	IncreaseAge()
}

type Creature interface {
	Mover
	Mater
	Ager
}

type Gender int

const (
	Male = iota
	Female
)

func CreateBunny(p Position, e *Environment) *Bunny {
	return &Bunny{
		Animal{
			e,
			Gender(rand.Intn(2)),
			p,
			0,
		},
	}
}

func CreateFox(p Position, e *Environment) *Fox {
	return &Fox{
		Animal{
			e,
			Gender(rand.Intn(2)),
			p,
			0,
		},
	}
}
/*
func (p []Creature) FilterBunnies() []Bunny {
	var results []Bunny
	for i, _ := range p {
		v, ok := p[i].(*Bunny)
		if ok {
			results = append(results, *v)
		}
	}
	return results
}
*/
func InitPopulation(e *Environment) []Creature {
	population := make([]Creature, 0)

	for i := 0; i < foxCount; i++ {
		population = append(population, CreateFox(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	for i := 0; i < bunnyCount; i++ {
		population = append(population, CreateBunny(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	return population
}

func Step(population []Creature) []Creature {
	var newPopulation []Creature

	for i, _ := range population {
		// pass rest of the population to Move() method
		//population[i].Move(append(population[:i], population[i+1:]...))
		population[i].Move()
		//population[i].TryToMate()
		population[i].IncreaseAge()

		// Carry current object over to population of next round
		newPopulation = append(newPopulation, population[i])
	}

	return newPopulation
}
