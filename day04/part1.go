package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func checkXmas(lines []string, i, j, dy, dx int) bool {
	n := len(lines)
	line := lines[i]
	m := len(line)
	if dy > 0 {
		if i+3 >= n {
			return false
		}
	} else if dy < 0 {
		if i-3 < 0 {
			return false
		}
	}
	if dx > 0 {
		if j+3 >= m {
			return false
		}
	} else if dx < 0 {
		if j-3 < 0 {
			return false
		}
	}
	return lines[i+dy][j+dx] == 'M' &&
		lines[i+2*dy][j+2*dx] == 'A' &&
		lines[i+3*dy][j+3*dx] == 'S'
}

func countXmas(lines []string, i, j int) int {
	var count int
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if checkXmas(lines, i, j, dy, dx) {
				count++
			}
		}
	}
	return count
}

func main() {
	content, err := os.ReadFile("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := slices.DeleteFunc(
		strings.Split(string(content), "\n"),
		func(s string) bool { return s == "" })
	var count int
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if line[j] == 'X' {
				count += countXmas(lines, i, j)
			}
		}
	}
	fmt.Println(count)
}
