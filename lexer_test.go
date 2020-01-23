package main

import (
	"testing"
)

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
