package main
//package scope

import "fmt"

//burada packvar paket scopuna sahip bu paketten her yerden ulasırız
/* var packVar = "Package Scope"
var funcVar = "Func(Package) Scope" */

//var := "test" scope harici := izin verilmez

func main() {

	/* var packVar = "Package Scope"
	var funcVar = "Func(Package) Scope" */

	//block degiskeni varolduguu blockda erişleblir bu lbock fonsiyon veya if icinde gercli olblr
	/* 	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	} */

	var name = "Arin"
	name, surname := "Elis", "Software"
	fmt.Println(name, surname)

	//fmt.Println(blockVar) ulasamassn

	/* 	fmt.Println(packVar)
	   	myFunc() */

}

/* func myFunc() {
	fmt.Println(funcVar)
	fmr.Println(packvar)
}
*/