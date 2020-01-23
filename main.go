package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"
)

type lexer struct {
	input  string
	start  int
	pos    int
	tokens chan Token
	states []stateFunc
}

func lex(input string) (*lexer, *Memory) {
	m := newMem()
	l := &lexer{
		input:  input,
		start:  0,
		pos:    0,
		tokens: make(chan Token),
		states: make([]stateFunc, 0),
	}

	l.run(m)
	return l, m
}

func (l *lexer) run(m *Memory) {
	scanner := bufio.NewScanner(strings.NewReader(l.input))
	scanner.Split(bufio.ScanRunes)

	var jumpStart int

	for scanner.Scan() {
		r := scanner.Text()
		switch r {
		case "[":
			jumpStart = l.pos
		case "]":
			if m.Get() != 0 {
				l.pos = jumpStart
			}
		case ">":
			m.RShift()
		case "<":
			m.LShift()
		case "+":
			m.Inc()
		case "-":
			m.Dec()
		case ".":
			// fmt.Printf("%s", m.Get())
		case ",":
			var v int
			n, err := fmt.Scanf("%d", &v)
			// i'm gonna be a troglodyte and pretend that only ascii numerals exist
			if n != 1 || err != nil {
				return // uhhhhhhhhhhhhhh
				// i hate this but i'm too lazy to figure out a better solution right now
			}
			m.Set(v)
		default:
			continue
		}
		l.pos += 1 // it's important to increment the position counter AFTER we do the work
	}

	close(l.tokens)
}

func main() {
	var c string
	flag.StringVar(&c, "code", "", "brainfuck program to execute")

	flag.Parse()

	_, m := lex(c)
	fmt.Printf("%#v\n", m)
}
