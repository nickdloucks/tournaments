package entities

import "errors"

type SimpleError string

var (
	ErrImpermissibleTie = errors.New("impermissible tie")
)

func (e SimpleError) Error() string {
	return string(e)
}