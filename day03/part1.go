package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Mul struct {
	x uint
	y uint
}

func parseNum(content []byte, index *int, value *uint) bool {
	i := *index
	var v uint
	for {
		c := content[i]
		if c >= '0' && c <= '9' {
			v *= 10
			v += uint(c - '0')
			i++
			if i == len(content) {
				*index = i
				*value = v
				return true
			}
		} else {
			if i == *index {
				return false
			}
			*index = i
			*value = v
			return true
		}
	}
}

func parseMul(content []byte, index *int, mul *Mul) bool {
	substring := content[*index:]
	i := bytes.Index(substring, []byte("mul("))
	if i == -1 {
		*index += len(substring)
		return false
	}
	*index += i + 4
	if !parseNum(content, index, &mul.x) {
		return false
	}
	if content[*index] != byte(',') {
		return false
	}
	*index++
	if !parseNum(content, index, &mul.y) {
		return false
	}
	if content[*index] != byte(')') {
		return false
	}
	*index++
	return true
}

func main() {
	content, err := os.ReadFile("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum uint
	var mul Mul
	index := 0
	for index < len(content) {
		if parseMul(content, &index, &mul) {
			sum += mul.x * mul.y
		}
	}
	fmt.Println(sum)
}
