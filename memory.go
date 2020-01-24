package main

// Storage describes the linear "tape" style storage used by the brainfuck
// language.
// I COULD break this interface down into further ones, but at the moment
// i don't really see a need to
type Storage interface {
	RShift(*lexer)
	LShift(*lexer)

	Inc(*lexer)
	Dec(*lexer)

	Get(*lexer)
	Set(*lexer)
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

/*

func (m *Memory) RShift(m *lexer) {
	m.active++
}

func (m *Memory) RShift(m *lexer) { m.active++ }

func (m *Memory) LShift(m *lexer) { m.active-- }

func (m *Memory) Inc(m *lexer) { m.cells[m.active]++ }
func (m *Memory) Dec(m *lexer) { m.cells[m.active]-- }

func (m *Memory) Get(m *lexer) int { return m.cells[m.active] }
func (m *Memory) Set(m *lexer)     { m.cells[m.active] = n }
*/
