package simforest

import (
	"math"
)

type Position struct {
	X, Y int
}

func (p Position) Add(other Position) Position {
	p.X += other.X
	p.Y += other.Y
	return p
}

func (p Position) Subtract(other Position) Position {
	p.X -= other.X
	p.Y -= other.Y
	return p
}

func (p Position) GetDistance(other Position) float64 {
	difference := p.Subtract(other)
	return math.Sqrt(float64(difference.X*difference.X) + float64(difference.Y*difference.Y))
}

func (p *Position) IsWithinEnvironment(e *Environment) bool {
	return p.X >= 0 && p.X < e.width && p.Y >= 0 && p.Y < e.height
}

func (p *Position) IsTaken(population []Creature) bool {
	for _, c := range population {
		if c.Pos() == *p {
			return true
		}
	}
	return false
}

func (p Position) IsNearby(o Position) bool {
	xWithin1 := math.Abs(float64(o.X-p.X)) <= 1
	yWithin1 := math.Abs(float64(o.Y-p.Y)) <= 1
	return xWithin1 && yWithin1
}
