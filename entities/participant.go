package entities

// A TournamentParticipant is an entity representing a participant (such as a team or player) in a tournament match
type TournamentParticipant struct {
	Name     string `json:"name"`
	Seed     uint8  `json:"seed"` // used to determine order of matchups and opponents especially in a bracket-type tournament
	Id       string `json:"id,omitempty"`
}