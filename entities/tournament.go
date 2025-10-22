package entities

// configuration options for a tournament bracket
type TournamentCfg struct {
	MatchType SeriesType       `json:"match_type" default:"1"` // must be MFV (0) if overriding the Match's MarginForVictory attribute
	SetType   SeriesType       `json:"set_type" default:"1"`
	MatchMFV  MarginForVictory `json:"match_margin_for_victory" default:"1"`
	SetMFV    MarginForVictory `json:"set_margin_for_victory" default:"1"`
}

type TournamentBracket struct {
	// A tree-like data structure containing Match objects
}

type TournamentRoundRobin struct {
	//
}

// takes a sorted list of participants and generates a tournament bracket tree
func (tb *TournamentBracket) GenerateBracket(participants []TournamentParticipant, cfg *TournamentCfg) error {
	// TO-DO
	return nil
}

// takes a list of participants and generates a collection of round-robin matches
func (trr *TournamentRoundRobin) GenerateMatches(participants []TournamentParticipant, cfg *TournamentCfg) error {
	// TO-DO
	return nil
}