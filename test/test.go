package test

import "testing"

func TestAbs(t *testing.T) {
	got := abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -(-a)
}
