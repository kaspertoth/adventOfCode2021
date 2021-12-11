package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const size = 100

type grid [size][size]int
type freqs map[int]int

func (f freqs) print() {
	for i, i2 := range f {
		fmt.Printf("%d -> %d\n", i, i2)
	}
}

func (f *freqs) getScore() {
	var ordered []int
	for _, v := range *f {
		ordered = append(ordered, v)
	}
	sort.Ints(ordered)
	l := len(ordered)
	if ordered != nil && l > 0 {
		result := ordered[l-1] * ordered[l-2] * ordered[l-3]
		fmt.Printf("Part 2: %d\n", result)
	}
}

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
			g[curLine][i] = val
		}
		curLine++
	}
}

func (g *grid) print() {
	for _, r := range g {
		for _, c := range r {
			fmt.Printf("[%d]", c)
		}
		fmt.Printf("\n")
	}
}

func (g *grid) makeRegions() *grid {
	var result grid
	currentRegion := 0
	for y, row := range g {
		for x, val := range row {
			top := result.get(x, y-1, 0)
			left := result.get(x-1, y, 0)
			if val == 9 {
				result[y][x] = 0
			} else if top != 0 {
				result[y][x] = top
				if left != 0 && left != top {
					result.replace(left, top)
				}
			} else if left != 0 {
				result[y][x] = left
			} else {
				currentRegion++
				result[y][x] = currentRegion
			}
		}
	}
	return &result
}

func (g *grid) get(x int, y int, d int) int {
	if x < 0 || x >= size || y < 0 || y >= size {
		return d
	}
	return g[y][x]
}

func (g *grid) replace(find int, replace int) {
	for y, row := range g {
		for x, val := range row {
			if val == find {
				g[y][x] = replace
			}
		}
	}
}

func (g *grid) getFrequencies() map[int]int {
	var result = make(map[int]int)
	for _, row := range g {
		for _, val := range row {
			if val == 0 {
				continue
			}
			if count, ok := result[val]; ok {
				result[val] = count + 1
			} else {
				result[val] = 1
			}
		}
	}
	return result
}

func (g *grid) getLowScore() {
	var result int
	for y, row := range g {
		for x, val := range row {
			top := g.get(x, y-1, 9)
			left := g.get(x-1, y, 9)
			right := g.get(x+1, y, 9)
			bottom := g.get(x, y+1, 9)
			if val < top && val < bottom && val < left && val < right {
				result += val + 1
			}
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

func main() {
	var g, result grid
	var frequencies freqs
	g.processFile()
	g.getLowScore()
	result = *(g.makeRegions())
	frequencies = result.getFrequencies()
	frequencies.getScore()
}
