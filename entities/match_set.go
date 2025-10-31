package entities

import (
	"fmt"
	"math"
)

type MatchSetEvent struct {
	Id             string              `json:"id"` // uuid V7
	Games          map[uint8]GameEvent `json:"games"`
	SetType        SeriesType          `json:"set_type" default:"1"`
	ParentMatch    *MatchEvent
	HomeOrFav      *TournamentParticipant
	AwayOrUnderdog *TournamentParticipant
	CompetitionEventResult
}

func (s *MatchSetEvent) CalcWinThreshold() (uint8, error) {
	if s.SetType == 0 {
		return 0, fmt.Errorf("unsupported set type")
	}
	middle := float64(s.SetType / 2)
	return uint8(math.Ceil(middle)), nil
}
