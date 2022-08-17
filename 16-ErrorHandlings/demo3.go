package _6_ErrorHandlings

import "fmt"

type borderException struct {
	parameter int
	message   string
}

//ister pointer ister pointer yani basnda yıldız olmadan calısır pinter gibi verilirse performans artar
func (b *borderException) Error() string {   //error interfacenin impleemente ettigimiz yapı
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {//sartlar saglanırsa direk error fırlatcagız
		return "", &borderException{tahmin, "Sınırların dışında kaldı"}
	}

	return "Bildiniz", nil
}

// KENDI HATA YAPILARIMIZI OLUSTURMA
// CUSTOM ECXEPTION YAZMAK
//struct ile hatanın detayları belirlenecek