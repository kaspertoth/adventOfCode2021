package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 1500

type grid [size][size]byte

func (g *grid) processFile() (int, int) {
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	foldMode := false
	var curWidth, curHeight, dotCount int
	isFirst := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			foldMode = true
		} else if !foldMode {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			if x >= curWidth {
				curWidth = x + 1
			}
			y, _ := strconv.Atoi(parts[1])
			if y >= curHeight {
				curHeight = y + 1
			}
			g[x][y] = 1
		} else {
			var axis string
			var n int
			_, _ = fmt.Sscanf(line,
				"fold along %1s=%d", &axis, &n)
			curWidth, curHeight, dotCount = g.fold(axis, n, curWidth, curHeight)
			if isFirst {
				isFirst = false
				fmt.Printf("Answer 1: %d\n", dotCount)
			}
		}
	}
	return curWidth, curHeight
}

func (g *grid) print(width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if g[x][y] == 1 {
				fmt.Printf("██")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

func (g *grid) fold(axis string, n int, curWidth int, curHeight int) (int, int, int) {
	var dotCount int
	if axis == "y" {
		for x := 0; x < curWidth; x++ {
			for y := 0; y < n; y++ {
				if g[x][y] == 1 || g[x][curHeight-1-y] == 1 {
					g[x][y] = 1
					dotCount++
				}
			}
		}
		curHeight = n
	}
	if axis == "x" {
		for x := 0; x < n; x++ {
			for y := 0; y < curWidth; y++ {
				if g[x][y] == 1 || g[curWidth-1-x][y] == 1 {
					g[x][y] = 1
					dotCount++
				}
			}
		}
		curWidth = n
	}
	return curWidth, curHeight, dotCount
}

func main() {
	var g grid
	curWidth, curHeight := g.processFile()
	fmt.Printf("Answer 2:\n")
	g.print(curWidth, curHeight)
}
