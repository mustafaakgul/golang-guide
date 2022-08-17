package _8_RestfulArthitecture

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//benzetecegimizi veri turu struct ile yapılır
//direk map leme kendisinde var süper
type Todo struct {
	UserId    int    `json:"userId"`       //userId apideki neye karsılık gelecegini gosterr neye map ledigimiz digeri bizimkidir
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	//kendi hhtp.Get var kendisinden kullanıyoruz 2 deger donuyor hep cevap ve error dur hersey 2 donus tipi var
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()   //refer ilk calısıyor eger hata alırsak http.get de direk baglantıyı kapatır bu super

	bodyBytes, _ := ioutil.ReadAll(response.Body)    //_ pasif braktık koymamız lazım ama kullanmayacagım demek
	//byte formatında alıyoruz bunu donusu bu seklde oluyor

	bodyString := string(bodyBytes)  // bu body datasının string haline cevirdk
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)   //bu method ile json objesini  vartodo objesine aktardık
	//json to &tdo object data convertion
	fmt.Println(todo)

}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo)

	//2 parametre contnttype dır json formatnda data yolluyorm birden fazla data format varsa ; ile yan yana yazarız
	//3. parametre body post eetigiiz json turunde veri tipdir bu bytes formatında buffer ile yollarız
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()    //bu defer http istedi icin bir hata fln cıakrsa yapsın diye oluyor
	//hata yazılımı ve http.post defer ile kapatma işlemei bir problem olmasına karsın

	//burdan sonrası post edilecek seyi geri donuyor yaptımı get in aynısını yapıyoruz normalde bu satıra kadar post yeterli
	//test db si kaydetmiyor kaydettigini geti donuor test icin bunu yaptık

	bodyBytes, _ := ioutil.ReadAll(response.Body)   //body okumak string cevirmek donen seyi

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
