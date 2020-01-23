package main

// Storage describes the linear "tape" style storage used by the brainfuck
// language.
// I COULD break this interface down into further ones, but at the moment
// i don't really see a need to
type Storage interface {
	RShift()
	LShift()

	Inc()
	Dec()

	Get() int
	Set(int)
}

func newMem() *Memory {
	var m Memory
	m.cells = make(map[int]int)
	return &m
}

// Memory represents an interpreter's memory "tape"
type Memory struct {
	// active is the index of the currently "live" memory cell
	active int
	cells  map[int]int
}

func (m *Memory) RShift() { m.active++ }

func (m *Memory) LShift() { m.active-- }

func (m *Memory) Inc() { m.cells[m.active]++ }
func (m *Memory) Dec() { m.cells[m.active]-- }

func (m *Memory) Get() int  { return m.cells[m.active] }
func (m *Memory) Set(n int) { m.cells[m.active] = n }
