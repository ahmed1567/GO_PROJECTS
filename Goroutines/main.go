package main

import "fmt"

func main() {
	demoFireAndForget()
	fmt.Println()

	demoWaitGroup()
	fmt.Println()

	demoUnbufferedChannel()
	fmt.Println()

	demoBufferedChannel()
	fmt.Println()

	demoChannelDirections()
	fmt.Println()

	demoClosingChannel()
	fmt.Println()

	demoSelect()
	fmt.Println()

	demoSelectDefault()
	fmt.Println()

	demoSelectTimeout()
	fmt.Println()

	demoRaceCondition()
	fmt.Println()

	demoMutexFix()
	fmt.Println()

	demoWorkerPool()
	fmt.Println()

	fmt.Println("--- continuous link monitor (real example, in linkchecker.go) ---")
	// fmt.Println("not run here since it never returns — call monitorLinksForever() yourself to try it")
}
