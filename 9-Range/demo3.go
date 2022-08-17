package __Range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range sozluk { // key ve valualeri gezeriz
		fmt.Print(k)
		fmt.Println(v)
	}

	for k := range sozluk { // ilk yazılan her zaman key dir
		fmt.Print(k)
	}
}
