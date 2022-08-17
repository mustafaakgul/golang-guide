package _7_Strings

//alias
import (
	"fmt"
	s "strings"  // alias demek string i s yaptık kısayol bulmak
)

//s.methods OR string.methods

//case sensitive
//ascii
func Demo1() {
	isim := "Engin"
	fmt.Println(s.Count(isim, "n"))  //ismn icinde kac tane n var
	fmt.Println(s.Contains(isim, "a"))
	sonuc := s.Index(isim, "a") //gordugu ilk index i dondurur

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))
}
