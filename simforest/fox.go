package simforest

import (
	"math/rand"
)

const FoxTicksBetweenMating = 40

type Fox struct {
	Animal
}

func (f *Fox) Mate(other Creature) Creature {
	_, ok := other.(*Fox)
	if ok {
		f.tickLastMated = f.Env().tickCount
		return NewFox(f.Pos(), f.Env())
	}
	return nil
}

func (b *Fox) CanStartMating() bool {
	return b.Animal.CanStartMating(FoxTicksBetweenMating)
}

func NewFox(p Position, e *Environment) *Fox {
	return &Fox{
		Animal{
			e,
			Gender(rand.Intn(2)),
			p,
			0,
			e.tickCount,
		},
	}
}

func (f Fox) Render() string {
	return f.Animal.Render("f")
}
