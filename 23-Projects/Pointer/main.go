package main

import "fmt"

func main() {

	x := new(Account)
	x.Title = "Ahmet"
	x.setTitleWithPointer()
	fmt.Print(x.Title)

	//--------------------------

	b := x.setTitle()
	fmt.Print(b.Title)

}

func (a *Account) setTitleWithPointer() {
	a.Title = "Software Developer"
}

func (a Account) setTitle() Account {
	a.Title = "Software Developer"
	return a
}

func (a Account) NameTitle() string {
	return a.Title + a.Name
}

type Account struct {
	ID    int
	Name  string
	Title string
}
