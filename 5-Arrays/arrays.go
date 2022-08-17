package arrays

import "fmt"

func Demo1() {
	var sayilar [5]int
	sayilar[0] = 50
	fmt.Println(sayilar[2])  //digerleri 0 oılur int lerin default degeri 0 hepsini doldurur
	fmt.Println(sayilar)
}
