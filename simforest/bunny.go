package simforest

import (
	"math/rand"
)

const BunnyTicksBetweenMating = 20

type Bunny struct {
	Animal
}

func (b *Bunny) Mate(other Entity) []Entity {
	_, ok := other.(*Bunny)
	if ok {
		b.tickLastMated = b.Env().tickCount
		return []Entity{NewBunny(b.Pos(), b.Env())}
	}
	return []Entity{}
}

func NewBunny(p Position, e *Environment) *Bunny {
	return &Bunny{
		Animal{
			e,
			Gender(rand.Intn(2)),
			p,
			0,
			e.tickCount,
			true,
			e.tickCount,
		},
	}
}

func (b Bunny) Render() Marker {
	return Marker{
		Cyan,
		"b",
	}
}

func (b Bunny) IsAdult() bool {
	return b.Age() > 20
}

func (b *Bunny) Act(population []Entity) []Entity {
	return b.Animal.CommonAct(population, BunnyTicksBetweenMating, b)
}
