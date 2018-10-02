package service

//Battleship a destroyer class
type Battleship struct {
	Ship
}

//NewBattleship create instance of a destroyer
func NewBattleship() *Battleship {
	battleship := &Battleship{}
	battleship.Length = 5
	battleship.Type = ShipTypeBattleShip
	battleship.Name = battleship.Type.String()
	return battleship
}
