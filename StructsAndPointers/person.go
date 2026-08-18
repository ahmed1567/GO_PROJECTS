package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
	accounts  []*account
}

type contactInfo struct {
	email string
	phone string
}

func (p person) updateName(newFirstName string, newLastName string) {
	// value receiver: only changes the LOCAL copy
	p.firstName = newFirstName
	p.lastName = newLastName
}

func (p *person) updateNamePointer(newFirstName string, newLastName string) {
	// pointer receiver: changes the real struct
	p.firstName = newFirstName
	p.lastName = newLastName
}

func (p person) printPerson() {
	fmt.Printf("%+v\n", p)
}

// openAccount creates a new account owned by p, links it back to p via
// owner, and appends it to p's own list of accounts.
func (p *person) openAccount(initialBalance float64) *account {
	acc := &account{
		owner:   p,
		balance: initialBalance,
	}
	p.accounts = append(p.accounts, acc)
	return acc
}

func (p person) totalBalance() float64 {
	total := 0.0
	for _, acc := range p.accounts {
		total += acc.balance
	}
	return total
}
