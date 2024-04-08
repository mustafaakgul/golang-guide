package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int { //parametr and last int is a return type
	var toplam = sayi1 + sayi2
	return toplam
}

func SelamVer(kullaniciAdi string) { //no return type
	fmt.Println("Merhaba ", kullaniciAdi)
}
