package main

import (
	"testing"
)

func newMem() *Memory {
	var m Memory
	return &m
}

func Test_StateFuncs(t *testing.T) {
	t.Run("unary ops", func(t *testing.T) {
		tt := []struct {
			in       *Storage
			ops      []stateFunc
			expected int
		}{
			{newMem(), []stateFunc{RShift}, 0},
		}
	})
}
