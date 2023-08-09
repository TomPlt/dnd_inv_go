package main

import (
	"fmt"
	"github.com/TomPlt/dnd_inv_go/inventory" // Replace with the actual path to your inventory package
)

func main() {
	inv := inventory.NewInventory(5)
	fmt.Printf("Current inventory: %v\n", inv.Items)
	item1 := inventory.Item{ID: 1, Name: "Potion", Value: 50, Weight: 0.5, Magical: true, Description: "Heals 50 HP"}
	item2 := inventory.Item{ID: 2, Name: "Sword", Value: 200, Weight: 5.0, Magical: false, Description: "A sharp blade"}
	inv.AddItem(item1, 1)
	fmt.Printf("Current inventory: %v\n", inv.Items)
	err := inv.RemoveItem(1, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	inv.AddItem(item2, 2)
	fmt.Printf("Current inventory: %v\n", inv.Items)
	inv.AddItem(item2, 2)
	fmt.Printf("Current inventory: %v\n", inv.Items)
	// fmt.Printf("Error: %v\n", test)
	// inv.AddItem(item2)
	// fmt.Printf("Current inventory: %v\n", inv.Items)

}
