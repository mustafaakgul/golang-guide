package structs

import (
	"fmt"
)

func Demo1() {
	x:= product{"Bilgisayar", 500, "XYZ"}
	fmt.Println(x)
	fmt.Println(product{"Bilgisayar", 500, "ABC"})
	fmt.Println(product{"Bilgisayar", 500, ""})  //bosta olablr ama olmaması olmaz
	fmt.Println(product{name: "Bilgisayar"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}


//structlar oop class ve objeler gibidir
//deger güdümlü calısır primitve tipler gibi referanslar vererek koruruz bnları