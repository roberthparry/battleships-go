package main

//CellStatus enumeration type
type CellStatus int

const (
	// Unshelled cell is not shelled
	Unshelled CellStatus = iota
	// Shelled cell is shelled
	Shelled
)

var status = [...]string{
	"Unshelled",
	"Shelled",
}

func (s CellStatus) String() string {
	return status[s]
}
