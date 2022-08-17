package _5_DeferStatements

import (
	"fmt"
)

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {  //strunc baglamak icin p product yazarız
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log()  //burada log tutmak cok mantıklı islem btmesnn ardnda direk calıstılmalı
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {//agoritm icinde deferi nereye konacagı cok onemlidir sırası felan
	//arada defer ile fonk olsaydı defer i son calıstırırdı
	p := product{productName: "Laptop", unitPrice: 5000}
	//defer p.save()  burada olsayd laptopu calıstırrıdı cnku burada lifoya ekledi laptop icin ekledi
	p = product{productName: "Mouse", unitPrice: 45}

	fmt.Println("İşlem başarılı")
	fmt.Println(p.productName)
	defer p.save()  //burad ise su calısır su an p mouse oldugunda mouse icin calısacak
}
