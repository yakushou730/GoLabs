package chapter3

import "testing"

func Test_Point(t *testing.T) {
	a := NewPoint(1, 1)
	b := NewPoint(1, 1)
	if a != b {
		t.Fatal("a and b were not equal")
	}
}
