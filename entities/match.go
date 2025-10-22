package entities

import (
	"fmt"
	"math"
)

type Match struct {
	MatchType            SeriesType            `json:"match_type" default:"1"` // must be MFV (0) if overriding the Match's MarginForVictory attribute
	MatchMFV             MarginForVictory      `json:"match_margin_for_victory" default:"1"`
	SetMFV               MarginForVictory      `json:"set_margin_for_victory" default:"1"`
	Sets                 map[uint8]MatchSet    `json:"sets"`
	HigerSeedParticipant TournamentParticipant `json:"higher_seed"` // higher seed meens a "better" participant and thus a lower number
	LowerSeedParticipant TournamentParticipant `json:"lower_seed"`  // lower seed means a "worse" participant and this a higher number
}

func (m *Match) CalcWinThreshold() (uint8, error) {
	if m.MatchType == 0 {
		return 0, fmt.Errorf("unsupported match type")
	}
	middle := float64(m.MatchType / 2)
	return uint8(math.Ceil(middle)), nil
}
