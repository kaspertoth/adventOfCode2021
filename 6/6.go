package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxDays = 9

type fishCounter [maxDays]int64

func stringToResultNumbers(s string) fishCounter {
	var result fishCounter
	v := strings.Split(s, ",")
	for _, val := range v {
		curInt, _ := strconv.Atoi(val)
		result[curInt]++
	}
	return result
}

func readFile() fishCounter {
	var resultNumbers fishCounter
	counter := 0
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if counter == 0 {
			resultNumbers = stringToResultNumbers(line)
		}
		counter++
	}
	return resultNumbers
}

func iterate(numbers fishCounter) fishCounter {
	var result fishCounter
	for i := 0; i < 9; i++ {
		if i == 0 {
			result[6] += numbers[i]
			result[8] += numbers[i]
		} else {
			result[i-1] += numbers[i]
		}
	}
	return result
}

func getCount(numbers fishCounter) int64 {
	var result int64
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	numbers := readFile()
	for i := 1; i <= 256; i++ {
		numbers = iterate(numbers)
		if i == 80 {
			fmt.Printf("After 80: %d\n", getCount(numbers))
		}
	}
	fmt.Printf("After 256: %d\n", getCount(numbers))
}
