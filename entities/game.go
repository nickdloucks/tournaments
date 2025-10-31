package entities

import (
	"fmt"
)

type GameEvent struct {
	CompetitionEventResult
	Id string
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
	if resultPair.FavoriteOutcome.Score == resultPair.UndedogOutcome.Score {
		if ge.TieAllowed {
			ge.Winner = resultPair.FavoriteOutcome
			ge.Loser = resultPair.UndedogOutcome
		} else {
			return ErrImpermissibleTie
		}
	} else if resultPair.FavoriteOutcome.Score > resultPair.UndedogOutcome.Score {
		ge.Winner = resultPair.FavoriteOutcome
		ge.Loser = resultPair.UndedogOutcome
	} else { // underdog has higher score
		ge.Winner = resultPair.UndedogOutcome
		ge.Loser = resultPair.FavoriteOutcome
	}
	ge.SetCompUnitStatus(Final)
	return nil
}
