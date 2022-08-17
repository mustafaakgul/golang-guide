package __Slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)  //slide definition

	fmt.Println(isimler)
	isimler[0] = "Engin"
	isimler[1] = "Derin"
	isimler[2] = "Salih"
	isimler = append(isimler, "Büşra")  //arkada yeni bir dizi olusturur digerini silip aslında
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
