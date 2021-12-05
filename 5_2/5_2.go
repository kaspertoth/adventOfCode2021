package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

const gridSize = 1000

type grid [gridSize][gridSize]int

func compare(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func (g *grid) addLine(l line) {
	numberOfMoves := l.y2 - l.y1
	directionX := compare(l.x2, l.x1)
	directionY := compare(l.y2, l.y1)
	if directionY == 0 {
		numberOfMoves = l.x2 - l.x1
	}
	if numberOfMoves < 0 {
		numberOfMoves = -numberOfMoves
	}
	numberOfMoves++
	for i := 0; i < numberOfMoves; i++ {
		g[l.y1+i*directionY][l.x1+i*directionX]++
	}
}

func (g *grid) overlapCount() int {
	result := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if g[y][x] > 1 {
				result++
			}
		}
	}
	return result
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var l line
	var g grid
	for scanner.Scan() {
		_, err := fmt.Sscanf(
			scanner.Text(), "%d,%d -> %d,%d", &l.x1, &l.y1, &l.x2, &l.y2,
		)
		if err != nil {
			continue
		}
		g.addLine(l)
	}
	fmt.Printf("Overlapcount: %d\n", g.overlapCount())
}
