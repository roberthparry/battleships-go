package main

import "fmt"

//FiringResult the enumeration defining result of a shelling
type FiringResult int

const (
	// Missed shell has missed
	Missed FiringResult = iota
	// Hit shell has hit a ship
	Hit = 1
	// Repeat shell has hit an already shelled cell occupied by a ship
	Repeat = 2
)

var firingResult = [...]string{
	"Missed",
	"Hit",
	"Repeat",
}

func (fr FiringResult) String() string {
	return firingResult[fr]
}

// PrintResponse print a user response to the result.
func (fr FiringResult) PrintResponse() {
	switch fr {
	case Hit:
		fmt.Println("Hit!")
	case Missed:
		fmt.Println("Missed!")
	default:
		fmt.Println("You've already been there!")
	}
}
