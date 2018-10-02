package service

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
