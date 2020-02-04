package main

import (
	"math"
	"os"
	"testing"
)

func Test_InterpeterPure(t *testing.T) {
	// testing functionality that doesn't involve i/o

	t.Run("single cell ops", func(t *testing.T) {

		cases := []struct {
			in       string
			expected uint64
		}{
			{"+++", 3},
			{"----", math.MaxUint64 - 4},
		}

		for _, tt := range cases {
			m := newMem()
			i := New(tt.in, os.Stdin, os.Stdout)
			i.Run(m)

			v := m.Get()
			if v != tt.expected {
				t.Logf("expected %d, got %d", tt.expected, v)
				t.Fail()
			}
		}
	})

}
