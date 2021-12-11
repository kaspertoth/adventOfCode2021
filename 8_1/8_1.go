package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type segment []string

func stringToSegment(s string) segment {
	var result segment
	v := strings.Split(s, " ")
	for _, val := range v {
		if val != "|" {
			result = append(result, val)
		}
	}
	return result
}

func processFile() int {
	result := 0
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		seg := stringToSegment(line)
		for i, s := range seg {
			if i > 9 {
				l := len(s)
				if l == 2 || l == 3 || l == 4 || l == 7 {
					result++
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Printf("1: %d\n", processFile())
}
