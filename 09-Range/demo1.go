package __Range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for i, sehir := range sehirler { //buradaki i kacıncı index te oldugunu gosterir i kullanmak istemiyorsak alttaki gibi
		fmt.Print(i)
		fmt.Println(sehir)
	}

	for _, sehir := range sehirler { //i kllnmıcaksan _ koyarsak orada sorun cıkarmaz kullanmak istemedgmz her yerde _
		fmt.Println(sehir)
	}
}


//for-range yada range gezinmeyi saglayan yapılrı iteratorler gibi
//yukarda ilki for ile digeri ise for range mantıgı ile geznmek