package entities

import (
	"fmt"
	"math"
)

type Match struct {
	MatchType      SeriesType            `json:"match_type" default:"1"` // must be MFV (0) if overriding the Match's MarginForVictory attribute
	MFVInMatch     MarginForVictory      `json:"margin_for_victory_in_match" default:"1"`
	MFVInSet       MarginForVictory      `json:"margin_for_victory_in_set" default:"1"`
	Sets           map[uint8]MatchSet    `json:"sets"`
	HomeOrFav      TournamentParticipant `json:"home_or_favored"`  // higher seed meens a "better" participant and thus a lower number
	AwayOrUnderdog TournamentParticipant `json:"away_or_underdog"` // lower seed means a "worse" participant and this a higher number
}

func (m *Match) CalcWinThreshold() (uint8, error) {
	if m.MatchType == 0 {
		return 0, fmt.Errorf("unsupported match type")
	}
	middle := float64(m.MatchType / 2)
	return uint8(math.Ceil(middle)), nil
}
