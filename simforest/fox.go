package simforest

import (
	"math/rand"
)

const FoxTicksBetweenMating = 40

type Fox struct {
	Animal
}

func (f *Fox) Mate(other Creature) []Creature {
	_, ok := other.(*Fox)
	if ok {
		f.tickLastMated = f.Env().tickCount
		return []Creature{NewFox(f.Pos(), f.Env())}
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
			e.tickCount,
			true,
		},
	}
}

func (f Fox) Render() string {
	return f.Animal.Render("f")
}

func (f *Fox) Act(population []Creature) []Creature {
	offspring := f.Animal.CommonAct(population, FoxTicksBetweenMating, f)

	// eat a nearby Bunny!
	for _, c := range population {
		b, isBunny := c.(*Bunny)
		if isBunny && b.Pos().IsNearby(f.Pos()) {
			b.Die()
			break
		}
	}

	return offspring
}
