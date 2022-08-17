package functions

//variadic functions c# params denk gelir
//parametre sayısı belli olmatan args gibi

func ToplaVariadic(sayilar ...int) int { // 3 tane nokta koyarak girenler dizi olarak parameter verilir
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}
	return toplam

}
