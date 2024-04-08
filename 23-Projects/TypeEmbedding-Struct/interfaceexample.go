package main

import (
	"fmt"
	"time"
)

func main() {

}

type TeamMember interface{}

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Println(e.FirstName)
}

func (e Employee) PrintDetails() {
	fmt.Println(e.FirstName, e.LastName)
}

func (d Developer) PrintDetails() {
	d.Employee.PrintDetails()
}
