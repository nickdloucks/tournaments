package entities

import (
	"testing"
)

func TestUpdateGameStatus(t *testing.T) {
	t.Run("using known GameResult type", func(t *testing.T) {
		testGR := GameResult{
			GameStatus: Upcoming,
		}
		testGR.UpdateGameStatus(Final)
		got := testGR.GameStatus
		want := Final

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("using unkown GameResult type", func(t *testing.T) {
		testGR := GameResult{
			GameStatus: InProgress,
		}
		got := testGR.UpdateGameStatus(255)
		want := UnknownGameStatus

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

// func TestFinalizeGame(t *testing.T) {


// 	// if got != want {
// 	// 	t.Errorf("got %q want %q", got, want)
// 	// }
// }