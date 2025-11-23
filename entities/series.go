package entities

type CompetitionSeries interface {
	IncrementWinCount(TournamentParticipant)
	DeclareSeriesWinner(TournamentParticipant)
	InviteParticipant(newParticipant TournamentParticipant, isHomeOrFav bool)
}