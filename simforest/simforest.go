package simforest

import (
	"math/rand"
	_ "fmt"
)

const (
	foxCount = 30
	bunnyCount = 80
	carrotCount = 30
	treeCount = 250
)

const (
	LightBrown = "\033[38;5;101m%s\033[0m"
	DarkGreen = "\033[38;5;22m%s\033[0m"
	LightBlue  = "\033[1;34m%s\033[0m"
	Salmon = "\033[38;5;216m%s\033[0m"
	LightPink = "\033[38;5;182m%s\033[0m"
	Teal = "\033[1;36m%s\033[0m"
	Yellow = "\033[1;33m%s\033[0m"
	Red = "\033[1;31m%s\033[0m"
	Cyan = "\033[0;36m%s\033[0m"
	Orange = "\033[38;5;214m%s\033[0m"
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

type Marker struct {
	Color string
	Character string
}

type Entity interface {
	Pos() Position
	Act([]Entity) []Entity
	Move([]Entity)
	Gender() Gender
	Mate(Entity, []Entity) []Entity
	Render() Marker
	IsAlive() bool
	IsAdult() bool
}

type Gender int

const (
	Male = iota
	Female
)

func InitPopulation(e *Environment) []Entity {
	population := make([]Entity, 0)

	for i := 0; i < foxCount; i++ {
		population = append(population, NewFox(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	for i := 0; i < bunnyCount; i++ {
		population = append(population, NewBunny(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	for i := 0; i < carrotCount; i++ {
		population = append(population, &Carrot{Position{rand.Intn(e.width), rand.Intn(e.height)}})
	}

	for i := 0; i < treeCount; i++ {
		population = append(population, &Tree{Position{rand.Intn(e.width), rand.Intn(e.height)}})
	}	

	return population
}

func TryToMate(c Entity, others []Entity) []Entity {
	for _, o := range others {
		if c == o { continue } // skip self
		if o.IsAlive() && c.Gender() != o.Gender() && c.Pos().IsNearby(o.Pos()) {
			return c.Mate(o, others)
		}
	}
	return nil
}

func Tick(population []Entity, env *Environment) []Entity {
	var populationForNextTick []Entity

	for i, _ := range population {
		if !population[i].IsAlive() {
			continue
		}
		newEntities := population[i].Act(population)

		// add any new objects to population
		populationForNextTick = append(populationForNextTick, newEntities...)

		// Carry current object over to population of next round
		populationForNextTick = append(populationForNextTick, population[i])
	}

	env.tickCount += 1

	return filterOutDead(populationForNextTick)
}

func filterOutDead(population []Entity) []Entity {
	results := make([]Entity, 0, len(population))
	for _, c := range population {
		if c.IsAlive() {
			results = append(results, c)
		}
	}
	return results
}
