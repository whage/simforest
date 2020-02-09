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
	bunnyCount = 4
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

type Creature interface {
	GetPos() Position
	Move()
}

type Fox struct {
	pos Position
}

type Bunny struct {
	gender Gender
	pos Position
}

func (b *Bunny) Move() {
	randomStep(&b.pos)
}

func (b *Bunny) GetPos() Position {
	return b.pos
}

func (f *Fox) GetPos() Position {
	return f.pos
}

func (f *Fox) Move() {
	randomStep(&f.pos)
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

	// TODO: implement moving towards potential mate if within radius

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
		step(population)
	}
}

func step(population []Creature) {
	for i, _ := range population {
		population[i].Move()
	}
}

func initPopulation() []Creature {
	population := make([]Creature, 0)

	for i := 0; i < foxCount; i++ {
		population = append(population, &Fox{Position{rand.Intn(width), rand.Intn(height)}})
	}

	for i := 0; i < bunnyCount; i++ {
		population = append(population, &Bunny{
			Gender(rand.Intn(2)),
			Position{rand.Intn(width), rand.Intn(height)},
		})
	}

	return population
}

func drawWorld(population []Creature) {
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

func getMarker(x, y int, population []Creature) string {
	for _, c := range population {
		if c.GetPos().X == x && c.GetPos().Y == y {
			switch c.(type) {
				case *Fox:
					return "F"
				case *Bunny:
					b := c.(*Bunny)
					if b.gender == Female {
						return fmt.Sprintf(ErrorColor, "B")
					} else {
						return fmt.Sprintf(InfoColor, "B")
					}
				default:
					return " "
			}
		}
	}
	return " "
}
