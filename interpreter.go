package main

import (
	"fmt"
	"io"
	// "log"
)

type Interpreter struct {
	tokens []Token
	stack  []int
	r      io.Reader
	w      io.Writer
}

func New(p string, w io.Writer, r io.Reader) *Interpreter {
	_, ts := lex(p)

	tokens := make([]Token, 0)
	for t := range ts {
		tokens = append(tokens, t)
	}

	return &Interpreter{
		tokens: tokens,
		stack:  make([]int, 0),
		w:      w,
		r:      r,
	}
}

func NewInterpreter(ts chan Token, m Storage, w io.Writer, r io.Reader) *Interpreter {
	tokens := make([]Token, 0)
	for t := range ts {
		tokens = append(tokens, t)
	}

	return &Interpreter{
		tokens: tokens,
		stack:  make([]int, 0),
		w:      w,
		r:      r,
	}
}

func (i *Interpreter) Run(m *Memory) {
	for j := 0; j < len(i.tokens); j++ {
		// log.Printf("stepping through token: %s\n", i.tokens[j])
		switch i.tokens[j] {
		case RShiftToken:
			m.RShift()
		case LShiftToken:
			m.LShift()
		case IncToken:
			m.Inc()
		case DecToken:
			m.Dec()
		case GetToken:
			fmt.Fprintf(i.w, "%d\n", m.Get())
		case SetToken:
			m.Set()
		case LJumpToken:
			i.stack = append(i.stack, j)
			if m.LJump() {
				j = i.stack[len(i.stack)-1]
				i.stack = i.stack[:len(i.stack)-1]
			}
		case RJumpToken:
			i.stack = append(i.stack, j)
			if m.RJump() {
				j = i.stack[len(i.stack)-1]
				i.stack = i.stack[:len(i.stack)-1]
			}
		default:
			return
		}
	}

}
