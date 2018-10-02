package service

//Destroyer a destroyer class
type Destroyer struct {
	Ship
}

//NewDestroyer create instance of a destroyer
func NewDestroyer() *Destroyer {
	destroyer := &Destroyer{}
	destroyer.Length = 4
	destroyer.Type = ShipTypeDestroyer
	destroyer.Name = destroyer.Type.String()
	return destroyer
}
