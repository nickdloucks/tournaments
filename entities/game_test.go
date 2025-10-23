package entities

import (
	"testing"
	"reflect"
	// "github.com/nickdloucks/tournaments/tournamenterrors"
)

func TestSetCompUnitStatus(t *testing.T) {
	t.Run("using known GameResult type", func(t *testing.T) {
		testGR := GameResult{
			CompetitionUnitResult: CompetitionUnitResult{
				CompUnitStatus: Upcoming,
			},
		}
		testGR.SetCompUnitStatus(Final)
		got := testGR.CompUnitStatus
		want := Final

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("using unkown GameResult type", func(t *testing.T) {
		testGR := GameResult{
			CompetitionUnitResult: CompetitionUnitResult{
				CompUnitStatus: InProgress,
			},
		}
		got := testGR.SetCompUnitStatus(255)
		want := UnknownCompUnitStatus

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestFinalizeGame(t *testing.T) {
	t.Run("using tie score", func(t *testing.T) {
		testPOs := [2]ParticipantOutcome{
			{Participant: TournamentParticipant{}, Score: 0}, 
			{Participant: TournamentParticipant{}, Score: 0},
		}
		testPair := OutcomePair{
			FavoriteOutcome: testPOs[0],
			UndedogOutcome: testPOs[1],
		}
		testGR := GameResult{
			CompetitionUnitResult: CompetitionUnitResult{
				TieAllowed: false,
			},
		}
		got := testGR.FinalizeGame(testPair)
		want := ErrImpermissibleTie
		if (got == nil) {
			t.Errorf("impermissible tie score: %v-%v. got %s want %s", testGR.Winner.Score, testGR.Loser.Score, got, want.Error())
		} else if (got != want){
			t.Errorf("impermissible tie score: %v-%v. got %s want %s", testGR.Winner.Score, testGR.Loser.Score, got.Error(), want.Error())
		}
	})
	t.Run("winner selection based on score", func(t *testing.T) {
		testPOs := [2]ParticipantOutcome{
			{Participant: TournamentParticipant{Name: "team-with-seed1-and-score0", Seed: 1}, Score: 0},
			{Participant: TournamentParticipant{Name: "team-with-seed2-and-score2", Seed: 2}, Score: 2},
		}
		testPair := OutcomePair{
			FavoriteOutcome: testPOs[0],
			UndedogOutcome: testPOs[1],
		}
		testGR := GameResult{
			CompetitionUnitResult: CompetitionUnitResult{
				TieAllowed: false,
			},
		}
		testGR.FinalizeGame(testPair)
		got := testGR.Winner
		want := testPOs[1]
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got incorrect winner: got  %v, want %v", got.Participant.Name, want.Participant.Name)
		}
	})
	t.Run("winner selection based on seed", func(t *testing.T) {
		testPOs := [2]ParticipantOutcome{
			{Participant: TournamentParticipant{Name: "team-with-seed1-and-score0", Seed: 1}, Score: 0},
			{Participant: TournamentParticipant{Name: "team-with-seed2-and-score0", Seed: 2}, Score: 0},
		}
		testPair := OutcomePair{
			FavoriteOutcome: testPOs[0],
			UndedogOutcome: testPOs[1],
		}
		testGR := GameResult{
			CompetitionUnitResult: CompetitionUnitResult{
				TieAllowed: true,
			},
		}
		testGR.FinalizeGame(testPair)
		got := testGR.Winner
		want := testPOs[0]
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got incorrect winner: got  %v, want %v", got.Participant.Name, want.Participant.Name)
		}
	})
	t.Run("status of GameResult", func(t *testing.T) {
		testPOs := [2]ParticipantOutcome{
			{Participant: TournamentParticipant{Name: "team-with-seed1-and-score0", Seed: 1}, Score: 0},
			{Participant: TournamentParticipant{Name: "team-with-seed2-and-score0", Seed: 2}, Score: 0},
		}
		testPair := OutcomePair{
			FavoriteOutcome: testPOs[0],
			UndedogOutcome: testPOs[1],
		}
		testGR := GameResult{
			CompetitionUnitResult: CompetitionUnitResult{
				CompUnitStatus: InProgress,
				TieAllowed: true,
			},
		}
		testGR.FinalizeGame(testPair)
		got := testGR.CompUnitStatus
		want := Final
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})


}