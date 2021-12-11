package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const size = 10

type octopus struct {
	value      int
	hasFlashed bool
}

type grid [size][size]octopus

var totalFlashes int

func (g *grid) processFile() {
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	curLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			val, _ := strconv.Atoi(string(r))
			g[curLine][i] = octopus{val, false}
		}
		curLine++
	}
}

func (g *grid) print() {
	for _, r := range g {
		for _, c := range r {
			fmt.Printf("[%d]", c.value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (g *grid) increase(x int, y int) {
	if x < 0 || x >= size || y < 0 || y >= size {
		return
	}
	g[y][x].value++
	return
}

func (g *grid) replaceFlashes() bool {
	allFlashed := true
	for y, row := range g {
		for x, val := range row {
			if val.value > 9 {
				g[y][x].value = 0
				g[y][x].hasFlashed = false
				totalFlashes++
			} else {
				allFlashed = false
			}
		}
	}
	return allFlashed
}

func (g *grid) flash(x, y, increase int) {
	if x < 0 || x >= size || y < 0 || y >= size {
		return
	}
	g[y][x].value += increase
	if g[y][x].value > 9 && !g[y][x].hasFlashed {
		g[y][x].hasFlashed = true
		g.flash(x-1, y-1, 1)
		g.flash(x-1, y, 1)
		g.flash(x-1, y+1, 1)
		g.flash(x, y-1, 1)
		g.flash(x, y+1, 1)
		g.flash(x+1, y-1, 1)
		g.flash(x+1, y, 1)
		g.flash(x+1, y+1, 1)
	}
}

func (g *grid) iterate() bool {
	// Increase all by 1:
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			g[y][x].value++
		}
	}
	// Flash what needs to be flashed:
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if g[y][x].value > 9 {
				g.flash(x, y, 0)
			}
		}
	}
	return g.replaceFlashes()
}

func main() {
	var g grid
	g.processFile()
	for i := 0; true; i++ {
		if i == 100 {
			fmt.Printf("Answer 1: %d\n", totalFlashes)
		}
		if g.iterate() {
			fmt.Printf("Answer 2: %d\n", i+1)
			break
		}

	}
}
