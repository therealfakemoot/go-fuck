package main

type Token struct {
	Type Type
	Raw  byte
}

type stateFn func(m *Memory)

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

// Memory represents an interpreter's memory "tape"
type Memory struct {
	// active is the index of the currently "live" memory cell
	active int
	cells  []int
}
