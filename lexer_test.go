package main

import (
	"reflect"
	"testing"
)

func Test_Lexer(t *testing.T) {
	t.Run("single ops", func(t *testing.T) {
		cases := []struct {
			in       string
			expected []Token
		}{
			{"+", []Token{IncToken}},
			{"-", []Token{DecToken}},
			{">", []Token{RShiftToken}},
			{"<", []Token{LShiftToken}},
			{"[", []Token{LJumpToken}},
			{"]", []Token{RJumpToken}},
			{",", []Token{SetToken}},
			{".", []Token{GetToken}},
		}

		for _, tt := range cases {

			_, tokens := lex(tt.in)
			var actual []Token
			for token := range tokens {
				actual = append(actual, token)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Logf("Exepcted %s, got %s", tt.expected, actual)
				t.Fail()
			}
		}
	})
	t.Run("short, working code", func(t *testing.T) {
		cases := []struct {
			in       string
			expected []Token
		}{
			{"++++", []Token{IncToken, IncToken, IncToken, IncToken}},
			{">--", []Token{RShiftToken, DecToken, DecToken}},
			{" >--", []Token{RShiftToken, DecToken, DecToken}},
			{"<- -", []Token{LShiftToken, DecToken, DecToken}},
		}

		for _, tt := range cases {

			_, tokens := lex(tt.in)
			var actual []Token
			for token := range tokens {
				actual = append(actual, token)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Logf("Exepcted %s, got %s", tt.expected, actual)
				t.Fail()
			}
		}
	})
}

/*
func Test_StateFuncs(t *testing.T) {
	t.Run("unary ops", func(t *testing.T) {
		cases := []struct {
			in       Storage
			ops      string
			expected int
		}{
			{newMem(), "", 0},
			{newMem(), "+", 1},
		}

		for _, tt := range cases {
			_, m := lex(tt.ops)
			if m.Get() != tt.expected {
				t.Logf("expected %d, received %d", tt.expected, m.Get())
				t.Logf("testing op %s", tt.ops)
				t.Fail()
			}
		}
	})

}
*/
