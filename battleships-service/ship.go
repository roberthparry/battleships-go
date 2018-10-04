package service

//Ship base class of ships
type Ship struct {
	Length int
	Hits   int
}

//NewDestroyer create instance of a destroyer
func NewDestroyer() *Ship {
	return &Ship{Length: 4}
}

//NewBattleship create instance of a destroyer
func NewBattleship() *Ship {
	return &Ship{Length: 5}
}

//IsDestroyed returns true if the ship is destroyed
func (ship *Ship) IsDestroyed() bool {
	return ship.Hits == ship.Length
}
