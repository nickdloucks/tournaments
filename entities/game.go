package entities

import (
	"fmt"
)

type GameEvent struct {
	Id             string
	HomeOrFav      TournamentParticipant
	AwayOrUnderdog TournamentParticipant
	CompetitionEventResult
}

func NewGameEvent(homeOrFavored TournamentParticipant, awayOrUnderdog TournamentParticipant, tieAllowed bool, invitesDeferred bool) GameEvent {
	var initialStatus CompUnitStatus
	if invitesDeferred { // game will be scheduled, but Participants are not yet being invited
		initialStatus = TBA
	} else {
		initialStatus = Upcoming
	}
	return GameEvent{
		CompetitionEventResult: CompetitionEventResult{
			CompUnitStatus: initialStatus,
			TieAllowed:     tieAllowed,
		},
		HomeOrFav:      homeOrFavored,
		AwayOrUnderdog: awayOrUnderdog,
	}
}

func (ge *GameEvent) SetCompUnitStatus(s CompUnitStatus) error {
	if s.String() == UnknownCompUnitStatus {
		return fmt.Errorf("%v", UnknownCompUnitStatus)
	}
	ge.CompUnitStatus = s
	return nil
}

// Updates the game status and declares the winner and loser.
// If the final score is tied AND a tie is permitted for this game, the "winner" will be the participant with the better seed.
func (ge *GameEvent) FinalizeGame(resultPair OutcomePair) error {
	if resultPair.HomeOrFavOutcome.Score == resultPair.AwayOrUndedogOutcome.Score {
		if ge.TieAllowed {
			ge.Winner = resultPair.HomeOrFavOutcome
			ge.Loser = resultPair.AwayOrUndedogOutcome
		} else {
			return ErrImpermissibleTie
		}
	} else if resultPair.HomeOrFavOutcome.Score > resultPair.AwayOrUndedogOutcome.Score {
		ge.Winner = resultPair.HomeOrFavOutcome
		ge.Loser = resultPair.AwayOrUndedogOutcome
	} else { // underdog has higher score
		ge.Winner = resultPair.AwayOrUndedogOutcome
		ge.Loser = resultPair.HomeOrFavOutcome
	}
	ge.SetCompUnitStatus(Final)
	return nil
}
