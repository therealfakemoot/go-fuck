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
		states: make([]stateFunc),
	}

	go l.run()
	return l, l.tokens
}

func (l *lexer) run() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		r := scanner.Text()
		switch r {
		case "[":
			l.emit(LeftJump)
		case "]":
			l.emit(RightJump)
		case ">":
			l.emit(RShift)
		case "<":
			l.emit(LShift)
		case "+":
			l.emit(Inc)
		case "-":
			l.emit(Dec)
		case ".":
			l.emit(LeftJump)
		case ",":
			l.emit(LeftJump)
		default:
			continue
		}
	}

	for state := l.lexText; state != nil; {
		state = state(l)
		l.states = append(l.states, stateFunc)
	}
	close(l.tokens)
}

func (l *lexer) lexText() stateFn {
}

func (l *lexer) emit(t Type) {
	l.tokens <- Token{t, l.input[l.pos]}
}

func main() {
	fmt.Println("vim-go")
}
