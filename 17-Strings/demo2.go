package _7_Strings

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng"))  //boyle baslarmı
	fmt.Println(s.HasSuffix(isim, "en"))  //boyle bitermi
	fmt.Println(s.Index(isim, "gi"))  //basladıgın yerin index i doner
	harfler := []string{"e", "n", "g", "i", "n"}   //string joinleri
	sonuc := s.Join(harfler, "*")     //sep koyuyorsal arasına * koyar optional sep sanırım
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", 3))  //stringleri yer degistirmek n yerine -1 koyarsak hepsini degistir pozitif degerde kac tane degistirlecegi
	fmt.Println(s.Split(sonuc, "-"))   //stringleri ayırır belli formatlara gore ihtiyac oluyor  array olarak return eder
	fmt.Println(s.Repeat(sonuc, 5))
}
