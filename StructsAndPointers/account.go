package main

import (
	"errors"
	"fmt"
)

type account struct {
	owner   *person
	balance float64
}

// value receiver: only changes the LOCAL copy, caller's balance is untouched
func (a account) deposit(amount float64) {
	a.balance += amount
}

// pointer receiver: reassigns through the pointer, caller's balance actually changes
func (a *account) depositPointer(amount float64) {
	a.balance += amount
}

func (a *account) withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New(a.owner.firstName + " has insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a account) printAccount() {
	fmt.Printf("owner=%s balance=%.2f\n", a.owner.firstName, a.balance)
}
