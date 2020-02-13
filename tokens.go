package main

// type stateFunc func(m *Memory)
type stateFunc func(*lexer) stateFunc

type Token struct {
	Raw  string
	Jump int // Index of corresponding jump
}

func (t Token) String() string {
	return t.Raw
}

var (
	RShiftToken = Token{Raw: ">"}
	LShiftToken = Token{Raw: "<"}
	IncToken    = Token{Raw: "+"}
	DecToken    = Token{Raw: "-"}
	GetToken    = Token{Raw: "."}
	SetToken    = Token{Raw: ","}
	LJumpToken  = Token{Raw: "["}
	RJumpToken  = Token{Raw: "]"}
)
