package main

import "fmt"

func main() {

	var p person = person{
		firstName: "Ahmed",
		lastName:  "Ali",
		age:       21,
		contact: contactInfo{
			email: "[EMAIL_ADDRESS]",
			phone: "1234567890",
		},
	}
	p.updateName("Osama", "Ali")
	p.printPerson()
	p.updateNamePointer("Mohamed", "Ali")
	p.printPerson()

	// person -> account relation: p owns both accounts below,
	// and each account points back to p through owner.
	savings := p.openAccount(100)
	checking := p.openAccount(50)

	savings.deposit(50)        // value receiver bug: only a local copy changes
	savings.printAccount()     // still balance=100.00
	savings.depositPointer(50) // pointer receiver: the real account changes
	savings.printAccount()     // balance=150.00

	if err := checking.withdraw(200); err != nil {
		fmt.Println("Error:", err) // insufficient funds
	}
	checking.depositPointer(200)
	checking.withdraw(200)
	checking.printAccount() // balance=50.00

	fmt.Printf("%s's total balance across %d accounts: %.2f\n",
		p.firstName, len(p.accounts), p.totalBalance())

	fmt.Println("address of savings in main:", &savings)
	demoPointerBasics(savings) // shows *, &, and the copied-pointer-different-address rules
}
