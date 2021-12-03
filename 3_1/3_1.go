package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	// Get input:
	scanner := bufio.NewScanner(f)
	var values [12]int
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '1' {
				values[i]++
			}
		}
		lineCount++
	}
	// Determine most common bits:
	result := ""
	for _, v := range values {
		if v > lineCount/2 {
			result = result + "1"
		} else {
			result = result + "0"
		}
	}
	// Binary operating:
	binMap, _ := strconv.ParseInt("111111111111", 2, 64)
	gamma, _ := strconv.ParseInt(result, 2, 64)
	epsilon := (gamma ^ binMap) & binMap
	print(gamma * epsilon)
}
