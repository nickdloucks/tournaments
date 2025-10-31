package entities

// The status of a competition unit (game, set, or match)
type CompUnitStatus uint8

const (
	TBA CompUnitStatus = iota // not yet scheduled
	Upcoming
	InProgress
	Final
)
const UnknownCompUnitStatus = "unkown status"

// Displays a string representation of the game state.
// Uses a value-receiver as it is intended to be used for read-only.
func (c CompUnitStatus) String() string {
	switch c {
	case TBA:
		return "TBA"
	case Upcoming:
		return "upcoming"
	case InProgress:
		return "in-progress"
	case Final:
		return "final"
	default:
		return UnknownCompUnitStatus
	}
}

// A SeriesType is a
type SeriesType uint8

const (
	Single      SeriesType = 1
	BestOfThree SeriesType = 3
	BestOfFive  SeriesType = 5
	BestOfSeven SeriesType = 7
)
const UnknownSeriesType SeriesType = 0

type MarginForVictory uint8

// The participant and their final score in a competition unit (game, set, or match)
type ParticipantOutcome struct {
	Participant TournamentParticipant `json:"participant"`
	Score       uint8                 `json:"score"` // Score can be points scored in a game, games won in a set, or sets won in a match
}

// Maps the two participants to their score in a competition unit (game, set, or match)
type OutcomePair struct {
	HomeOrFavOutcome     ParticipantOutcome `json:"home_or_favorite_outcome"`
	AwayOrUndedogOutcome ParticipantOutcome `json:"away_or_underdog_outcome"`
}

// Generic type to represent the state of any competition unit (game, set, or match)
type CompetitionEventResult struct {
	Winner         ParticipantOutcome `json:"winner"`
	Loser          ParticipantOutcome `json:"loser"`
	CompUnitStatus CompUnitStatus     `json:"competition_unit_status"`
	TieAllowed     bool               `json:"tie_allowed"`
}
