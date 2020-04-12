package simforest

import (
	"math/rand"
	_ "fmt"
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

type Creature interface {
	Pos() Position
	Env() *Environment
	Move()
	Age() int
	IncreaseAge()
	Gender() Gender
	Mate(Creature) Creature
}

type Gender int

const (
	Male = iota
	Female
)

func InitPopulation(e *Environment) []Creature {
	population := make([]Creature, 0)

	for i := 0; i < foxCount; i++ {
		population = append(population, NewFox(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	for i := 0; i < bunnyCount; i++ {
		population = append(population, NewBunny(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	return population
}

func TryToMate(c Creature, others []Creature) Creature {
	for _, o := range others {
		if c == o { continue } // skip self
		if c.Gender() != o.Gender() && c.Pos().IsNearby(o.Pos()) {
			return c.Mate(o)
		}
	}
	return nil
}

func Step(population []Creature) []Creature {
	var newPopulation []Creature

	for i, _ := range population {
		population[i].Move()
		// add potential newborn to population
		newBorn := TryToMate(population[i], population)
		if newBorn != nil {
			newPopulation = append(newPopulation, newBorn)
		}
		population[i].IncreaseAge()
		// Carry current object over to population of next round
		newPopulation = append(newPopulation, population[i])
	}
	//fmt.Println(len(newPopulation))
	return newPopulation
}
