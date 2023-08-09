package inventory

import (
	"errors"
	"strconv"
)

type Inventory struct {
	Items       map[int]Item
	Capacity    int
	CurrentSize int
}

func NewInventory(capacity int) *Inventory {
	return &Inventory{
		Items:    make(map[int]Item),
		Capacity: capacity,
	}
}

func (inv *Inventory) AddItem(item Item, quantity int) error {
	if inv.CurrentSize+quantity > inv.Capacity {
		return errors.New("inventory will exceed capacity")
	}
	existingItem, exists := inv.Items[item.ID]
	if exists {
		existingItem.Quantity += quantity
		inv.Items[item.ID] = existingItem
	} else {
		item.Quantity = quantity
		inv.Items[item.ID] = item
	}
	inv.CurrentSize += quantity
	return nil
}

func (inv *Inventory) RemoveItem(itemID int, quantity int) error {
	if _, exists := inv.Items[itemID]; !exists {
		// panic("item does not exist")
		return errors.New("item with id " + strconv.Itoa(itemID) + " does not exist")
	}
	item := inv.Items[itemID]
	item.Quantity -= quantity
	if item.Quantity > 0 {
		inv.Items[itemID] = item
		inv.CurrentSize -= quantity
		return nil
	}
	delete(inv.Items, itemID)
	inv.CurrentSize -= quantity
	return nil
}
