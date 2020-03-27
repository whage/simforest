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

type Breeder interface {
	TryToMate(others []Breeder) Breeder
}

type Ager interface {
	IncreaseAge()
}

type Creature interface {
	Mover
	Breeder
	Ager
}

type Animal struct {
	environment *Environment
	gender Gender
	pos Position
	age int
}

func (a Animal) Gender() Gender {
	return a.gender
}

func (a Animal) Age() int {
	return a.age
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

func InitPopulation(e *Environment) Population {
	population := make(Population, 0)

	for i := 0; i < foxCount; i++ {
		population = append(population, CreateFox(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	for i := 0; i < bunnyCount; i++ {
		population = append(population, CreateBunny(Position{rand.Intn(e.width), rand.Intn(e.height)}, e))
	}

	return population
}

func Step(population Population) Population {
	var newPopulation Population

	for i, _ := range population {
		population[i].Move()
		population[i].IncreaseAge()

		switch current := population[i].(type) {
		//case *Fox:
			// TODO: continue (as with *Bunny below)
		case *Bunny:
			bunnies := population.FilterBunnies()

			// Convert Bunnies to Breeders to satisfy TryToMate() signature
			breeders := make([]Breeder, len(bunnies), len(bunnies))
			for i := range bunnies {
				breeders[i] = &bunnies[i]
			}
			if newBreeder := current.TryToMate(breeders); newBreeder != nil {
				// Can this type assertion fail at runtime?
				// If not, how does the type checker know?
				newPopulation = append(newPopulation, newBreeder.(*Bunny))
			}
		}

		// Carry current object over to population of next round
		newPopulation = append(newPopulation, population[i])
	}

	return newPopulation
}

func randomStep(p *Position, e *Environment) {
	steps := []int{-1, 0, 1}
	direction := Position{
		steps[rand.Intn(len(steps))],
		steps[rand.Intn(len(steps))],
	}

	if newPosition := p.Add(direction); newPosition.IsWithinEnvironment(e) {
		*p = p.Add(direction)
	}
}
