package main

import (
	"fmt"
	"math"
	"math/rand"
	//"os"
	//"strconv"
	"strings"
	"time"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	width = 50
	height = 25
	foxCount = 2
	bunnyCount = 8
	gameLoopInterval = 100
)

type Position struct {
	X, Y int
}

type Gender int

const (
	Male = iota
	Female
)

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

func (p *Position) IsWithinEnvironment() bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

type Mover interface {
	Pos() Position
	Move()
}

type Breeder interface {
	TryToMate(others []Breeder) Breeder
}

type Ager interface {
	IncreaseAge()
}

type Creature interface {
	Mover
	Breeder
	Ager
}

type Population []Creature

func (p Population) FilterBunnies() []Bunny {
	var results []Bunny
	for i, _ := range p {
		v, ok := p[i].(*Bunny)
		if ok {
			results = append(results, *v)
		}
	}
	return results
}

type Fox struct {
	gender Gender
	pos Position
	age int
}

type Bunny struct {
	gender Gender
	pos Position
	age int
}

func (b *Bunny) Move() {
	randomStep(&b.pos)
}

func (b *Bunny) Pos() Position {
	return b.pos
}

func (b *Bunny) IncreaseAge() {
	b.age += 1
}

func (b *Bunny) IsNearby(o *Bunny) bool {
	xWithin1 := math.Abs(float64(o.pos.X-b.pos.X)) <= 1
	yWithin1 := math.Abs(float64(o.pos.Y-b.pos.Y)) <= 1
	return xWithin1 && yWithin1
}

func (b *Bunny) TryToMate(others []Breeder) Breeder {
	if (b.gender == Male) {
		for _, o := range others {
			o, ok := o.(*Bunny)
			if ok {
				if b.IsNearby(o) && (o.gender == Female) && b.age > 30 && o.age > 30 {
					return createBunny(Position{b.pos.X, b.pos.Y})
				}
			}
		}
	}
	return nil
}

func (f *Fox) Pos() Position {
	return f.pos
}

func (f *Fox) Move() {
	randomStep(&f.pos)
}

func (f *Fox) IncreaseAge() {
	f.age += 1
}

func (f *Fox) IsNearby(o *Fox) bool {
	xWithin1 := math.Abs(float64(o.pos.X-f.pos.X)) <= 1
	yWithin1 := math.Abs(float64(o.pos.Y-f.pos.Y)) <= 1
	return xWithin1 && yWithin1
}

func (f *Fox) TryToMate(others []Breeder) Breeder {
	if (f.gender == Male) {
		for _, o := range others {
			o, ok := o.(*Fox)
			if ok {
				if f.IsNearby(o) && (o.gender == Female) && f.age > 30 && o.age > 30 {
					return createBunny(Position{f.pos.X, f.pos.Y})
				}
			}
		}
	}
	return nil
}

func randomStep(p *Position) {
	directions := []Position{
		{0, 0},
		{0, 1},
		{0, -1},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
	}

	direction := directions[rand.Intn(len(directions))]
	newPosition := p.Add(direction)

	if newPosition.IsWithinEnvironment() {
		*p = p.Add(direction)
	}
}

var environment [width][height]string

func main() {
	rand.Seed(time.Now().UnixNano())
	population := initPopulation()
	
	for {
		drawWorld(population)
		time.Sleep(gameLoopInterval * time.Millisecond)
		population = step(population)
		fmt.Println("Population count: ", len(population))
		//fmt.Println("Age of first: ", population[0].age)
	}
}

func step(population Population) Population {
	var newPopulation Population

	for i, _ := range population {
		population[i].Move()
		population[i].IncreaseAge()

		switch current := population[i].(type) {
		//case *Fox:
			// TODO: continue (as with *Bunny below)
		case *Bunny:
			bunnies := population.FilterBunnies()

			// Convert Bunnies to Breeders to satisfy TryToMate() signature
			breeders := make([]Breeder, len(bunnies), len(bunnies))
			for i := range bunnies {
				breeders[i] = &bunnies[i]
			}
			if newBreeder := current.TryToMate(breeders); newBreeder != nil {
				// Can this type assertion fail at runtime?
				// If not, how does the type checker know?
				newPopulation = append(newPopulation, newBreeder.(*Bunny))
			}
		}

		// Carry current object over to population of next round
		newPopulation = append(newPopulation, population[i])
	}

	return newPopulation
}

func createBunny(p Position) *Bunny {
	return &Bunny{
		Gender(rand.Intn(2)),
		p,
		0,
	}
}

func initPopulation() Population {
	population := make(Population, 0)

	for i := 0; i < foxCount; i++ {
		population = append(population, &Fox{Gender(rand.Intn(2)), Position{rand.Intn(width), rand.Intn(height)}, 0})
	}

	for i := 0; i < bunnyCount; i++ {
		population = append(population, createBunny(Position{rand.Intn(width), rand.Intn(height)}))
	}

	return population
}

func drawWorld(population Population) {
	fmt.Println(strings.Repeat("#", width+2))
	for y := 0; y < 25; y++ {
		fmt.Print("#")
		for x := 0; x < 50; x++ {
			fmt.Print(getMarker(x, y, population))
		}
		fmt.Println("#")
	}
	fmt.Println(strings.Repeat("#", width+2))
}

func getMarker(x, y int, population Population) string {
	for _, c := range population {
		if c.Pos().X == x && c.Pos().Y == y {
			switch c.(type) {
				case *Fox:
					return "F"
				case *Bunny:
					b := c.(*Bunny)
					if b.gender == Female {
						if b.age < 30 {
							return fmt.Sprintf(ErrorColor, "b")
						} else {
							return fmt.Sprintf(ErrorColor, "B")
						}
					} else {
						if b.age < 30 {
							return fmt.Sprintf(InfoColor, "b")
						} else {
							return fmt.Sprintf(InfoColor, "B")
						}
					}
				default:
					return " "
			}
		}
	}
	return " "
}
