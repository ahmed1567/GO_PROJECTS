package main

import "fmt"

// demoPointerBasics shows the two meanings of "*" and the one meaning of "&":
//   - "*" next to a TYPE (var x *float64)      -> declares a pointer type
//   - "*" next to a VARIABLE in an expression   -> dereference (read/write through the pointer)
//   - "&" next to a VARIABLE in an expression   -> address-of (get a pointer to it)
func demoPointerBasics(a *account) {
	var balancePointer *float64 = &a.balance // "*float64" = type, "&a.balance" = address-of
	fmt.Println("balancePointer holds address:", balancePointer)

	fmt.Println("value through *balancePointer (dereference):", *balancePointer)
	*balancePointer += 25 // writes straight into a.balance, no receiver needed at all
	fmt.Println("a.balance after *balancePointer += 25:", a.balance)

	printAccountAddress(a)
}

// printAccountAddress mirrors Maps/main.go: the pointer VALUE (the address of
// the real account) is copied into "a" here, but "a" itself, as a local
// variable, lives at a different address than the caller's "a".
func printAccountAddress(a *account) {
	fmt.Println("address of local parameter a:", &a)
}
