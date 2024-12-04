package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = input.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(input)
	lefts := make([]uint, 0, 1001)
	rights := make(map[uint]uint, 1001)
	var l, r uint
	for scanner.Scan() {
		line := scanner.Text()
		if n, err := fmt.Sscan(line, &l, &r); err != nil || n != 2 {
			log.Fatal(err)
		}
		lefts = append(lefts, l)
		rights[r]++
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	var score uint
	for _, l := range lefts {
		count, ok := rights[l]
		if ok {
			score += l * count
		}
	}
	fmt.Println(score)
}
