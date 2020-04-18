package simforest

import (
	"math/rand"
)

const FoxTicksBetweenMating = 40

type Fox struct {
	Animal
}

func (f *Fox) Mate(other Entity, population []Entity) []Entity {
	_, ok := other.(*Fox)
	if ok {
		if offSpringPosition := f.FindFreeNeighborTile(population, f.Env()); offSpringPosition != nil {
			f.tickLastMated = f.Env().tickCount
			return []Entity{NewFox(*offSpringPosition, f.Env())}
		}
	}
	return []Entity{}
}

func NewFox(p Position, e *Environment) *Fox {
	return &Fox{
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

func (f Fox) Render() Marker {
	var color string

	if f.Gender() == Female {
		color = LightPink
	} else {
		color = LightBrown
	}

	return Marker{
		color,
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
			if isBunny && b.IsNearby(f.Position) {
				b.Die()
				break
			}
		}
	}

	return offspring
}

func (f *Fox) IsAtEndOfLife() bool {
	return f.Age() > 200
}
