package main

import (
	"fmt"
	"onlineshop/app"
)

func main() {
	p := app.Employee{Accounts: app.Account{Name: "Jack", Surname: "Gaal"}, Credits: 100.0}
	fmt.Printf("Employee: \n%s", &p)
	for i := 0; i < 4; {
		fmt.Println("Choose options:")
		fmt.Scanln(&i)
		switch i {
		case 0:
			fmt.Println("Change name...")
			p.Accounts.ChangeName()
		case 1:
			fmt.Println("Add credits...")
			p.AddCredits()
		case 2:
			fmt.Println("Remove credits...")
			p.RemoveCredits()
		case 3:
			fmt.Println("Check credits...")
			p.CheckCredits()
		}
		fmt.Printf("%s", &p)
	}
}
