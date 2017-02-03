package rtsengine

/*
 Implements the  unit

*/

// Armory is an IUnit that maintains an Armory that produces combat units
type Armory struct {
	Poolable
	HealthAndAttack
	owner IPlayer
}

func (unit *Armory) name() string {
	return "Armory"
}