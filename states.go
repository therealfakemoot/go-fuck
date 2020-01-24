package main

import ()

func lexText(l *lexer) stateFunc {
	switch l.next() {
	case "[":
		return LJump
	case "]":
		return RJump
	case ">":
		return RShift
	case "<":
		return LShift
	case "+":
		return Inc
	case "-":
		return Dec
	case ".":
		return Get
	case ",":
		return Set
	default:
		return nil
	}
}

func RJump(l *lexer) stateFunc {
	l.emit(RJumpToken)
	return lexText
}

func LJump(l *lexer) stateFunc {
	l.emit(LJumpToken)
	return lexText
}

func RShift(l *lexer) stateFunc {
	l.emit(RShiftToken)
	return lexText
}

func LShift(l *lexer) stateFunc {
	l.emit(LShiftToken)
	return lexText
}

func Inc(l *lexer) stateFunc {
	l.emit(IncToken)
	return lexText
}

func Dec(l *lexer) stateFunc {
	l.emit(DecToken)
	return lexText
}

func Get(l *lexer) stateFunc {
	l.emit(GetToken)
	return lexText
}

func Set(l *lexer) stateFunc {
	l.emit(SetToken)
	return lexText
}
