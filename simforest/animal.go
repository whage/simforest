package simforest

import (
	"math/rand"
	"strings"
	"fmt"
)

type Animal struct {
	environment *Environment
	gender Gender
	pos Position
	age int
	tickLastMated int
}

func (a *Animal) Move(population []Creature) {
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

func (a Animal) Age() int {
	return a.age
}

func (a Animal) Pos() Position {
	return a.pos
}

func (a Animal) Env() *Environment {
	return a.environment
}

func (a *Animal) IncreaseAge() {
	a.age += 1
}

func (a Animal) CanStartMating(ticksBetweenMating int) bool {
	if a.Env().tickCount > a.tickLastMated + ticksBetweenMating {
		return true
	}
	return false
}

func (a Animal) Render(marker string) string {
	var colorCode string

	if a.Gender() == Female {
		colorCode = ErrorColor
	} else {
		colorCode = InfoColor
	}

	if a.Age() < 30 {
		return fmt.Sprintf(colorCode, marker)
	} else {
		return fmt.Sprintf(colorCode, strings.ToUpper(marker))
	}
}
