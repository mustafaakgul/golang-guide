package _5_DeferStatements

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc() //burada ise ilk defere lifoya atacak snra return olursa calısacak yani defer her sekilde calısır
	//yni deferi ilk yerleri almk mantıklı  yoksa returnu gorunce fonk sonlandırablir

	if sayi%2 == 0 {
		fmt.Println("çift sayı çalıştı")
		return "Çift sayıdır"
	}

	if sayi%2 != 0 {
		return "Tek sayıdır"
	}
	//defer DeferFunc()  //burada defer olursa LIFO ya atamayacak buraya gelmicek cnku fonk ama ilk returnleri
	// calıstıracagı icin oraya giremeyecek bnn yerine yukarda olacakki lıfoya atalım burada cnku lifoya atamıyoruz

	return "Belli değil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
