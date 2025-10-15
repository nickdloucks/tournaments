package main

import(
	"fmt"
	"math"
)

type TournamentParticipant struct {
	// A TournamentParticipant is an entity representing a participant (such as a team or player) in a tournament match
	Name: 		string	`json:"name"`,
	Seed:		uint8	`json:"seed"`, // used to determine order of matchups and opponents especially in a bracket-type tournament
	Id:			string	`json:"id,omitempty"`,
	Abbrv:		string	`json:"abbrv,omitempty"`,  // should be 2-4 characters, all-caps
	LogoPath:	string	`json:"logo_path"`
}

type SeriesType uint8

const (
	Single		 	SeriesType = 1,
	BestOfThree		SeriesType = 3,
	BestOfFive 		SeriesType = 5,
	BestOfSeven 	SeriesType = 7,
)

type MarginForVictory uint8

type GameStatus string

const (
	Upcoming	GameStatus = "upcoming",
	InProgress	GameResult = "in-progress",
	Final		GameResult = "final",
)

type GameResult struct {
	Winner:		TournamentParticipant		`json:"winner"`,
	Loser:		TournamentParticipant		`json:"loser"`,
	WinScore:	uint8						`json:"win_score"`,
	LoseScore:	uint8						`json:"lose_score"`,
	GameStatus:	GameStatus					`json:"game_status"`,
}

type MatchSet struct {
	Games:		[]GameResult		`json:"games"`,
}

type Match struct {
	MatchType:				SeriesType = 1			`json:"match_type" default:"1"`, // must be MFV (0) if overriding the Match's MarginForVictory attribute
	SetType:				SeriesType = 1			`json:"set_type" default:"1"`,
	MatchMFV:				MarginForVictory = 1	`json:"match_margin_for_victory" default:"1"`,
	SetMFV:					MarginForVictory = 1	`json:"set_margin_for_victory" default:"1"`,
	Sets:					[]MatchSet				`json:"sets"`,
	HigerSeedParticipant:	TournamentParticipant	`json:"higher_seed"`, // higher seed meens a "better" participant and thus a lower number
	LowerSeedParticipant:	TournamentParticipant	`json:"lower_seed"`, // lower seed means a "worse" participant and this a higher number
}

type TournamentBracket struct {
	// A tree-like data structure containing Match objects
}

func (m *Match) CalcWinThreshold() (uint8, error) {
	if m.MatchType == 0 {
		return 0, fmt.Errorf("unsupported match type")
	}
	return uint8(math.Ceil(m.MatchType / 2)), nil
}

func (s *MatchSet) CalcWinThreshold() (uint8, error) {
	if s.SetType == 0 {
		return 0, fmt.Errorf("unsupported set type")
	}
	return uint8(math.Ceil(s.SetType / 2)), nil
}