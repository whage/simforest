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

func (a *Animal) TryToMate(other Animal) Creature {
	if (a.gender != other.gender) && other.age > 30 && a.age > 30 {
		//return CreateBunny(Position{other.pos.X, other.pos.Y}, other.environment)
	}
	return nil
}

func (a *Animal) Pos() Position {
	return a.pos
}

func (a *Animal) IncreaseAge() {
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
