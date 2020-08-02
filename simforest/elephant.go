package simforest

import (
	"math/rand"
)

const ElephantTicksBetweenMating = 40

type Elephant struct {
	Animal
}

func (e *Elephant) Mate(other Entity, population []Entity) []Entity {
	return []Entity{}
}

func NewElephant(p Position, e *Environment) *Elephant {
	return &Elephant{
		Animal{
			Position: p,
			Environment: e,
			gender: Gender(rand.Intn(2)),
			age: 0,
			tickLastMated: e.tickCount,
			isAlive: true,
			tickOfBirth: e.tickCount,
			direction: nil,
		},
	}
}

func (f Elephant) Render() Marker {
	var color string

	if f.Gender() == Female {
		color = MaleElephantColor
	} else {
		color = FemaleElephantColor
	}

	return Marker{
		color,
		"e",
	}
}

func (f Elephant) IsAdult() bool {
	return f.Age() > 30
}

func (f *Elephant) Act(population []Entity) []Entity {
	f.Move(population)

	for _, c := range population {
		t, isTree := c.(*Tree)
		if isTree && t.IsNearby(f.Position) {
			t.Die()
			break
		}
	}

	return make([]Entity, 0)
}

func (f *Elephant) IsAtEndOfLife() bool {
	return f.Age() > 200
}
