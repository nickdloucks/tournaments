package entities

import (
	"fmt"
	"math"
)

type MatchSet struct {
	Games   map[uint8]GameEvent `json:"games"`
	Id      string              `json:"id"`
	SetType SeriesType          `json:"set_type" default:"1"`
}

type MatchSetResult struct {
	CompetitionEventResult
	Games []GameEvent
}

func (s *MatchSet) CalcWinThreshold() (uint8, error) {
	if s.SetType == 0 {
		return 0, fmt.Errorf("unsupported set type")
	}
	middle := float64(s.SetType / 2)
	return uint8(math.Ceil(middle)), nil
}
