package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	testCases := []struct {
		name         string
		a            int
		b            int
		wantQuestion string
		wantAnswer   int
	}{
		{"1", 1, 1, "1+1", 2},
		{"2", -1, 1, "(-1)+1", 0},
		{"2", 1, -1, "1+(-1)", 0},
	}

	for _, tc := range testCases {
		p := NewProblem(tc.a, tc.b)
		assert.Equal(t, tc.wantQuestion, p.Question())
		assert.True(t, p.IsCorrect(tc.wantAnswer))
	}
}
