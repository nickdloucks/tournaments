package entities

import (
	"testing"
)

func TestUpdateGameStatus(t *testing.T) {
	testGR := GameResult{
		GameStatus: Upcoming,
	}
	testGR.UpdateGameStatus(Final)
	got := testGR.GameStatus
	want := Final

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}