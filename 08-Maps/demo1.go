package __Maps

import "fmt"

func Demo1() {
	//key-value
	//diger dillerdeki dictinary yapıları gibidir
	//ilk string key digeri value string
	//make biseyler olusturak fonk
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı : ", len(sozluk))
	fmt.Println("Sözlük : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman sayısı : ", len(sozluk))
	fmt.Println("Sözlük : ", sozluk)

	deger, varMi := sozluk["pencil"]  //bir deger sozlugun icinde varmı yokmu
	//szlk 2 deger dondurebliryor deger ve boolean deger
	fmt.Println(deger)
	fmt.Println("Listede olma durumu : ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}
