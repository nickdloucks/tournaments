package entities

import (
	"testing"
)

func TestIncrementWinCount(t *testing.T) {
	
	t.Run("should result in the win count being one greater than previous", func(t *testing.T) {		
		prevWinTotal := 0
		homeTeam := NewGenericParticipant("homeTeam", 1)
		match1 := &SeriesMatchEvent{
			HomeOrFav:        homeTeam,
			HomeOrFavSetsWon: uint8(prevWinTotal),
		}
		err := match1.IncrementWinCount(homeTeam)
		if err != nil {
			t.Error(err)
		}
		want := prevWinTotal + 1
		if (match1.HomeOrFavSetsWon - 1) != uint8(prevWinTotal) {
			t.Errorf("competition series win count not incremented properly. got %v expected %v", match1.HomeOrFavSetsWon, want)
		}
	})
}
