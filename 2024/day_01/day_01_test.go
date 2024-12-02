package day_01

import "testing"

func TestDistanceCalculation(t *testing.T) {
	got := Calc("test.input", distance)
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSimilarityScoreCalculation(t *testing.T) {
	got := Calc("test.input", similarityScore)
	want := 31

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
