package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type runes []rune
type segment []runes

type frequencies map[rune]int

var signatureLookup = map[int]int{
	17: 1,
	34: 2,
	39: 3,
	30: 4,
	37: 5,
	41: 6,
	25: 7,
	49: 8,
	45: 9,
}

func (f *frequencies) add(s runes) {
	for _, c := range s {
		(*f)[c]++
	}
}

func (f *frequencies) val(s runes) int {
	result := 0
	for _, c := range s {
		result += (*f)[c]
	}
	return result
}

func stringToSegment(s string) segment {
	var result segment
	v := strings.Split(s, " ")
	for _, val := range v {
		if val != "|" {
			result = append(result, runes(val))
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
		freqs := frequencies{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0}
		line := scanner.Text()
		seg := stringToSegment(line)
		for _, s := range seg[0:10] {
			freqs.add(s)
		}
		output := ""
		for _, s := range seg[10:] {
			val := freqs.val(s)
			itm := signatureLookup[val]
			output = output + strconv.Itoa(itm)
		}
		lineResult, _ := strconv.Atoi(output)
		result += lineResult
	}
	return result
}

func main() {
	fmt.Printf("%d\n", processFile())
}
