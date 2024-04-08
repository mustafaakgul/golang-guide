package functions

//multiple return function oop dillerinde yoktur oop de boyle kullanılmıyor farklı seklde bnlar yapılır
//birden fazla donus tiplerinde parantez icine alır dizi gibi sanki

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2) //type casting

	return toplam, cikarim, carpim, bolum
}
