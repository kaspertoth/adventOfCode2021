package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type crabCounter []int

func stringToCrabs(s string) crabCounter {
	var result crabCounter
	v := strings.Split(s, ",")
	for _, val := range v {
		curInt, _ := strconv.Atoi(val)
		result = append(result, curInt)
	}
	return result
}

func readFile() crabCounter {
	var resultNumbers crabCounter
	counter := 0
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if counter == 0 {
			resultNumbers = stringToCrabs(line)
		}
		counter++
	}
	return resultNumbers
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (numbers *crabCounter) max() int {
	result := 0
	for _, i2 := range *numbers {
		if i2 > result {
			result = i2
		}
	}
	return result
}

func (numbers *crabCounter) getCost(position int) int64 {
	var result int64
	var cost int
	for _, i2 := range *numbers {
		cost = abs(i2 - position)
		result += int64(cost)
	}
	return result
}

func (numbers *crabCounter) getCost2(position int) int64 {
	var result int64
	var distance, cost int
	for _, i2 := range *numbers {
		distance = abs(i2 - position)
		cost = (distance * (distance + 1)) / 2
		result += int64(cost)
	}
	return result
}

func main() {
	crabs := readFile()
	var cost, cost2, lowest, lowest2 int64
	max := crabs.max()
	for i := 1; i <= max; i++ {
		cost = crabs.getCost(i)
		cost2 = crabs.getCost2(i)
		if i == 1 || cost < lowest {
			lowest = cost
		}
		if i == 1 || cost2 < lowest2 {
			lowest2 = cost2
		}
	}
	fmt.Printf("Lowest: %d\n", lowest)
	fmt.Printf("Lowest 2: %d\n", lowest2)
}
