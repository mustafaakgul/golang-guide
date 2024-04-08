package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"

	i := 1  //:= olayı tanımlamadan kullanmamızı sagladı
	//infinite loop if there is no i = i + 5
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1
	}

	/*
	for i := 0; i < ; i++ {

	}
	for i2, i3 := range collection {

	}
	*/
	
	fmt.Println("Bitti")
}
