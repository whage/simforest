package simforest

import (
	"math/rand"
)

type Fox struct {
	Animal
}

func (f *Fox) Mate(other Creature) Creature {
	_, ok := other.(*Fox)
	if ok {
		return NewFox(f.Pos(), f.Env())
	}
	return nil
}

func NewFox(p Position, e *Environment) *Fox {
	return &Fox{
		Animal{
			e,
			Gender(rand.Intn(2)),
			p,
			0,
		},
	}
}
