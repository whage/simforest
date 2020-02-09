package main

import (
	"fmt"
	"math/rand"
	//"os"
	//"strconv"
	"strings"
	"time"
)

const (
	width = 50
	height = 25
	foxCount = 8
	gameLoopInterval = 200
)

type Position struct {
	X, Y int
}

type Fox struct {
	pos Position
}

func (f *Fox) Move() {
	if rand.Intn(2) == 0 {
		f.pos.X--
	} else {
		f.pos.X++
	}

	if rand.Intn(2) == 0 {
		f.pos.Y--
	} else {
		f.pos.Y++
	}
}

var environment [width][height]string

func main() {
	rand.Seed(time.Now().UnixNano())
	foxes := initPopulation()
	
	for {
		drawWorld(foxes)
		time.Sleep(gameLoopInterval * time.Millisecond)
		step(foxes)
	}
}

func step(foxes []Fox) {
	for i, _ := range foxes {
		foxes[i].Move()
	}
}

func initPopulation() []Fox {
	foxes := make([]Fox, 0)

	for i := 0; i < foxCount; i++ {
		foxes = append(foxes, Fox{Position{rand.Intn(width), rand.Intn(height)}})
	}

	return foxes
}

func drawWorld(foxes []Fox) {
	fmt.Println(strings.Repeat("#", width+2))
	for y := 0; y < 25; y++ {
		fmt.Print("#")
		for x := 0; x < 50; x++ {
			if isFox(x, y, foxes) {
				fmt.Print("F")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("#")
	}
	fmt.Println(strings.Repeat("#", width+2))
}

func isFox(x, y int, foxes []Fox) bool {
	for _, f := range foxes {
		if f.pos.X == x && f.pos.Y == y {
			return true
		}
	}

	return false
}
