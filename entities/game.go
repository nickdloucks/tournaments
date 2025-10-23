package entities

import (
	"errors"
	"fmt"
)

type GameStatus uint8

const (
	Upcoming GameStatus = iota
	InProgress
	Final
)
const UnknownGameStatus = "unkown game status"

// The participant and their final score in a game, set, or match
type ParticipantOutcome struct {
	Participant TournamentParticipant `json:"participant"`
	Score       uint8                 `json:"score"`
}

type OutcomePair struct {
	FavoriteOutcome ParticipantOutcome `json:"favorite_outcome"`
	UndedogOutcome  ParticipantOutcome `json:"underdog_outcome"`
}

type GameResult struct {
	Winner     ParticipantOutcome `json:"winner"`
	Loser      ParticipantOutcome `json:"loser"`
	GameStatus GameStatus         `json:"game_status"`
	TieAllowed bool               `json:"tie_allowed"`
}

// Displays a string representation of the game state. 
// Uses a value-receiver as it is intended to be used for read-only.
func (gs GameStatus) String() string {
	switch gs {
	case Upcoming:
		return "upcoming"
	case InProgress:
		return "in-progress"
	case Final:
		return "final"
	default:
		return UnknownGameStatus
	}
}

// 
func (gr *GameResult) SetGameStatus(gs GameStatus) error {
	if gs.String() == UnknownGameStatus {
		return fmt.Errorf("%v", UnknownGameStatus)
	}
	gr.GameStatus = gs
	return nil
}

type SimpleError string

var (
	ErrImpermissibleTie = errors.New("impermissible tie")
)

func (e SimpleError) Error() string {
	return string(e)
}


// Updates the game status and declares the winner and loser.
// If the final score is tied AND a tie is permitted for this game, the "winner" will be the participant with the better seed.
func (gr *GameResult) FinalizeGame(resultPair OutcomePair) error {
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
	gr.SetGameStatus(Final)
	return nil
}
