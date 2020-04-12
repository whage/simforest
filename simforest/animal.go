package simforest

import (
	"math/rand"
)

type Animal struct {
	environment *Environment
	gender Gender
	pos Position
	age int
}

func (a *Animal) Move() {
	randomStep(&a.pos, a.environment)
}

func (a Animal) Gender() Gender {
	return a.gender
}

func (a Animal) Age() int {
	return a.age
}

func (a Animal) Pos() Position {
	return a.pos
}

func (a Animal) Env() *Environment {
	return a.environment
}

func (a Animal) IncreaseAge() {
	a.age += 1
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
