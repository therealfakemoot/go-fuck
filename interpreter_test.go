package main

import (
	"testing"
)

func Test_InterpeterPure(t *testing.T) {
	// testing functionality that doesn't involve i/o

	t.Run("single cell maths", func(t *testing.T) {

		cases := []struct {
			in       string
			expected int
		}{
			{"+++", 3},
		}

		for _, tt := range cases {
			i, m := New(tt.in)
		}
	})

}
