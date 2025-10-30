package entities

import (
	"errors"
	"fmt"
)

// The status of a competition unit (game, set, or match)
type CompUnitStatus uint8

const (
	TBA	CompUnitStatus = iota // not yet scheduled
	Upcoming 
	InProgress
	Final
)
const UnknownCompUnitStatus = "unkown status"

// The participant and their final score in a competition unit (game, set, or match)
type ParticipantOutcome struct {
	Participant TournamentParticipant `json:"participant"`
	Score       uint8                 `json:"score"` // Score can be points scored in a game, games won in a set, or sets won in a match
}

// Maps the two participants to their score in a competition unit (game, set, or match)
type OutcomePair struct {
	FavoriteOutcome ParticipantOutcome `json:"favorite_outcome"`
	UndedogOutcome  ParticipantOutcome `json:"underdog_outcome"`
}

type GameResult struct {
	CompetitionUnitResult
	Id string
}

// Generic type to represent the state of any competition unit (game, set, or matche)
type CompetitionUnitResult struct {
	Winner         ParticipantOutcome `json:"winner"`
	Loser          ParticipantOutcome `json:"loser"`
	CompUnitStatus CompUnitStatus     `json:"game_status"`
	TieAllowed     bool               `json:"tie_allowed"`
}

// Displays a string representation of the game state.
// Uses a value-receiver as it is intended to be used for read-only.
func (c CompUnitStatus) String() string {
	switch c {
	case TBA:
		return "TBA"
	case Upcoming:
		return "upcoming"
	case InProgress:
		return "in-progress"
	case Final:
		return "final"
	default:
		return UnknownCompUnitStatus
	}
}

func (gr *GameResult) SetCompUnitStatus(s CompUnitStatus) error {
	if s.String() == UnknownCompUnitStatus {
		return fmt.Errorf("%v", UnknownCompUnitStatus)
	}
	gr.CompUnitStatus = s
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
	gr.SetCompUnitStatus(Final)
	return nil
}
