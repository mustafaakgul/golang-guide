package main

import "fmt"

func main() {
	//x := Account{}
	// x.ID = 1
	// x.AvailableFunds()
	// x.ProcessPayment(20)

	// y := new(CreditAccount)
	// y.ID = 0

	// ------

	x := &CreditAccount{}
	//x.AvailableFunds()
	x.Title = "Ahmet"
	x.setTitle()
	fmt.Print(x.Title)
	//x.AvailableFunds()
}

func (a Account) setTitle() {
	a.Title = "Software Developer"
}

func (a *Account) AvailableFunds() bool {
	a.Title = "Software Developer"
	return true
}

func (a *Account) ProcessPayment(amount float32) bool {
	a.ID = 5
	fmt.Println("Processing payment")
	return true
}

type CreditAccount struct {
	Account
	ASD
}

type Account struct {
	ID    int
	Name  string
	Title string
}

type ASD struct {
	ID int
}
