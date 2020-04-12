package simforest

import (
	"math/rand"
	_ "fmt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	foxCount = 15
	bunnyCount = 25
)

type Environment struct {
	width int
	height int
	tickCount int
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
		0,
	}
}

type Creature interface {
	Pos() Position
	Env() *Environment
	Move([]Creature)
	Age() int
	IncreaseAge()
	Gender() Gender
	Mate(Creature) Creature
	CanStartMating() bool
	Render() string
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

func Tick(population []Creature, env *Environment) []Creature {
	var newPopulation []Creature

	for i, _ := range population {
		population[i].Move(population)
		// add potential newborn to population
		if population[i].CanStartMating() {
			offSpring := TryToMate(population[i], population)
			if offSpring != nil {
				newPopulation = append(newPopulation, offSpring)
			}
		}
		population[i].IncreaseAge()
		// Carry current object over to population of next round
		newPopulation = append(newPopulation, population[i])
	}

	env.tickCount += 1

	return newPopulation
}
