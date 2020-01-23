package main

type Token struct {
	Type Type
	Raw  rune
}

type stateFunc func(m *Memory)

type Type string

const (
	RShift    Type = ">"
	LShift    Type = "<"
	Inc       Type = "+"
	Dec       Type = "-"
	Output    Type = "."
	Input     Type = ","
	LeftJump  Type = "["
	RightJump Type = "]"
)
