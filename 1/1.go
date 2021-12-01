package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)
	lastValue := 0
	totalIncrements := 0
	first := true
	for scanner.Scan() {
		currentValue, _ := strconv.Atoi(scanner.Text())
		if first {
			first = false
		} else {
			if currentValue > lastValue {
				totalIncrements++
			}
		}
		lastValue = currentValue
	}
	fmt.Println(totalIncrements)
}