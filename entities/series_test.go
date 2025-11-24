package entities

import (
	"testing"
)

func TestDeclareSeriesWinner(t *testing.T) {
	t1 := NewGenericParticipant("t1", 1)
	t2 := NewGenericParticipant("t2", 2)
	series1 := &SetSeriesEvent{
		MFVInSet: 1,
		SetType: 1,
		HomeOrFav: t1,
		AwayOrUnderdog: t2,
		HomeOrFavGamesWon: 1,
		AwayOrUnderdogGamesWon: 0,
	}
	t.Run("declare winner in a series", func(t *testing.T) {
		err := declareSeriesWinner[SetSeriesEvent](series1)
		if err != nil {
			t.Errorf("%v", err)
		}
		got := series1.CompetitionEventResult.Winner.Participant.Id
		want := t1.Id
		if got != want {
			t.Errorf("incorrect winner declared in series. got %s want %s", got, want)
		}
	})
}