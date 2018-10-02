package service

//Destroyer a destroyer class
type Destroyer struct {
	Ship
}

//NewDestroyer create instance of a destroyer
func NewDestroyer() *Destroyer {
	ship := &Destroyer{}
	ship.Length = 4
	ship.Type = ShipTypeDestroyer
	ship.Name = ship.Type.String()
	return ship
}
