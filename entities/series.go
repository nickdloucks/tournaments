package entities

type SeriesType uint8

const (
	Single      SeriesType = 1
	BestOfThree SeriesType = 3
	BestOfFive  SeriesType = 5
	BestOfSeven SeriesType = 7
)

type MarginForVictory uint8