package main

import "fmt"

func main() {
	demoWithCancel()
	fmt.Println()

	demoWithTimeout()
	fmt.Println()

	demoWithDeadline()
	fmt.Println()

	demoPropagation()
	fmt.Println()

	demoGracefulShutdown()

	// realProductionExample() // it wait for ctrl+c
}
