package _6_ErrorHandlings

import (
	"fmt"
	"os"
)

//metin dosyasının okunması ile ilgili hata yakalama
//type assertion
func Demo1() {
	f, err := os.Open("demo12.txt")  //2 donusludr dosyann kendisi ve hata olayı iste bu go da cok iyi ve farklı
	//varsa dosyayı oyksa error donecek bu ust satırdkai yapı eger dosya bulursa error nil olacak dosya blundu yani
	//hata olup olmamasını burada nil oldugunda anlarız dosyayı bulamaması
	//nil
	if err != nil { //error varsa
		fmt.Println("dosya blunamadı")
		if pErr, ok := err.(*os.PathError); ok {   // bu syntax pErr, ok := err.(*os.PathError) buraya kadar atama yapıyor
			// bu true ise ok yani if icine gelecek bu yeni syntax if kısaltması gibidir bu

			// satır hatayı profesyonelleştirmedir hatanın turune gore sistemi
			//düzeltmekle ilgili hatayı println ile yaprız ama pErr satırı ile hatanın daha derinine gireiyorz
			//patherror dosya yolu ile ilgili bir hayaymıs yani patherror o yolda dosya olup olmadıgı ile ilgilidir
			//baska bi suru oalblrdi dosya baska biri acmıstır dosya orada yoktur ..... yani bnların butun olasılıklarını
			//aslında yazmalıyız
			fmt.Println("Dosya bulunamadı : ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı") //bu patherror türünde bir hata degilse ise girecek
			return
		}
	} else {
		fmt.Println(f.Name()) //bu dosya ile ilglil operasyonları gosterr
	}

}
