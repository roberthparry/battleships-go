package service

//Ship base class of ships
type Ship struct {
	Length int
	Hits   int
}

//Destroyer a destroyer class
type Destroyer struct {
	Ship
}

//Battleship a destroyer class
type Battleship struct {
	Ship
}

//NewDestroyer create instance of a destroyer
func NewDestroyer() *Ship {
	destroyer := &Destroyer{}
	destroyer.Length = 4
	return &destroyer.Ship
}

//NewBattleship create instance of a destroyer
func NewBattleship() *Ship {
	battleship := &Battleship{}
	battleship.Length = 5
	return &battleship.Ship
}

//IsDestroyed returns true if the ship is destroyed
func (ship *Ship) IsDestroyed() bool {
	return ship.Hits == ship.Length
}
