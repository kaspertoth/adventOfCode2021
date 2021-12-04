package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type number struct {
	nr     int
	picked bool
}

type grid [5][5]number

func stringToResultNumbers(s string) []int {
	var result []int
	v := strings.Split(s, ",")
	for _, val := range v {
		curInt, _ := strconv.Atoi(val)
		result = append(result, curInt)
	}
	return result
}

func stringToGridLine(s string) [5]number {
	var result [5]number
	v := strings.Split(s, " ")
	curPos := 0
	for _, val := range v {
		if len(val) >= 1 {
			curInt, err := strconv.Atoi(val)
			if err == nil {
				result[curPos] = number{
					nr:     curInt,
					picked: false,
				}
				curPos++
			}
		}
	}
	return result
}

func readFile() ([]int, []grid) {
	var resultNumbers []int
	var resultGrids []grid
	var currentGrid grid
	counter := 0
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		relativePosition := (counter - 1) % 6
		if counter == 0 {
			resultNumbers = stringToResultNumbers(line)
		} else if relativePosition >= 1 {
			gridLine := stringToGridLine(line)
			currentGrid[relativePosition-1] = gridLine
			if relativePosition == 5 {
				resultGrids = append(resultGrids, currentGrid)
				currentGrid = grid{}
			}
		}
		counter++
	}
	return resultNumbers, resultGrids
}

func printGrid(g grid) {
	println("GRID")
	for _, numbers := range g {
		for _, n := range numbers {
			print(n.nr)
			if n.picked {
				print("*")
			}
			print(", ")
		}
		print("\n")
	}
}

func addNumberToGrid(g *grid, picked int) {
	for y := 0; y < len(*g); y++ {
		for x := 0; x < len((*g)[y]); x++ {
			if (*g)[y][x].nr == picked {
				(*g)[y][x].picked = true
			}
		}
	}
}

func calculateScore(g *grid, n int) int {
	result := 0
	for y := 0; y < len(*g); y++ {
		for x := 0; x < len((*g)[y]); x++ {
			if (*g)[y][x].picked == false {
				result += (*g)[y][x].nr
			}
		}
	}
	return result * n
}

func isWinning(g *grid) bool {
	for y := 0; y < len(*g); y++ {
		successful := true
		for x := 0; x < len((*g)[y]); x++ {
			if (*g)[y][x].picked == false {
				successful = false
				break
			}
		}
		if successful {
			return true
		}
	}
	for x := 0; x < len((*g)[0]); x++ {
		successful := true
		for y := 0; y < len(*g); y++ {
			if (*g)[y][x].picked == false {
				successful = false
				break
			}
		}
		if successful {
			return true
		}
	}
	return false
}

func main() {
	numbers, grids := readFile()
	for _, n := range numbers {
		for i := 0; i < len(grids); i++ {
			addNumberToGrid(&(grids[i]), n)
			if isWinning(&(grids[i])) {
				println("Winner")
				printGrid(grids[i])
				println(n)
				print("Result: ")
				println(calculateScore(&(grids[i]), n))
				return
			}
		}
	}
}
