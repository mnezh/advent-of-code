package day_01

import "testing"

func DayOnePartOne(t *testing.T) {
	got := Calc("day_01/test.input")
	want := 11

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
