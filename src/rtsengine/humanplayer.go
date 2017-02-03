package rtsengine

import (
	"fmt"
	"image"
)

// HumanPlayer implements the IPlayer interface for a human player
type HumanPlayer struct {
	// Structures common to all players.
	BasePlayer

	// Live TCPWire to communicate with UI
	Wire *TCPWire
}

// NewHumanPlayer constructs a HumanPlayer
func NewHumanPlayer(description string, worldLocation image.Point, width int, height int, pool *Pool, pathing *AStarPathing, world *World) *HumanPlayer {
	player := HumanPlayer{}

	player.description = description
	player.GenerateView(worldLocation, width, height)
	player.ItemPool = pool
	player.Pathing = pathing
	player.OurWorld = world

	// Add mechanics

	return &player
}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
func (player *HumanPlayer) listen(wire *TCPWire) {
	player.Wire = wire
}

func (player *HumanPlayer) isHuman() bool {
	return true
}

func (player *HumanPlayer) isWireAlive() bool {
	// Best guess
	return player.Wire != nil
}

func (player *HumanPlayer) start() error {

	if !player.isWireAlive() {
		return fmt.Errorf("Failed: This player does not have an active wire connection.")
	}

	return nil
}

func (player *HumanPlayer) stop() {

}

/////////////////////////////////////////////////////////////////////////
//                           IPlayer interface                         //
/////////////////////////////////////////////////////////////////////////
