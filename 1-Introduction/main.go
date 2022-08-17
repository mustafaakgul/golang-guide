package main

import "fmt"

func Topla(sayi1 int, sayi2 int) (int, int, int) { //parametr and last int is a return type
	var toplam = sayi1 + sayi2
	return toplam, toplam+1, toplam+2
}

func SelamVer(kullaniciAdi string) {  //no return type
	fmt.Println("Merhaba ", kullaniciAdi)
}

func main()  {
	//fmt.Println("hello")
	//print println printf
	/*
	x := 5
	fmt.Println(&x)
	//x := 4
	y:=5
	fmt.Println(&y)

	{
	x := 7
	fmt.Println(&x)

	}*/

	x, y := 5, 7
	fmt.Println(Topla(x, y))




}

//compile : go build main.go

