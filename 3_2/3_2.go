package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const bitSize = 12

type bitLine [bitSize]int

var wg sync.WaitGroup

func readFile() []bitLine {
	var result []bitLine
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var value bitLine
		for i, c := range scanner.Text() {
			if c == '1' {
				value[i] = 1
			}
		}
		result = append(result, value)
	}
	return result
}

func getMostCommon(values []bitLine) bitLine {
	var result bitLine
	for _, row := range values {
		for bitPosition, val := range row {
			result[bitPosition] += val
		}
	}
	for i, v := range result {
		if len(values)-v <= v {
			result[i] = 1
		}
	}
	return result
}

func getMeasurement(values []bitLine, lookingFor bool, result *int) {
	defer wg.Done()
	for i := 0; i < bitSize; i++ {
		mostCommon := getMostCommon(values)
		var newValues []bitLine
		for _, row := range values {
			if (row[i] == mostCommon[i]) == lookingFor {
				newValues = append(newValues, row)
			}
		}
		if len(newValues) == 1 {
			for _, bit := range newValues[0] {
				*result = (*result)<<1 + bit
			}
			break
		}
		values = newValues
	}
}

func main() {
	values := readFile()
	wg.Add(2)
	var oxygenValue int
	var co2Value int
	go getMeasurement(values, true, &oxygenValue)
	go getMeasurement(values, false, &co2Value)
	wg.Wait()
	fmt.Printf("Result: %d\n", oxygenValue*co2Value)
}
