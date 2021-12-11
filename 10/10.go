package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type token struct {
	closer rune
	score  int
}

var tokens = map[rune]token{
	'(': {')', 1},
	'[': {']', 2},
	'{': {'}', 3},
	'<': {'>', 4},
}

var breakScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func processFile() {
	score := 0
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	curLine := 0
	var goodScores []int
	for scanner.Scan() {
		line := scanner.Text()
		breakScore, goodScore := processLine(line)
		score += breakScore
		if goodScore > 0 {
			goodScores = append(goodScores, goodScore)
		}
		curLine++
	}
	sort.Ints(goodScores)
	fmt.Printf("Score: %d\n", score)
	if goodScores != nil {
		fmt.Printf("Score 2: %d\n", goodScores[len(goodScores)/2])
	}
}

func processLine(line string) (int, int) {
	var stack []token
	for _, r := range line {
		if t, ok := tokens[r]; ok {
			stack = append(stack, t)
		} else if len(stack) > 0 && r == stack[len(stack)-1].closer {
			stack = stack[:len(stack)-1]
		} else {
			return breakScores[r], 0
		}
	}
	score := 0
	if len(stack) > 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + stack[i].score
		}
	}
	return 0, score
}

func main() {
	processFile()
}
