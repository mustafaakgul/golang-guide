package _4_Interfaces

import (
	"fmt"
)

//interface tipinde bir türdür interface her seyi kapsadıgı icin her sey buraya atanır
//ornegin bankada hersey tek intefaceden gider onların alt türlerin vardır yani en genelde bankada aynı interface gnderilir super bilgi
//her sey interface integer string felan o yuzden bu verileblir programlcılık dnyasında
//int stringde hesey interfacedir o yuzden herseyde kullanılablir go yu geliştirenler herseyi interface uzernden yamıslar
//alt yapısında bu vardır
func dogrula(i interface{}) {
	sayi, ok := i.(int)   //her seyde oldugu gibi hata yakalamayı yapmaız lazım ok ile atayabildimi atayamadımı hata oldumu buna bkıyrz
	fmt.Println(sayi, ok)
	//i.(int) tip donusumu yapmak icin type casting
}

func Demo3() {
	var sayi1 interface{} = "Engin"
	dogrula(sayi1)

	var sayi2 interface{} = 5
	dogrula(sayi2)

}
