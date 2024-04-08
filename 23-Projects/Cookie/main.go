package main

import (
	"fmt"
	"log"
	"net/http"
)

//Cookie
//Cookie üzerinde kritik olmayan bilgiler tutulmalıdır.

func main() {
	http.HandleFunc("/", setCerez)
	http.HandleFunc("/check", cerezKontrol)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello")
}

func setCerez(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "kullanicibilgi01",
		Value: "Bu da cookienin verisi",
		Path:  "/",
	})
	fmt.Fprintln(w, "Cookileri kontrol et!")
}

func cerezKontrol(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("kullanicibilgi01")
	if err != nil {
		log.Fatal("No cookie found")
	}

	if c == nil {
		log.Fatal("Cookie is null")
	}

	//Cookie var ise yönlendirme işlemi yapılabilir. Cookie yok ise login sayfasına yönlendirilebilir.
	if c.Value != "" {
		fmt.Fprintln(w, "Cookie var : "+c.Value)
	} else {
		fmt.Fprintln(w, "Cookie yok")
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("kullanicibilgi01")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Cookie var : "+c.Value)
}
