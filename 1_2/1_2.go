package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const slideCount = 3

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	var slides [slideCount]int
	currentTotal := 0
	scanner := bufio.NewScanner(f)
	lastTotal := 0
	totalIncrements := 0
	index := 0
	for scanner.Scan() {
		currentValue, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < slideCount; i++ {
			slides[i] = slides[i] + currentValue
			if index%slideCount == i {
				currentTotal = slides[i]
				slides[i] = 0
			}
		}
		if index >= slideCount && currentTotal > lastTotal {
			totalIncrements++
		}
		lastTotal = currentTotal
		index++
	}
	fmt.Println(totalIncrements)
}
