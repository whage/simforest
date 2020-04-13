package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"time"

	"github.com/whage/simforest/simforest"
)

const (
	gameLoopInterval = 50
)

func main() {
	rand.Seed(time.Now().UnixNano())

	environment := NewEnvironment()
	population := simforest.InitPopulation(environment)
	
	for {
		drawWorld(population, environment)
		time.Sleep(gameLoopInterval * time.Millisecond)
		population = simforest.Tick(population, environment)
		// https://stackoverflow.com/a/33509850/1772429
		fmt.Printf("\033[0;0H")
	}
}

func NewEnvironment() *simforest.Environment {
	sttyCommand := exec.Command("stty", "size")
	sttyCommand.Stdin = os.Stdin
	out, err := sttyCommand.Output()
	if err != nil {
		fmt.Println("Error when running `stty size`: ", err)
		os.Exit(1)
	}
	terminalSize := strings.Split(strings.Trim(string(out), "\n"), " ")
	w, _ := strconv.Atoi(terminalSize[1])
	h, _ := strconv.Atoi(terminalSize[0])

	return simforest.CreateEnvironment(w, h)
}

func drawWorld(population []simforest.Creature, e *simforest.Environment) {
	for y := 0; y < e.Height(); y++ {
		for x := 0; x < e.Width(); x++ {
			fmt.Print(getMarker(x, y, population))
		}
	}
}

func getMarker(x, y int, population []simforest.Creature) string {
	for _, c := range population {
		if c.Pos().X == x && c.Pos().Y == y {
			return c.Render()
		}
	}
	return " "
}
