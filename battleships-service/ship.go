package service

//Ship base class of ships
type Ship struct {
	Length int
	Name   string
	Type   ShipType
	Hits   int
}

//IsDestroyed returns true if the ship is destroyed
func (ship *Ship) IsDestroyed() bool {
	return ship.Hits == ship.Length
}
