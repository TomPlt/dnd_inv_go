package inventory 

import (
	"errors"
	// "fmt"
)

type Inventory struct {
	Items map[int]Item
	Capacity int
	CurrentSize int
}


func NewInventory(capacity int) *Inventory {
	return &Inventory{
		Items: make(map[int]Item),
		Capacity: capacity,
	}
}

func (inv *Inventory) AddItem(item Item) error {
	if inv.CurrentSize >= inv.Capacity {
		return errors.New("inventory is full")
	}
	inv.Items[item.ID] = item
	inv.CurrentSize++
	return nil
}

func (inv *Inventory) RemoveItem(itemID int) error{
	if _, exists := inv.Items[itemID]; !exists {
		return errors.New("item does not exist")
	}
	delete(inv.Items, itemID)
	inv.CurrentSize--
	return nil
}



