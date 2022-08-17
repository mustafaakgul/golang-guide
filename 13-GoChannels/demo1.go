package _3_GoChannels

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiCn chan int) { //bu satırı channal ile actık bu işlemin baslangıc noktasıdırı
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Çift sayı toplama çalışıyor")
		time.Sleep(1 * time.Second)     //burayı yorum yapıp calisma mantıgını direk goruleblr ama asenkron oldugu gozumuzle goremeyiz
	}

	ciftSayiCn <- toplam   //bu satırda işlemlerin bitiş noktasıdır
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
		fmt.Println("Tek sayı toplama çalışıyor")
		time.Sleep(1 * time.Second)
	}

	tekSayiCn <- toplam
}
