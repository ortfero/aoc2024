package main

import (
	"fmt"
	"log"
	"os"
)

type Instruction int

const (
	InstructionNone = iota
	InstructionMul
	InstructionDo
	InstructionDont
)

type Parser struct {
	source []byte
	index  int
}

type Mul struct {
	x uint
	y uint
}

func (p *Parser) isAllParsed() bool {
	return p.index >= len(p.source)
}

func (p *Parser) nextInstruction() Instruction {
	n := len(p.source)
	switch p.source[p.index] {
	case 'd':
		p.index++
		if n-p.index < 1 {
			return InstructionNone
		}
		if p.source[p.index] != byte('o') {
			return InstructionNone
		}
		p.index++
		if n-p.index < 3 {
			return InstructionDo
		}
		if p.source[p.index] == byte('n') &&
			p.source[p.index+1] == byte('\'') &&
			p.source[p.index+2] == byte('t') {
			p.index += 3
			return InstructionDont
		}
		return InstructionDo
	case 'm':
		p.index++
		if n-p.index < 2 {
			return InstructionNone
		}
		if p.source[p.index] == byte('u') &&
			p.source[p.index+1] == byte('l') {
			p.index += 2
			return InstructionMul
		}
		return InstructionNone
	default:
		p.index++
		return InstructionNone
	}
}

func (p *Parser) parseNoArgs() bool {
	if p.source[p.index] != byte('(') {
		return false
	}
	p.index++
	if p.source[p.index] != byte(')') {
		return false
	}
	p.index++
	return true
}

func (p *Parser) parseNum(value *uint) bool {
	i := p.index
	var v uint
	for {
		c := p.source[i]
		if c >= '0' && c <= '9' {
			v *= 10
			v += uint(c - '0')
			i++
			if i == len(p.source) {
				p.index = i
				*value = v
				return true
			}
		} else {
			if i == p.index {
				return false
			}
			p.index = i
			*value = v
			return true
		}
	}
}

func (p *Parser) parseTwoNums(x *uint, y *uint) bool {
	if p.source[p.index] != byte('(') {
		return false
	}
	p.index++
	if !p.parseNum(x) {
		return false
	}
	if p.source[p.index] != byte(',') {
		return false
	}
	p.index++
	if !p.parseNum(y) {
		return false
	}
	if p.source[p.index] != byte(')') {
		return false
	}
	p.index++
	return true
}

func main() {
	content, err := os.ReadFile("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parser := Parser{source: content, index: 0}
	var sum uint
	var mul Mul
	enabled := true
	for !parser.isAllParsed() {
		switch parser.nextInstruction() {
		case InstructionNone:
			continue
		case InstructionDo:
			if parser.parseNoArgs() {
				enabled = true
			}
		case InstructionDont:
			if parser.parseNoArgs() {
				enabled = false
			}
		case InstructionMul:
			if parser.parseTwoNums(&mul.x, &mul.y) && enabled {
				sum += mul.x * mul.y
			}
		}
	}
	fmt.Println(sum)
}
