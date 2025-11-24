package entities

import (
	"errors"
)

type CompetitionSeries interface {
	CalcWinThreshold()
	IncrementWinCount(TournamentParticipant) error
	DeclareSeriesWinner(TournamentParticipant) error
	InviteParticipant(newParticipant TournamentParticipant, isHomeOrFav bool)
}

func incrementSeriesWinCount[S MatchSeriesEvent | SetSeriesEvent](p TournamentParticipant, series *S) error {
	if series == nil {
		return errors.New("cannot increment win count on a nil series")
	}
	switch v := any(series).(type) {
	case *MatchSeriesEvent:
		if v.HomeOrFav.Id == p.Id {
			v.HomeOrFavSetsWon += 1
			return nil
		} else if v.AwayOrUnderdog.Id == p.Id {
			v.AwayOrUnderdogSetsWon += 1
			return nil
		} else {
			return errors.New("match participant not found")
		}
	case *SetSeriesEvent:
		if v.HomeOrFav.Id == p.Id {
			v.HomeOrFavGamesWon += 1
			return nil
		} else if v.AwayOrUnderdog.Id == p.Id {
			v.AwayOrUnderdogGamesWon += 1
			return nil
		} else {
			return errors.New("set participant not found")
		}
	default:
		return errors.New("bad series type, cannot increment win total")
	}
}

func declareSeriesWinner[S MatchSeriesEvent | SetSeriesEvent](s *S) error {
	return nil
}