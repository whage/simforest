package simforest

import (
	"math/rand"
)

type Animal struct {
	environment *Environment
	gender Gender
	pos Position
	age int
	tickLastMated int
	isAlive bool
	tickOfBirth int
}

func (a *Animal) Move(population []Entity) {
	steps := []int{-1, 0, 1}
	direction := Position{
		steps[rand.Intn(len(steps))],
		steps[rand.Intn(len(steps))],
	}

	newPosition := a.pos.Add(direction)
	canMoveThere := newPosition.IsWithinEnvironment(a.environment) && !newPosition.IsTaken(population)

	if canMoveThere {
		a.pos= newPosition
	}
}

func (a Animal) Gender() Gender {
	return a.gender
}

func (a *Animal) Age() int {
	return a.environment.tickCount - a.tickOfBirth
}

func (a Animal) Pos() Position {
	return a.pos
}

func (a Animal) Env() *Environment {
	return a.environment
}

func (a *Animal) CanStartMating(ticksBetweenMating int) bool {
	if a.Env().tickCount > a.tickLastMated + ticksBetweenMating {
		return true
	}
	return false
}

func (a *Animal) Die() {
	a.isAlive = false
}

func (a *Animal) IsAlive() bool {
	return a.isAlive
}

func (a *Animal) CommonAct(population []Entity, ticksBetweenMating int, currentEntity Entity) []Entity {
	a.Move(population)

	var newEntities []Entity

	// add potential newborn to population
	if a.CanStartMating(ticksBetweenMating) {
		offSpring := TryToMate(currentEntity, population)
		if offSpring != nil {
			newEntities = offSpring
		}
	}

	return newEntities
}
