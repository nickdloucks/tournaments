package entities

import "fmt"

type GameStatus uint8

const (
	Upcoming GameStatus = iota
	InProgress
	Final
)
const UnknownGameStatus = "unkown game status"

type GameResult struct {
	Winner     TournamentParticipant `json:"winner"`
	Loser      TournamentParticipant `json:"loser"`
	WinScore   uint8                 `json:"win_score"`
	LoseScore  uint8                 `json:"lose_score"`
	GameStatus GameStatus            `json:"game_status"`
}

// Displays a string representation of the game state. Uses a value-receiver as it is intended to be used for read-only.
func (gs GameStatus) String() string {
	switch gs {
	case Upcoming:
		return "upcoming"
	case InProgress:
		return "in-progress"
	case Final:
		return "final"
	default:
		return UnknownGameStatus
	}
}

func (gr *GameResult) UpdateGameStatus(gs GameStatus) error {
	if gs.String() == UnknownGameStatus {
		return fmt.Errorf("%v", UnknownGameStatus)
	}
	gr.GameStatus = gs
	return nil
}

func (gr *GameResult) FinalizeGame() {
	gr.UpdateGameStatus(Final)

}