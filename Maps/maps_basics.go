package main

import (
	"fmt"
	"reflect"
)

func demoNilMap() {
	fmt.Println("--- nil map ---")
	var m map[string]int      // zero value of a map is nil, not an empty map
	fmt.Println(m["missing"]) // reading a nil map is safe — returns the zero value, 0
	fmt.Println(m == nil)     // true

	// m["key"] = 1 // would PANIC: assignment to entry in nil map
	fmt.Println("writing to a nil map panics — must use make() or a literal first")
}

func demoCommaOk() {
	fmt.Println("--- comma-ok ---")
	ages := map[string]int{"ahmed": 21}

	age, ok := ages["ahmed"]
	fmt.Println(age, ok) // 21 true

	age, ok = ages["sara"]
	fmt.Println(age, ok) // 0 false — can't tell "missing" from "actually zero" without ok
}

func demoDeleteAndLen() {
	fmt.Println("--- delete and len ---")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("len before:", len(m))
	delete(m, "b")
	delete(m, "not-there") // deleting a missing key is a no-op, never panics
	fmt.Println("len after:", len(m))
}

func demoNestedMap() {
	fmt.Println("--- nested map ---")
	// category -> item name -> quantity
	stock := map[string]map[string]int{
		"fruits":     {"apple": 10, "banana": 5},
		"vegetables": {"carrot": 20},
	}
	fmt.Println(stock["fruits"]["apple"])       // 10
	fmt.Println(stock["vegetables"]["missing"]) // 0 — still safe, just a zero value

	// the inner map must exist before you can write into it directly
	if stock["grains"] == nil {
		stock["grains"] = make(map[string]int)
	}
	stock["grains"]["rice"] = 15
	fmt.Println(stock["grains"])
}

// map[string]struct{} is Go's idiomatic "set" — struct{} takes zero bytes,
// so you only pay for the keys, and membership is checked the comma-ok way.
func demoSet() {
	fmt.Println("--- map as a set ---")
	seen := map[string]struct{}{}
	seen["apple"] = struct{}{}
	seen["banana"] = struct{}{}

	_, exists := seen["apple"]
	fmt.Println("apple seen:", exists) // true
	_, exists = seen["mango"]
	fmt.Println("mango seen:", exists) // false
}

// mutateMap shows maps are reference types like slices: passing one to a
// function does NOT copy the underlying data, so changes are visible to
// the caller without needing a pointer receiver or *map.
func mutateMap(m map[string]int) {
	m["added-inside-function"] = 99
}

func demoMapIsReferenceType() {
	fmt.Println("--- maps as reference types ---")
	m := map[string]int{"x": 1}
	mutateMap(m)
	fmt.Println(m) // includes "added-inside-function" — no pointer needed, unlike a struct
}

// maps can't be compared with == (except to nil) — you need reflect.DeepEqual
// or your own key-by-key comparison.
func demoMapComparison() {
	fmt.Println("--- comparing maps ---")
	a := map[string]int{"x": 1}
	b := map[string]int{"x": 1}
	// fmt.Println(a == b) // compile error: invalid operation (map can only be compared to nil)
	fmt.Println("deep equal:", reflect.DeepEqual(a, b))
}
