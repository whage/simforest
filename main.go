package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"strconv"
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
	foxCount = 20
	bunnyCount = 80
	gameLoopInterval = 200
)

var width = 50
var height = 25

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

func (p Position) IsNearby(o Position) bool {
	xWithin1 := math.Abs(float64(o.X-p.X)) <= 1
	yWithin1 := math.Abs(float64(o.Y-p.Y)) <= 1
	return xWithin1 && yWithin1
}

func (b *Bunny) TryToMate(others []Breeder) Breeder {
	if (b.gender == Male) {
		for _, o := range others {
			o, ok := o.(*Bunny)
			if ok {
				if b.Pos().IsNearby(o.Pos()) && (o.gender == Female) && b.age > 30 && o.age > 30 {
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

func (f *Fox) TryToMate(others []Breeder) Breeder {
	if (f.gender == Male) {
		for _, o := range others {
			o, ok := o.(*Fox)
			if ok {
				if f.Pos().IsNearby(o.Pos()) && (o.gender == Female) && f.age > 30 && o.age > 30 {
					return createBunny(Position{f.pos.X, f.pos.Y})
				}
			}
		}
	}
	return nil
}

func randomStep(p *Position) {
    steps := []int{-1, 0, 1}
	direction := Position{
        steps[rand.Intn(len(steps))],
        steps[rand.Intn(len(steps))],
    }

	if newPosition := p.Add(direction); newPosition.IsWithinEnvironment() {
		*p = p.Add(direction)
	}
}

func initBoardSize() {
	sttyCommand := exec.Command("stty", "size")
	sttyCommand.Stdin = os.Stdin
	out, err := sttyCommand.Output()
	if err != nil {
		fmt.Println("Error when running `stty size`: ", err)
		os.Exit(1)
	}
	terminalSize := strings.Split(strings.Trim(string(out), "\n"), " ")
	h, _ := strconv.Atoi(terminalSize[0])
	w, _ := strconv.Atoi(terminalSize[1])

	width, height = w-2, h-2
}

func main() {
	rand.Seed(time.Now().UnixNano())

	initBoardSize()
	population := initPopulation()
	
	for {
		drawWorld(population)
		time.Sleep(gameLoopInterval * time.Millisecond)
		population = step(population)
		// https://stackoverflow.com/a/33509850/1772429
		fmt.Printf("\033[0;0H")
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
	for y := 0; y < height; y++ {
		fmt.Print("#")
		for x := 0; x < width; x++ {
			fmt.Print(getMarker(x, y, population))
		}
		fmt.Println("#")
	}
	fmt.Print(strings.Repeat("#", width+2))
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
