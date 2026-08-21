package main

import "fmt"

// ---- multiple interfaces sharing an identical method signature ----

func demoMultipleInterfaces() {
	fmt.Println("--- multiple interfaces, same method ---")
	eb := englishBot{}

	var b bot = eb
	var g greeter = eb // same value satisfies a second interface with zero extra work
	fmt.Println(b.getGreeting())
	fmt.Println(g.getGreeting())
}

// ---- interface embedding ----

func demoInterfaceEmbedding() {
	fmt.Println("--- interface embedding ---")
	var mb multilingualBot = bilingualBot{}
	fmt.Println(mb.getGreeting())
	fmt.Println(mb.translate("good morning"))
}

// ---- empty interface / any ----

func describe(i any) {
	fmt.Printf("value=%v type=%T\n", i, i)
}

func demoEmptyInterface() {
	fmt.Println("--- empty interface (any) ---")
	describe(42)
	describe("hello")
	describe(englishBot{})
}

// ---- type assertion ----

func demoTypeAssertion() {
	fmt.Println("--- type assertion ---")
	var b bot = spanishBot{}

	if sb, ok := b.(spanishBot); ok {
		fmt.Println("it IS a spanishBot:", sb.getGreeting())
	}

	eb, ok := b.(englishBot) // wrong type — comma-ok form returns ok=false, never panics
	fmt.Println("wrong-type assertion:", eb, ok)

	// b.(englishBot) without ", ok" would PANIC here, since b actually holds a spanishBot
}

// ---- type switch ----

func demoTypeSwitch(b bot) {
	switch v := b.(type) {
	case englishBot:
		fmt.Println("switch matched englishBot:", v.getGreeting())
	case spanishBot:
		fmt.Println("switch matched spanishBot:", v.getGreeting())
	default:
		fmt.Println("switch matched an unknown bot type")
	}
}

// ---- the nil interface gotcha ----

type myError struct{}

func (*myError) Error() string {
	return "boom"
}

func mayFail(fail bool) error {
	var err *myError // a nil pointer of a concrete type
	if fail {
		err = &myError{}
	}
	return err // BUG: this ALWAYS returns a non-nil error interface, even when err is nil!
}

func demoNilInterfaceGotcha() {
	fmt.Println("--- nil interface gotcha ---")
	err := mayFail(false)
	fmt.Println("err == nil:", err == nil) // false! surprising, but correct
	fmt.Println("why: the interface holds a (type=*myError, value=nil) pair — only a truly untyped nil interface equals nil")
}

// ---- pointer receiver required to satisfy an interface ----

type counter struct {
	count int
}

func (c *counter) increment() {
	c.count++
}

type incrementer interface {
	increment()
}

var _ incrementer = &counter{} // compiles; "var _ incrementer = counter{}" would NOT

func demoPointerReceiverInterface() {
	fmt.Println("--- pointer receiver required ---")
	c := counter{}
	var inc incrementer = &c // must pass &c — counter{} alone doesn't satisfy incrementer
	inc.increment()
	inc.increment()
	fmt.Println("count:", c.count)
}

// ---- comparing interface values ----

func demoInterfaceComparison() {
	fmt.Println("--- comparing interfaces ---")
	var b1 bot = englishBot{}
	var b2 bot = englishBot{}
	var b3 bot = spanishBot{}
	fmt.Println("same concrete type & value:", b1 == b2) // true
	fmt.Println("different concrete type:", b1 == b3)    // false
}

type withSlice struct {
	data []int
}

func demoInterfaceComparisonPanic() {
	fmt.Println("--- comparing interfaces holding non-comparable types ---")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	var i1 any = withSlice{data: []int{1, 2}}
	var i2 any = withSlice{data: []int{1, 2}}
	fmt.Println(i1 == i2) // never reached — panics at RUNTIME: comparing uncomparable type withSlice
}
