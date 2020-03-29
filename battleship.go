package main

//Battleship a destroyer class
type Battleship struct {
	Ship
}

//NewBattleship create instance of a destroyer
func NewBattleship() *Battleship {
	ship := &Battleship{}
	ship.Length = 5
	ship.Type = BattleShipType
	ship.Name = ship.Type.String()
	return ship
}
