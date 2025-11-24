package entities

type CompetitionSeries interface {
	CalcWinThreshold()
	IncrementWinCount(TournamentParticipant)
	DeclareSeriesWinner(TournamentParticipant)
	InviteParticipant(newParticipant TournamentParticipant, isHomeOrFav bool)
}