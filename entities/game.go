package entities

import (
	"fmt"
)

type GameEvent struct {
	CompetitionEventResult
	Id string
}



func (gr *GameEvent) SetCompUnitStatus(s CompUnitStatus) error {
	if s.String() == UnknownCompUnitStatus {
		return fmt.Errorf("%v", UnknownCompUnitStatus)
	}
	gr.CompUnitStatus = s
	return nil
}


// Updates the game status and declares the winner and loser.
// If the final score is tied AND a tie is permitted for this game, the "winner" will be the participant with the better seed.
func (gr *GameEvent) FinalizeGame(resultPair OutcomePair) error {
	if resultPair.FavoriteOutcome.Score == resultPair.UndedogOutcome.Score {
		if gr.TieAllowed {
			gr.Winner = resultPair.FavoriteOutcome
			gr.Loser = resultPair.UndedogOutcome
		} else {
			return ErrImpermissibleTie
		}
	} else if resultPair.FavoriteOutcome.Score > resultPair.UndedogOutcome.Score {
		gr.Winner = resultPair.FavoriteOutcome
		gr.Loser = resultPair.UndedogOutcome
	} else { // underdog has higher score
		gr.Winner = resultPair.UndedogOutcome
		gr.Loser = resultPair.FavoriteOutcome
	}
	gr.SetCompUnitStatus(Final)
	return nil
}
