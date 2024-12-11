package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules := make(map[int][]int)
	content, err := os.ReadFile("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	var i int
	for i = 0; lines[i] != ""; i++ {
		line := strings.Split(lines[i], "|")
		first, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		rules[first] = append(rules[first], second)
	}
	numbers := make([]int, 0, 32)
	var sum int
nextLine:
	for i = i + 1; i != len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		seq := strings.Split(lines[i], ",")
		numbers = numbers[:0]
		for j := 0; j != len(seq); j++ {
			number, err := strconv.Atoi(seq[j])
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, number)
			numberRules := rules[number]
			if numberRules == nil {
				continue
			}
			for k := 0; k != j; k++ {
				for l := 0; l != len(numberRules); l++ {
					if numbers[k] == numberRules[l] {
						continue nextLine
					}
				}
			}
		}
		sum += numbers[len(numbers)/2]
	}
	fmt.Println(sum)
}
