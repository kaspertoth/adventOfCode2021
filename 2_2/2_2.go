package main

import (
	"bufio"
	"fmt"
	"os"
)

var vectors = map[string][2]int{
	"forward": {1, 0},
	"down":    {0, 1},
	"up":      {0, -1},
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var job string
	var amount int
	var line string

	horizontal := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		line = scanner.Text()
		_, err := fmt.Sscanf(line, "%s %d", &job, &amount)
		if err != nil {
			continue
		}
		v := vectors[job]
		horizontal = horizontal + (v[0] * amount)
		depth = depth + (v[0] * aim * amount)
		aim = aim + (v[1] * amount)
	}
	fmt.Println(horizontal * depth)
}
