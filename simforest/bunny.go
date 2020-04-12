package simforest

import (
	"math/rand"
)

type Bunny struct {
	Animal
}

func (b *Bunny) Mate(other Creature) Creature {
	_, ok := other.(*Bunny)
	if ok {
		return NewBunny(b.Pos(), b.Env())
	}
	return nil
}

func NewBunny(p Position, e *Environment) *Bunny {
	return &Bunny{
		Animal{
			e,
			Gender(rand.Intn(2)),
			p,
			0,
		},
	}
}
