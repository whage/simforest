package simforest

import (
	"math/rand"
	_ "fmt"
)

const (
	foxCount = 40
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
	Act([]Creature) []Creature
	Move([]Creature)
	Age() int
	Gender() Gender
	Mate(Creature) []Creature
	CanStartMating(int) bool
	Render() string
	Die()
	IsAlive() bool
	IsAdult() bool
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

func TryToMate(c Creature, others []Creature) []Creature {
	for _, o := range others {
		if c == o { continue } // skip self
		if o.IsAlive() && c.Gender() != o.Gender() && c.Pos().IsNearby(o.Pos()) {
			return c.Mate(o)
		}
	}
	return nil
}

func Tick(population []Creature, env *Environment) []Creature {
	var populationForNextTick []Creature

	for i, _ := range population {
		if !population[i].IsAlive() {
			continue
		}
		newCreatures := population[i].Act(population)

		// add any new objects to population
		populationForNextTick = append(populationForNextTick, newCreatures...)

		// Carry current object over to population of next round
		populationForNextTick = append(populationForNextTick, population[i])
	}

	env.tickCount += 1

	return filterOutDead(populationForNextTick)
}

func filterOutDead(population []Creature) []Creature {
	results := make([]Creature, 0, len(population))
	for _, c := range population {
		if c.IsAlive() {
			results = append(results, c)
		}
	}
	return results
}
