package _2_GoRoutines

import (
	"fmt"
	"time"
)

//Faz 1
/*
func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı : ", i)
	}
}

func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı : ", i)
	}
}*/

//Faz2

func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift sayı : ", i)
		time.Sleep(1 * time.Second)
	}
}

func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek sayı : ", i)
		time.Sleep(1 * time.Second)
	}
}