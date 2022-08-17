package __Slices

import (
	"fmt"
)

func Demo2() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}  //slice tnmlama buyuklugunu belirtmeden another slice defi.
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3])    //python gibi hangi indisten baslayacaksn hangisine hangisi dahil hangisi degil aynısı
	fmt.Println(sehirler[1:])
	fmt.Println(sehirler[:2])

}
