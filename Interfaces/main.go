package main

import "fmt"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
	fmt.Println()

	demoMultipleInterfaces()
	fmt.Println()

	demoInterfaceEmbedding()
	fmt.Println()

	demoEmptyInterface()
	fmt.Println()

	demoTypeAssertion()
	fmt.Println()

	demoTypeSwitch(eb)
	demoTypeSwitch(sb)
	fmt.Println()

	demoNilInterfaceGotcha()
	fmt.Println()

	demoPointerReceiverInterface()
	fmt.Println()

	demoInterfaceComparison()
	fmt.Println()

	demoInterfaceComparisonPanic()
}
