package main

type Token struct {
	Type Type
	Raw  byte
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

// Memory represents an interpreter's memory "tape"
type Memory struct {
	// active is the index of the currently "live" memory cell
	active int
	cells  []int
}

func (m *Memory) RShift() { m.active++ }
func (m *Memory) LShift() { m.active-- }

func (m *Memory) Inc() { m.cells[m.active]++ }
func (m *Memory) Dec() { m.cells[m.active]-- }

func (m *Memory) Get() int  { return m.cells[m.active] }
func (m *Memory) Set(n int) { m.cells[m.active] = n }
