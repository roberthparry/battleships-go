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
func NewDestroyer() *Destroyer {
	ship := &Destroyer{}
	ship.Length = 4
	return ship
}

//NewBattleship create instance of a destroyer
func NewBattleship() *Battleship {
	ship := &Battleship{}
	ship.Length = 5
	return ship
}

//IsDestroyed returns true if the ship is destroyed
func (ship *Ship) IsDestroyed() bool {
	return ship.Hits == ship.Length
}
