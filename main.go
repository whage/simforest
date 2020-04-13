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

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
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

func GetMarker(c simforest.Creature) string {
	var colorCode string
	sign := c.Render()

	if !c.IsAlive() {
		return " "
	}

	if c.Gender() == simforest.Female {
		colorCode = ErrorColor
	} else {
		colorCode = InfoColor
	}

	if c.IsAdult() {
		return fmt.Sprintf(colorCode, strings.ToUpper(sign))
	} else {
		return fmt.Sprintf(colorCode, sign)
	}
}

func getMarker(x, y int, population []simforest.Creature) string {
	for _, c := range population {
		if c.Pos().X == x && c.Pos().Y == y {
			return GetMarker(c)
		}
	}
	return " "
}
