package simforest

import (
	"math/rand"
)

const BunnyTicksBetweenMating = 20

type Bunny struct {
	Animal
}

func (b *Bunny) Mate(other Entity, population []Entity) []Entity {
	_, ok := other.(*Bunny)
	if ok {
		if offSpringPosition := b.FindFreeNeighborTile(population, b.Env()); offSpringPosition != nil {
			b.tickLastMated = b.Env().tickCount
			return []Entity{NewBunny(*offSpringPosition, b.Env())}
		}
	}
	return []Entity{}
}

func NewBunny(p Position, e *Environment) *Bunny {
	return &Bunny{
		Animal{
			Position: p,
			environment: e,
			gender: Gender(rand.Intn(2)),
			age: 0,
			tickLastMated: e.tickCount,
			isAlive: true,
			tickOfBirth: e.tickCount,
			direction: nil,
		},
	}
}

func (b Bunny) Render() Marker {
	var color string

	if b.Gender() == Female {
		color = Salmon
	} else {
		color = Cyan
	}

	return Marker{
		color,
		"b",
	}
}

func (b Bunny) IsAdult() bool {
	return b.Age() > 20
}

func (b *Bunny) Act(population []Entity) []Entity {
	return b.Animal.CommonAct(population, BunnyTicksBetweenMating, b)
}

func (b *Bunny) IsAtEndOfLife() bool {
	return b.Age() > 100
}
