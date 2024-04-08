// addition + => 15 + 10 => 15, 10 operand, + operator
// substruction -
// product *
// division /
// remainder %

package main

import "fmt"

func main() {
	//x, y := 15, 10

	//Type and value
	//fmt.Printf("%T, %v\n", x, x)
	//fmt.Printf("%T, %v\n", y, y)
	//fmt.Printf("%T, %v\n", (x + y), (x + y))
	//fmt.Printf("%T, %v\n", (x - y), (x - y))
	//fmt.Printf("%T, %v\n", (x * y), (x * y))
	//fmt.Printf("%T, %v\n", (9.0 / 3), (9.0 / 3))
	//fmt.Printf("%T, %v\n", (x % y), (x % y))

	//z:= float64(5.0/2)

	//z := 5.0 / 2 // (2.5)
	//fmt.Printf("%T, %v\n", z, z)

	// Increment ++, Decrement -- POSTFIX

	x := 10

	fmt.Println(x)
	//Assignement
	x = x - 1

	fmt.Println(x)
	//fmt.Println(x--) bunu yapamassın cnku arttırma azaltma birere statement
	//bunlar bir stırda sadece bir stament blunursa yapamaz

	x--

	fmt.Println(x)
}
