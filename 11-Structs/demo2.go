package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

//class icine fonk yapılıyor nrmalde ancak class icin oldugunu c customer ile verilebliryor
func (c customer) save() {      //normal class icinde verilirdi ama burada boyle tanımlanabliyor
	fmt.Println("Eklendi : ", c.firstName)
}

func (c customer) update() {
	fmt.Println("Güncellendi : ", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Engin", lastName: "Demiroğ", age: 35}
	c.save()
	c.update()
}
