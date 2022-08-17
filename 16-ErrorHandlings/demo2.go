package _6_ErrorHandlings

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {  //donus tipim string sonuc yani ve error artık byle kullancaz

	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")   //hata doncezya bunu yeni nesne ureterek ypılır
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı giriniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı giriniz", nil
	}
	//hic bi return calısmassa tek olasılık bilmek kalıyor alltaki calısır herhngi biri bilinirse ilk returndeki ile fonk dan cıkacak
	return "Bildiniz", nil

}

func Demo2() {
	mesaj, hata := TahminEt(67)

	//hatayı  kullanmak istemessek syle yazılmalı mesaj, _ :=TahminEt(5)

	fmt.Println(mesaj)
	fmt.Println(hata)
}
