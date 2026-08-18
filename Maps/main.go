package main

import "fmt"

func main() {
	demoNilMap()
	fmt.Println()

	demoCommaOk()
	fmt.Println()

	demoDeleteAndLen()
	fmt.Println()

	demoNestedMap()
	fmt.Println()

	demoSet()
	fmt.Println()

	demoMapIsReferenceType()
	fmt.Println()

	demoMapComparison()
	fmt.Println()

	fmt.Println("--- inventory app ---")
	inv := newInventory()
	inv.addOrUpdate("apple", 50, 0.5)
	inv.addOrUpdate("banana", 30, 0.3)
	inv.addOrUpdate("apple", 20, 0.5) // adds onto the existing quantity, doesn't overwrite

	if item, ok := inv.get("apple"); ok {
		fmt.Printf("apple: quantity=%d price=%.2f\n", item.Quantity, item.Price)
	}

	inv.remove("banana")
	inv.addOrUpdate("mango", 2, 1.5)

	fmt.Println("low stock (<5):")
	for _, item := range inv.lowStock(5) {
		fmt.Printf("  %s: %d\n", item.Name, item.Quantity)
	}

	fmt.Println("full inventory (sorted):")
	for _, item := range inv.listSorted() {
		fmt.Printf("  %s: qty=%d price=%.2f\n", item.Name, item.Quantity, item.Price)
	}

	fmt.Printf("total inventory value: %.2f\n", inv.totalValue())
}
