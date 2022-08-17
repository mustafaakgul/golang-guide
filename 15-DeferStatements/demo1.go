package _5_DeferStatements

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}

func C() {
	fmt.Println("c fonksiyonu çalıştı")
}

func D() {
	fmt.Println("d fonksiyonu çalıştı")
}

func B() {
	defer A()  //defer a tanımlanırsa bir fonksiyon defer olarak yani bu ne olursa olsun B fonksitonu bitince calısacak
	//yani a ve c b nin en sonunda calısacak garanti edecek
	// sırasıda su sekilde ilk defer olanlar calısmayacak deferdeki sıra ise su  fonk bitimndeki suslu paranteze gelcek
	//snra deferleri calıstıracak deferler LIFO en son giren ilk calısır veri yapısı ile calısır ilk c calısaak o yuzden
    defer C()
	defer D()
	// defer
	fmt.Println("b fonksiyonu çalıştı")

	//bu soylede ooablirdi
	/*
	D()
	C()
	A()
	boylede olurdu ama fonk hata olursa sistem patlardı defer hata oluncada calısmasını garanti ediyor
	 */
	
}