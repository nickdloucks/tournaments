package main

import(

)

type TournamentParticipant struct {
	// A TournamentParticipant is an entity representing a participant (such as a team or player) in a tournament match
	Name: 		string	`json:"name"`,
	Seed:		uint8	`json:"seed"`, // used to determine order of matchups and opponents especially in a bracket-type tournament
	Id:			string	`json:"id,omitempty"`,
	Abbrv:		string	`json:"abbrv,omitempty"`,  // should be 2-4 characters, all-caps
	LogoPath:	string	`json:"logo_path"`
}

type MatchType uint8

const (
	BestOfOne 	MatchType = 1,
	BestOfThree MatchType = 3,
	BestOfFive 	MatchType = 5,
	BestOfSeven MatchType = 7,
	MFV			MatchType = 0, // use this MatchType when overriding the default MarginForVictory
)

type SetType uint8

type MarginForVictory uint8

type Match struct {
	MatchType:				MatchType = 1			`json:"match_type" default:"1"`, // must be MFV (0) if overriding the Match's MarginForVictory attribute
	SetType:				SetType	= 1				`json:"set_type" default:"1"`,
	MarginForVictory:		MarginForVictory = 1	`json:"margin_for_victory" default:"1"`
	HigerSeedParticipant:	TournamentParticipant	`json:"higher_seed"`,
	LowerSeedParticipant:	TournamentParticipant	`json:"lower_seed"`,
}

type TournamentBracket struct {
	// A tree-like data structure containing GameOrMatch objects
}