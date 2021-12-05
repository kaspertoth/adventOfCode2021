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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g *grid) addLine(l line) {
	if l.x1 != l.x2 && l.y1 != l.y2 {
		return
	}
	yStart := min(l.y1, l.y2)
	yEnd := max(l.y1, l.y2)
	xStart := min(l.x1, l.x2)
	xEnd := max(l.x1, l.x2)
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			g[y][x] += 1
		}
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
