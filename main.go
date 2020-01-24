package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"
)

type lexer struct {
	input  []string
	pos    int
	tokens chan Token
	states []stateFunc
}

func lex(input string) (*lexer, chan Token) {
	l := &lexer{
		input:  make([]string, 0),
		pos:    0,
		tokens: make(chan Token),
		states: make([]stateFunc, 0),
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		l.input = append(l.input, scanner.Text())
	}

	go l.run()
	return l, l.tokens
}

func (l *lexer) next() string {
	if l.pos > len(l.input) {
		return "eof"
	}
	r := l.input[l.pos]
	l.pos++
	return r
}

func (l *lexer) backup() {
	l.pos--
}

func (l *lexer) peek() string {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) emit(t Token) {
	l.tokens <- t
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}

func main() {
	var c string
	flag.StringVar(&c, "code", "", "brainfuck program to execute")

	flag.Parse()

	_, tokens := lex(c)
	for t := range tokens {
		fmt.Printf("%#v\n", t)
	}
	// here is where i plug os.Stdout into some slot in my lexer/interpreter type so the brainfuck program can produce its output
}
