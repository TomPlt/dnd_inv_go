package player

import (
	"github.com/TomPlt/dnd_inv_go/inventory"
)

type Player struct {
	Name      string
	Inventory inventory.Inventory
	Money     int
}

func NewPlayer(name string, capacity int, money int) *Player {
	return &Player{
		Name:      name,
		Inventory: *inventory.NewInventory(capacity),
		Money:     money,
	}
}

