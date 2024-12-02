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

func isSafeReport(levels []uint64) bool {
	if len(levels) < 2 {
		return false
	}
	increasing := levels[1] >= levels[0]
	for i := 1; i < len(levels); i++ {
		if !isSafeLevel(levels[i-1], levels[i], increasing) {
			return false
		}
	}
	return true
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
	levels := make([]uint64, 0, 8)
	unsafeLevels := make([]uint64, 0, 8)
NextLine:
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		levels = levels[:0]
		for i := 0; i < len(fields); i++ {
			level, err := strconv.ParseUint(fields[i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, level)
		}
		if isSafeReport(levels) {
			counter++
			continue
		}
		for i := 0; i < len(levels); i++ {
			unsafeLevels = unsafeLevels[:0]
			for j := 0; j < len(levels); j++ {
				if i == j {
					continue
				}
				unsafeLevels = append(unsafeLevels, levels[j])
			}
			if isSafeReport(unsafeLevels) {
				counter++
				continue NextLine
			}
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Println(counter)
}
