package entities

import (
	"fmt"
	"math"
)

type SetSeriesEvent struct {
	// TO-DO: need to handle case where the "home" participant changes from game to game but
	// the perticipant with home-field advantage in the series remains unchanged throughout
	Id                     string               `json:"id"` // uuid V7
	SetType                SeriesType           `json:"set_type" default:"1"`
	MFVInSet               MarginForVictory     `json:"margin_for_victory_in_set" default:"1"`
	Games                  map[uint8]*GameEvent `json:"games"`
	ParentMatch            *MatchSeriesEvent
	HomeOrFav              TournamentParticipant // TO-DO: be sure these are consistent whether they are pointers or not
	AwayOrUnderdog         TournamentParticipant
	HomeOrFavGamesWon      uint8 `json:"home_or_fav_games_won"`
	AwayOrUnderdogGamesWon uint8 `json:"away_or_underdog_games_won"`
	CompetitionEventResult
	CompetitionSeries
}

func (s *SetSeriesEvent) CalcWinThreshold() (uint8, error) {
	if s.SetType == 0 {
		return 0, fmt.Errorf("unsupported set type")
	}
	middle := float64(s.SetType / 2)
	return uint8(math.Ceil(middle)), nil
}

func (s *SetSeriesEvent) IncrementWinCount(p TournamentParticipant) error {
	return incrementSeriesWinCount[SetSeriesEvent](p, s)
}

// TO-DO...
func (s *SetSeriesEvent) DeclareSeriesWinner(TournamentParticipant) error {
	// TO-DO...
	return nil
}

// TO-DO...
func (s *SetSeriesEvent) InviteParticipant(newParticipant TournamentParticipant, isHomeOrFav bool) {

}