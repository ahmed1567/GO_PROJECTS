package main

import "sort"

// Item is what an inventory map value points to. Storing *Item (not Item)
// means anyone holding a pointer can mutate quantity/price in place.
type Item struct {
	Name     string
	Quantity int
	Price    float64
}

// inventory wraps a map so we can attach methods to it — a plain
// map[string]*Item can't have its own methods.
type inventory map[string]*Item

func newInventory() inventory {
	return make(inventory) // make() is required for a map you intend to write to
}

func (inv inventory) addOrUpdate(name string, quantity int, price float64) {
	if existing, ok := inv[name]; ok {
		existing.Quantity += quantity // mutate through the pointer already stored
		return
	}
	inv[name] = &Item{Name: name, Quantity: quantity, Price: price}
}

func (inv inventory) remove(name string) bool {
	if _, ok := inv[name]; !ok {
		return false
	}
	delete(inv, name)
	return true
}

// get uses the "comma-ok" idiom: the second value tells you whether the key
// was actually present, since a missing key just returns the zero value
// (nil here) with no error at all.
func (inv inventory) get(name string) (*Item, bool) {
	item, ok := inv[name]
	return item, ok
}

func (inv inventory) totalValue() float64 {
	total := 0.0
	for _, item := range inv { // iteration order is randomized every run
		total += float64(item.Quantity) * item.Price
	}
	return total
}

// listSorted returns items sorted by name — maps have no built-in order,
// so "sorted iteration" means: collect the keys, sort THEM, then look each up.
func (inv inventory) listSorted() []*Item {
	names := make([]string, 0, len(inv))
	for name := range inv {
		names = append(names, name)
	}
	sort.Strings(names)

	items := make([]*Item, 0, len(names))
	for _, name := range names {
		items = append(items, inv[name])
	}
	return items
}

func (inv inventory) lowStock(threshold int) []*Item {
	var low []*Item
	for _, item := range inv {
		if item.Quantity < threshold {
			low = append(low, item)
		}
	}
	return low
}
