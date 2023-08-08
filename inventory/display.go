package inventory

import "fmt"

func (inv *Inventory) DisplayInventory() {
    fmt.Println("Current Inventory:")
    for _, item := range inv.Items {
        fmt.Printf("ID: %d, Name: %s, Value: %d, Weight: %.2f, Magical: %v\n", item.ID, item.Name, item.Value, item.Weight, item.Magical)
    }
}
