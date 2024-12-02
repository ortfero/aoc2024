package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

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
	lefts := make([]uint, 0, 1001)
	rights := make([]uint, 0, 1001)
	var l, r uint
	for scanner.Scan() {
		line := scanner.Text()
		if n, err := fmt.Sscan(line, &l, &r); err != nil || n != 2 {
			log.Fatal(err)
		}
		lefts = append(lefts, l)
		rights = append(rights, r)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	sort.Slice(lefts, func(i, j int) bool { return lefts[i] < lefts[j] })
	sort.Slice(rights, func(i, j int) bool { return rights[i] < rights[j] })
	var sum uint
	for i := 0; i < len(lefts); i++ {
		l = lefts[i]
		r = rights[i]
		if l >= r {
			sum += l - r
		} else {
			sum += r - l
		}
	}
	fmt.Println(sum)
}
