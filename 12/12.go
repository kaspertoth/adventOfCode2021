package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type pathList []string
type pathsMap map[string]pathList

func isLc(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func (pl pathList) occurrenceCount(s string) int {
	var result int
	for _, p := range pl {
		if p == s {
			result++
		}
	}
	return result
}

func (paths *pathsMap) calculateRoutes(currentPath pathList, maxLc int) int {
	var result int
	current := currentPath[len(currentPath)-1]
	for _, s := range (*paths)[current] {
		newPath := make(pathList, len(currentPath))
		copy(newPath, currentPath)
		newPath = append(newPath, s)
		if s == "end" {
			result++
		} else if s == "start" {
			// Can't go into start again
		} else if isLc(s) {
			occurrences := currentPath.occurrenceCount(s)
			newMaxLc := maxLc
			if occurrences < maxLc {
				if occurrences >= 1 {
					newMaxLc = 1
				}
				result += paths.calculateRoutes(newPath, newMaxLc)
			}
		} else {
			result += paths.calculateRoutes(newPath, maxLc)
		}
	}
	return result
}

func processFile() pathsMap {
	paths := pathsMap{}
	f, _ := os.Open("input.txt")
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if _, ok := paths[parts[0]]; ok {
			paths[parts[0]] = append(paths[parts[0]], parts[1])
		} else {
			paths[parts[0]] = []string{parts[1]}
		}
		if _, ok := paths[parts[1]]; ok {
			paths[parts[1]] = append(paths[parts[1]], parts[0])
		} else {
			paths[parts[1]] = []string{parts[0]}
		}
	}
	return paths
}

func main() {
	paths := processFile()
	routeCount1 := paths.calculateRoutes(pathList{"start"}, 1)
	fmt.Printf("Answer 1: %d\n", routeCount1)
	routeCount2 := paths.calculateRoutes(pathList{"start"}, 2)
	fmt.Printf("Answer 2: %d\n", routeCount2)
}
