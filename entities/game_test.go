package entities

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewGameEvent(t *testing.T) {
	t.Run("init game event with two actual participants", func(t *testing.T) {
		teamA := TournamentParticipant{Name: "team_A"}
		teamB := TournamentParticipant{Name: "team_B"}

		got := NewGameEvent(teamA, teamB, false, false)
		want := GameEvent{
			CompetitionEventResult: CompetitionEventResult{CompUnitStatus: Upcoming},
			HomeOrFav:              teamA,
			AwayOrUnderdog:         teamB,
		}
		if !reflect.DeepEqual(got, want) {
			gotStr := strings.Join([]string{got.HomeOrFav.Name, got.AwayOrUnderdog.Name}, " vs. ")
			gotStr = gotStr + " | status: " + got.CompUnitStatus.String()
			wantStr := strings.Join([]string{want.HomeOrFav.Name, want.AwayOrUnderdog.Name}, " vs. ")
			wantStr = wantStr + " | status: " + want.CompUnitStatus.String()
			t.Errorf("got %q want %q",
				gotStr,
				wantStr,
			)
		}
	})
	t.Run("init game event with only one actual participant", func(t *testing.T) {
		teamA := TournamentParticipant{Name: "team_A"}
		teamB := NewGenericParticipant(ByeParticipantName)

		got := NewGameEvent(teamA, teamB, false, false)
		want := GameEvent{
			CompetitionEventResult: CompetitionEventResult{CompUnitStatus: Upcoming},
			HomeOrFav:              teamA,
			AwayOrUnderdog:         teamB,
		}
		if !reflect.DeepEqual(got, want) {
			gotStr := strings.Join([]string{got.HomeOrFav.Name, got.AwayOrUnderdog.Name}, " vs. ")
			gotStr = gotStr + " | status: " + got.CompUnitStatus.String()
			wantStr := strings.Join([]string{want.HomeOrFav.Name, want.AwayOrUnderdog.Name}, " vs. ")
			wantStr = wantStr + " | status: " + want.CompUnitStatus.String()
			t.Errorf("got %q want %q",
				gotStr,
				wantStr,
			)
		}
	})
	t.Run("init game event with zero actual participant", func(t *testing.T) {
		teamA := NewGenericParticipant(TbaParticipantName)
		teamB := NewGenericParticipant(TbaParticipantName)

		got := NewGameEvent( teamA, teamB, false, true)
		want := GameEvent{
			CompetitionEventResult: CompetitionEventResult{CompUnitStatus: TBA},
			HomeOrFav:              teamA,
			AwayOrUnderdog:         teamB,
		}
		if !reflect.DeepEqual(got, want) {
			gotStr := strings.Join([]string{got.HomeOrFav.Name, got.AwayOrUnderdog.Name}, " vs. ")
			gotStr = gotStr + " | status: " + got.CompUnitStatus.String()
			wantStr := strings.Join([]string{want.HomeOrFav.Name, want.AwayOrUnderdog.Name}, " vs. ")
			wantStr = wantStr + " | status: " + want.CompUnitStatus.String()
			t.Errorf("got %q want %q",
				gotStr,
				wantStr,
			)
		}
	})
}

func TestSetCompUnitStatus(t *testing.T) {
	t.Run("using known GameEvent type", func(t *testing.T) {
		testGE := GameEvent{
			CompetitionEventResult: CompetitionEventResult{
				CompUnitStatus: Upcoming,
			},
		}
		testGE.SetCompUnitStatus(Final)
		got := testGE.CompUnitStatus
		want := Final

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("using unkown GameEvent type", func(t *testing.T) {
		testGE := GameEvent{
			CompetitionEventResult: CompetitionEventResult{
				CompUnitStatus: InProgress,
			},
		}
		got := testGE.SetCompUnitStatus(255)
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
			HomeOrFavOutcome:     testPOs[0],
			AwayOrUndedogOutcome: testPOs[1],
		}
		testGE := GameEvent{
			CompetitionEventResult: CompetitionEventResult{
				TieAllowed: false,
			},
		}
		got := testGE.FinalizeGame(testPair)
		want := ErrImpermissibleTie
		if got == nil {
			t.Errorf("impermissible tie score: %v-%v. got %s want %s", testGE.Winner.Score, testGE.Loser.Score, got, want.Error())
		} else if got != want {
			t.Errorf("impermissible tie score: %v-%v. got %s want %s", testGE.Winner.Score, testGE.Loser.Score, got.Error(), want.Error())
		}
	})
	t.Run("winner selection based on score", func(t *testing.T) {
		testPOs := [2]ParticipantOutcome{
			{Participant: TournamentParticipant{Name: "team-with-seed1-and-score0", Seed: 1}, Score: 0},
			{Participant: TournamentParticipant{Name: "team-with-seed2-and-score2", Seed: 2}, Score: 2},
		}
		testPair := OutcomePair{
			HomeOrFavOutcome:     testPOs[0],
			AwayOrUndedogOutcome: testPOs[1],
		}
		testGE := GameEvent{
			CompetitionEventResult: CompetitionEventResult{
				TieAllowed: false,
			},
		}
		testGE.FinalizeGame(testPair)
		got := testGE.Winner
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
			HomeOrFavOutcome:     testPOs[0],
			AwayOrUndedogOutcome: testPOs[1],
		}
		testGE := GameEvent{
			CompetitionEventResult: CompetitionEventResult{
				TieAllowed: true,
			},
		}
		testGE.FinalizeGame(testPair)
		got := testGE.Winner
		want := testPOs[0]

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got incorrect winner: got  %v, want %v", got.Participant.Name, want.Participant.Name)
		}
	})
	t.Run("status of GameEvent", func(t *testing.T) {
		testPOs := [2]ParticipantOutcome{
			{Participant: TournamentParticipant{Name: "team-with-seed1-and-score0", Seed: 1}, Score: 0},
			{Participant: TournamentParticipant{Name: "team-with-seed2-and-score0", Seed: 2}, Score: 0},
		}
		testPair := OutcomePair{
			HomeOrFavOutcome:     testPOs[0],
			AwayOrUndedogOutcome: testPOs[1],
		}
		testGE := GameEvent{
			CompetitionEventResult: CompetitionEventResult{
				CompUnitStatus: InProgress,
				TieAllowed:     true,
			},
		}
		testGE.FinalizeGame(testPair)
		got := testGE.CompUnitStatus
		want := Final
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
