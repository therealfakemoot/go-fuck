package main

// type stateFunc func(m *Memory)
type stateFunc func(*lexer) stateFunc

type Token string

const (
	RShiftToken Token = ">"
	LShiftToken Token = "<"
	IncToken    Token = "+"
	DecToken    Token = "-"
	GetToken    Token = "."
	SetToken    Token = ","
	LJumpToken  Token = "["
	RJumpToken  Token = "]"
)
