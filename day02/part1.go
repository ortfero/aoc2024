package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafeLevel(lastLevel, level uint64, increasing bool) bool {
	var diff uint64
	if level >= lastLevel {
		if !increasing {
			return false
		}
		diff = level - lastLevel
	} else {
		if increasing {
			return false
		}
		diff = lastLevel - level
	}
	return diff > 0 && diff < 4
}

func main() {
	input, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = input.Close(); err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(input)
	var counter uint64
NextLine:
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		lastLevel, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		level, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		increasing := level >= lastLevel
		if !isSafeLevel(lastLevel, level, increasing) {
			continue NextLine
		}
		lastLevel = level
		for i := 2; i < len(fields); i++ {
			level, err := strconv.ParseUint(fields[i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if !isSafeLevel(lastLevel, level, increasing) {
				continue NextLine
			}
			lastLevel = level
		}
		counter++
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Println(counter)
}
