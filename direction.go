package main

//Direction direction in which the cells of a ship is layed out in the game board
type Direction int

const (
	// Down ship occupies cells going down
	Down Direction = iota
	// Right ship occupies cells going right
	Right
)

var direction = [...]string{
	"Down",
	"Right",
}

func (d Direction) String() string {
	return direction[d]
}
