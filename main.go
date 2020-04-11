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
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

const (
	gameLoopInterval = 200
)

func main() {
	rand.Seed(time.Now().UnixNano())

	environment := NewEnvironment()
	population := simforest.InitPopulation(environment)
	
	for {
		drawWorld(population, environment)
		time.Sleep(gameLoopInterval * time.Millisecond)
		population = simforest.Step(population)
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
			switch c.(type) {
				case *simforest.Fox:
					return "F"
				case *simforest.Bunny:
					b := c.(*simforest.Bunny)
					if b.Animal.Gender() == simforest.Female {
						if b.Animal.Age() < 30 {
							return fmt.Sprintf(ErrorColor, "b")
						} else {
							return fmt.Sprintf(ErrorColor, "B")
						}
					} else {
						if b.Animal.Age() < 30 {
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
