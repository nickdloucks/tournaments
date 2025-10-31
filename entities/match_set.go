package entities

import (
	"fmt"
	"math"
)

type MatchSetEvent struct {
	Id                     string               `json:"id"` // uuid V7
	SetType                SeriesType           `json:"set_type" default:"1"`
	MFVInSet               MarginForVictory     `json:"margin_for_victory_in_set" default:"1"`
	Games                  map[uint8]*GameEvent `json:"games"`
	ParentMatch            *MatchEvent
	HomeOrFav              *TournamentParticipant
	AwayOrUnderdog         *TournamentParticipant
	HomeOrFavGamesWon      uint8 `json:"home_or_fav_games_won"`
	AwayOrUnderdogGamesWon uint8 `json:"away_or_underdog_games_won"`
	CompetitionEventResult
}

func (s *MatchSetEvent) CalcWinThreshold() (uint8, error) {
	if s.SetType == 0 {
		return 0, fmt.Errorf("unsupported set type")
	}
	middle := float64(s.SetType / 2)
	return uint8(math.Ceil(middle)), nil
}
