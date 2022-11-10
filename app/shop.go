package app

import (
	"fmt"
)

type Account struct {
	Name, Surname string
}

type Employee struct {
	Accounts Account
	Credits  float64
}

var credits float64

func (e *Employee) String() string {
	return fmt.Sprintf("%s has %v credits\n", e.Accounts, e.Credits)
}

func (c *Account) ChangeName() (string, string) {
	var newName, newSurname string
	fmt.Println("New name")
	fmt.Scanln(&newName)
	fmt.Println("New surname")
	fmt.Scanln(&newSurname)
	c.Name = newName
	c.Surname = newSurname
	return c.Name, c.Surname
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func (e *Employee) AddCredits() float64 {
	fmt.Scanln(&credits)
	e.Credits += credits
	return e.Credits
}

func (e *Employee) RemoveCredits() float64 {
	fmt.Scanln(&credits)
	e.Credits -= credits
	if e.Credits == 0 {
		fmt.Println("0 actual credits, add them now ;)")
		defer e.AddCredits()
	}
	return e.Credits
}
