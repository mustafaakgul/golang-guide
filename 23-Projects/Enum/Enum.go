package main

type Season int64

// Const ile Enum olustururken değerleri manuel olarak atamak yerine kullanmak iota kullanarak aynı değeri yanlışlıkla birden fazla sabite atamamızı da önler,
// aksi takdirde bir çakışmaya neden olabilir.
const (
	Summer Season = iota
	Autumn
	Winter
	Spring
)

//Iota kullanilmadan degerler manuel olarak atanmistir.
const (
	Yaz      Season = 0
	Ilkbahar        = 1
	Kis             = 2
	Sonbahar        = 3
)

func main() {

	var today Season = Winter

	if today == Winter {
		print("Winter")
	}
}

//Kaynak : https://www.sohamkamani.com/golang/enums/
