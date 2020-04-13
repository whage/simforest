package simforest

import (
	"math/rand"
)

const BunnyTicksBetweenMating = 20

type Bunny struct {
	Animal
}

func (b *Bunny) Mate(other Creature) []Creature {
	_, ok := other.(*Bunny)
	if ok {
		b.tickLastMated = b.Env().tickCount
		return []Creature{NewBunny(b.Pos(), b.Env())}
	}
	return []Creature{}
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
		},
	}
}

func (b Bunny) Render() string {
	return b.Animal.Render("b")
}

func (b *Bunny) Act(population []Creature) []Creature {
	return b.Animal.CommonAct(population, BunnyTicksBetweenMating, b)
}
