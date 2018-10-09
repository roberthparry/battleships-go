package service

import "fmt"

const (
	//FiringResultMissed shell has missed
	FiringResultMissed = 0
	//FiringResultHit shell has hit a ship
	FiringResultHit = 1
	//FiringResultRepeat shell has hit an already shelled cell
	FiringResultRepeat = 2
)

//FiringResult the enumeration defining result of a shelling
type FiringResult int

//PrintResponse print a user response to the result.
func (result FiringResult) PrintResponse() {
	switch result {
	case FiringResultHit:
		fmt.Println("Hit!")
	case FiringResultMissed:
		fmt.Println("Missed!")
	default:
		fmt.Println("You've already been there!")
	}
}
