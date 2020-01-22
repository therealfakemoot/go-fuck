package main

import (
	"bufio"
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

func lex(input string) (*lexer, chan Token) {
	l := &lexer{
		input:  input,
		start:  0,
		pos:    0,
		tokens: make(chan Token),
		states: make([]stateFunc, 0),
	}

	go l.run()
	return l, l.tokens
}

func (l *lexer) run(m *Memory) {
	scanner := bufio.NewScanner(strings.NewReader(l.input))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		r := scanner.Text()
		switch r {
		case "[":
			// okay, the jumps are gonna be weird
		case "]":
			// l.emit(RightJump)
		case ">":
			m.RShift()
		case "<":
			m.LShift()
		case "+":
			m.Inc()
		case "-":
			m.Dec()
		case ".":
			var v int
			n, err := fmt.Scanf("%d", &v)
			if len(n) > 1 || err != nil {
				return // uhhhhhhhhhhhhhh
			}
			m.Set(v)
			// fmt.Printf("%s", m.Get())
		case ",":
			// l.emit(LeftJump)
		default:
			continue
		}
	}

	close(l.tokens)
}

func main() {
	fmt.Println("vim-go")
}
