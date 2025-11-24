package entities

import (
	"errors"
)

type CompetitionSeries interface {
	CalcWinThreshold()
	IncrementWinCount(TournamentParticipant) error
	DeclareSeriesWinner(TournamentParticipant)
	InviteParticipant(newParticipant TournamentParticipant, isHomeOrFav bool)
}


func incrementSeriesWinCount[S SeriesMatchEvent | SeriesSetEvent](p TournamentParticipant, series *S) error {
	if series == nil {
		return errors.New("cannot increment win count on a nil series")
	}
	switch v := any(series).(type) {
	case *SeriesMatchEvent:
		if v.HomeOrFav.Id == p.Id {
			v.HomeOrFavSetsWon += 1
			return nil
		} else if v.AwayOrUnderdog.Id == p.Id {
			v.AwayOrUnderdogSetsWon += 1
			return nil
		} else {
			return errors.New("match participant not found")
		}
	case *SeriesSetEvent:
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
