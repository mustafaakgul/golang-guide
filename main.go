package main //package main maini paketledik buradaki kodları buutnlestirdik bu pakette bunlar var demek
import (
	_9_Project "btktraining/19-Project"
	"fmt"
)

//import variables "btktraining/2-Variables"

//birbiri ile ilişkili dosysları aynı yerlere koyarız

//import "btktraining/2-Variables"
/*
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "ACME INC")
}*/

func main() {

	//variables.Demo1()
	//conditionals.Demo1()

	//variables.Demo1()
	//fmt.Print()
	//conditionals.Demo3()
	//loops.Demo6()
	//arrays.Demo4()
	//slices.Demo2()
	//functions.SelamVer("Engin")    //fmt package name and selamver is a function name
	//functions.SelamVer("Derin")
	//var sonuc = functions.Topla(2, 6)
	//fmt.Println(sonuc * 10)

	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6) //istenmeyen yere _ koyulursa o alınmamıs oluyor super bu
	//sonuc4 u ihtiyacın yoksa bunu kullanılalbr
	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)

	// fmt.Println("Toplam : ", sonuc1)
	// fmt.Println("Çıkarım : ", sonuc2)
	// fmt.Println("Çarpım : ", sonuc3)
	//fmt.Println("Bölüm : ", sonuc4)

	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 7, 0, 11}   // array bile parametre olarak gonderileblir
	// fmt.Println(functions.ToplaVariadic(sayilar...))  //variadic funcs

	/*sayi := 20
	pointers.Demo0(sayi)
	fmt.Println("Maindeki sayı : ", sayi)
	sayipointer := 30
	pointers.Demo1(&sayipointer)  //yollarken ve operator ile sayiyi adresini yollamak demek
	//bellektekini yolladık direk onu buldu ve scope icinde bile arttırabldi
	fmt.Println("Maindeki sayı : ", sayipointer)*/

	// sayilar := []int{1, 2, 3}  //arrayde tek deger 20 olan gibi calısmadı global calıstı cnku array degerleri uzernden calısmaz
	//arraylaer referansları ile calisir bellekteki adresi ile arkadan gonderir array biraz ile veri ypısı primitive degil
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayı : ", sayilar[0])

	//program adım adım calistirr sırayla maine gelir buradan digerine gider calisir sirali seklde calisir
	//_2_Go_Routines.CiftSayilar()
	//_2_Go_Routines.TekSayilar()
	//islemlerin paralel calismasını istersek birbirlerinin calısmasını engellemesin buradan go routinelerde yapılıyor
	//aynı asenkron komutları calıstırmak icin basına go koy bunları baska thread acar main fonk
	//go _2_Go_Routines.CiftSayilar()
	//go _2_Go_Routines.TekSayilar()
	//bu notasyonda calısmadı cunku main basına go gorunce 2 satır icin 2 thread acar kendisini bitirir tam calıscakkken
	//calısamaz bunun calısan hali su sekldr
	//Faz1
	//go _2_GoRoutines.CiftSayilar()
	//go _2_GoRoutines.TekSayilar()
	//time.Sleep(5*time.Second)
	//fmt.Println("main bitiyor burada")
	//threadler cok hızlı acarken thredi ilk satır hemen bitiyor bile buradan asenkron oldugu anlamak zor ama cok uzunda gorulur
	//yani 2 fonk normalde aynı anda calısacak ama ilki su an cabuk bitiyor normalde farklı threadler calisti
	//thread calisiyor ancak burada kendi algoritmalar kullanıyor
	//burada mainin bitirmeyecegiz 5 sn beklicek o arada threadler calisacak snra main kendini bitirene kdar bu 5 sn
	//digerlerinin islemleri bitmiş olacak
	//peki karısık calıstıgını nasıl ispat ederiz thread olarak soyle kod yazarız fonk icine bekletme koyarız faz2 fonksiyonları ile
	//go _2_GoRoutines.CiftSayilar()
	//go _2_GoRoutines.TekSayilar()

	//structs.Demo2()
	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	//channel mantıgı
	//routşnlerde asenkron calısma mantıgını ispatlamıstık gercekten thread acıp asenkron calisiyor
	//ciftSayiCn := make(chan int)  //tipi channel int ayrı degil onlar bunlar sayi degiller sayi tasıyan bir channel tipinde
	//tekSayiCn := make(chan int)
	//go _3_GoChannels.CiftSayilar(ciftSayiCn)
	//go _3_GoChannels.TekSayilar(tekSayiCn)
	//channler işlemleri bitmeden alt satırı gecmez fonk sleep silinerek goruleblr bitmesini bekler
	//baslarına go yazmassak biri bitmeden digeri baslamazdı asenkron calısamazdık ama go yazarrak asenkron calisirdi
	//asenkron problemndede islemin bittigi anlayamamak channel ilede bunu yapıyoruz işlem bittimi bitmedimi
	//channgel bitince yaptıgını garanti eden yapı
	//asenkronda eslenik calısmasını garanti ediyoruzaynı anda fonk calisiiyor channge hem asenskron calısmasını saglıyor
	//go routinler ile hemde onlaarın yasam dongusunu kontrol etmemizi saglıyro
	//ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn  //2 channgel carpamassn snuclarını bi yere atanablr
	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım : ", carpim)
	//diger mantık sistem kitlenmeden 2 channel acıp kendi islemlerini yapıp sonucta onlarla bir işlem yaptırmak
	//channel işlemlerin bittigini garanti ettigi yapıdır

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	//fmt.Println(error_handling.TahminEt2(99))

	//restful.Demo2()

	/*product, _ := project.AddProduct()
	fmt.Println(product)

	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
	*/

	//_8_Restful_Arthitecture.Demo1()

	//_5_DeferStatements.B()
	//_5_DeferStatements.Test()
	//_5_DeferStatements.Demo3()
	//_6_ErrorHandlings.Demo1()
	//_4_Interfaces.Demo3()
	//_6_ErrorHandlings.Demo2()
	//_6_ErrorHandlings.TahminEt2(6)

	_9_Project.AddProduct()
	//products,_ := _9_Project.GetAllProducts()
	products, _ := _9_Project.AddProduct()
	fmt.Println("all products", products)

	i := 0
	y := 0
	for i = 1; i <= 10; i = i + 2 {
		y++
	}
	fmt.Println(i)
	fmt.Println(y)
}
