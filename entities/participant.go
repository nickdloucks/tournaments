package entities

// A TournamentParticipant is an entity representing a participant (such as a team or player) in a tournament match
type TournamentParticipant struct {
	Name string `json:"name"`
	Seed uint8  `json:"seed,omitempty"` // used to determine order of matchups and opponents especially in a bracket-type tournament
	Id   string `json:"id,omitempty"`
}

// supports BYE or TBA
type GenericParticipantName string

const (
	ByeParticipantName = "BYE"
	TbaParticipantName = "TBA"
)

func NewGenericParticipant(name GenericParticipantName, seed ...uint8) TournamentParticipant {
	if len(seed) > 0 {
		return TournamentParticipant{
			Name: string(name),
			Seed: seed[0],
			// Id: NewUuidV7().String(),
		}
	} else {
		return TournamentParticipant{
			Name: string(name),
			// Id: NewUuidV7().String(),
		}
	}
}

// func NewByeParticipant(seed ...uint8) TournamentParticipant {
// 	if len(seed) > 0 {
// 		return TournamentParticipant{
// 			Name: ByeParticipantName,
// 			Seed: seed[0],
// 			// Id: NewUuidV7().String(),
// 		}
// 	} else {
// 		return TournamentParticipant{
// 			Name: ByeParticipantName,
// 			// Id: NewUuidV7().String(),
// 		}
// 	}
// }

// func NewTbaParticipant()
