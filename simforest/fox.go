package simforest

import (
	"math/rand"
)

const FoxTicksBetweenMating = 40

type Fox struct {
	Animal
}

func (f *Fox) Mate(other Entity) []Entity {
	_, ok := other.(*Fox)
	if ok {
		f.tickLastMated = f.Env().tickCount
		return []Entity{NewFox(f.Pos(), f.Env())}
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
			e.tickCount,
		},
	}
}

func (f Fox) Render() Marker {
	return Marker{
		LightBrown,
		"f",
	}
}

func (f Fox) IsAdult() bool {
	return f.Age() > 30
}

func (f *Fox) Act(population []Entity) []Entity {
	offspring := f.Animal.CommonAct(population, FoxTicksBetweenMating, f)

	// eat a nearby Bunny!
	if f.IsAdult() {
		for _, c := range population {
			b, isBunny := c.(*Bunny)
			if isBunny && b.Pos().IsNearby(f.Pos()) {
				b.Die()
				break
			}
		}
	}

	return offspring
}
