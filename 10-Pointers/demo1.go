package pointers

import "fmt"

//demo0 icin gonderilen sayi ve fonk icindeki sayi tamamen bellekte farklı yerlerdeki sayılardır
//fonk icinde bir tane daha acar sayi ve ikisini farklı sonuclandırır print yazarsak farklı
//sayilar oldugunu goruru
//pointer ise fonk icindeki sayi mainden gonderileni degistirsin diyorsak pointer kullanırız
//yani gonderilen sayi bellekteki yerini gondeririz direk o sayiyi işlem yaptırırız onu degistirmiş oluruz
//ramdeki yernden direk onu arttırırz ama normal fonk yeni degisken acar dil ve farklısını 1 arttırır
func Demo0(sayi int) {
	sayi = sayi + 1
	fmt.Println("Demodaki sayı : ", sayi)
}


func Demo1(sayi *int) {//onune yıldız koyarsak degerinden ziyade bellekteki adresini gondeririz
	*sayi = *sayi + 1   //burada adres olan sayi kullanıyoruz deger olan int degil
	fmt.Println("Demodaki sayı : ", *sayi) //bellekteki degeri * ile gsterriz
	fmt.Println("Demodaki sayı : ", sayi) //gnderilen sayinin bellekteki adresi
}


/*
func Test()  {
	sayi:=10
	Demo1(&sayi)
	//numara?? 10
}
*/