package entities

import (
	"fmt"
	"math"
)

type MatchSeriesEvent struct {
	Id                    string                    `json:"id"`                     // uuid v7
	MatchType             SeriesType                `json:"match_type" default:"1"` // must be MFV (0) if overriding the Match's MarginForVictory attribute
	MFVInMatch            MarginForVictory          `json:"margin_for_victory_in_match" default:"1"`
	Sets                  map[uint8]*SetSeriesEvent `json:"sets"`
	HomeOrFav             TournamentParticipant     `json:"home_or_favored"`  // higher seed meens a "better" participant and thus a lower number
	AwayOrUnderdog        TournamentParticipant     `json:"away_or_underdog"` // lower seed means a "worse" participant and this a higher number
	HomeOrFavSetsWon      uint8                     `json:"home_or_fav_sets_won"`
	AwayOrUnderdogSetsWon uint8                     `json:"away_or_underdog_sets_won"`
	CompetitionEventResult
	CompetitionSeries
}

func (m *MatchSeriesEvent) CalcWinThreshold() (uint8, error) {
	if m.MatchType == 0 {
		return 0, fmt.Errorf("unsupported match type")
	}
	middle := float64(m.MatchType / 2)
	return uint8(math.Ceil(middle)), nil
}

func (m *MatchSeriesEvent) IncrementWinCount(p TournamentParticipant) error {
	return incrementSeriesWinCount[MatchSeriesEvent](p, m)
}

// TO-DO...
func (s *MatchSeriesEvent) DeclareSeriesWinner(TournamentParticipant) error {
	// TO-DO...
	return nil
}

// TO-DO...
func (s *MatchSeriesEvent) InviteParticipant(newParticipant TournamentParticipant, isHomeOrFav bool) {

}