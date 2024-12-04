package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func checkX(lines []string, i, j int, a, b, c, d byte) bool {
	if lines[i-1][j-1] == a &&
		lines[i-1][j+1] == b &&
		lines[i+1][j+1] == c &&
		lines[i+1][j-1] == d {
		return true
	}
	return false
}

func checkXmas(lines []string, i, j int) bool {
	n := len(lines)
	line := lines[i]
	m := len(line)
	if j == 0 || j == m-1 {
		return false
	}
	if i == 0 || i == n-1 {
		return false
	}
	if checkX(lines, i, j, 'M', 'S', 'S', 'M') ||
		checkX(lines, i, j, 'M', 'M', 'S', 'S') ||
		checkX(lines, i, j, 'S', 'M', 'M', 'S') ||
		checkX(lines, i, j, 'S', 'S', 'M', 'M') {
		return true
	}
	return false
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
	for i := 1; i < len(lines)-1; i++ {
		line := lines[i]
		for j := 1; j < len(line)-1; j++ {
			if line[j] == 'A' && checkXmas(lines, i, j) {
				count++
			}
		}
	}
	fmt.Println(count)
}
