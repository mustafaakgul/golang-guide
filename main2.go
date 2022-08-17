package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}
*/

func TimeFromTicks(ticks int64) time.Time {
	//base2 := time.Now().Unix()
	base := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	timeconverted := time.Unix(ticks/10000000+base, ticks%10000000).UTC()
	return timeconverted
}
func main() {

	/*sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı : ", len(sozluk))
	fmt.Println("Sözlük : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman sayısı : ", len(sozluk))
	fmt.Println("Sözlük : ", sozluk)

	deger, varMi := sozluk["pencil"] //bir deger sozlugun icinde varmı yokmu
	//szlk 2 deger dondurebliryor deger ve boolean deger
	fmt.Println(deger)
	fmt.Println("Listede olma durumu : ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)*/

	/*x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	fmt.Println(values)*/

	/*s := Student{"Chetan", "Kumar", "Bangalore", 7777777777}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}*/

	/*const name, dept = "GeeksforGeeks", "CS"

	// Calling Sprintf() function
	s := fmt.Sprintf("%s is a %s Portal.\n", name, dept)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	fmt.Println(s)*/

	/*currentTime := time.Now()
	base := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)  //.Unix()
	base1970 := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) //.Unix()
	fmt.Println("Current     ", currentTime)             //.String()
	fmt.Println("origin unix.", currentTime.UnixNano())
	fmt.Println("base   unix.", base.Unix())
	fmt.Println(TimeFromTicks(637836374825217065))
	test := 637836374825217065 / 10000000
	fmt.Println(test)
	fmt.Println("test", TimeFromTicks(int64(test)))
	fmt.Println(currentTime.UnixNano() - base.UnixNano())
	fmt.Println("637836400051995066")
	fmt.Println(currentTime.UnixNano())
	tarihx := 621355968000000000
	new := int64(currentTime.UnixNano()/100) + int64(tarihx)
	fmt.Println("new", new)
	fmt.Println(new)
	fmt.Println(TimeFromTicks(int64(new)))

	fmt.Println(currentTime.Sub(base1970).Nanoseconds() / 100)

	*/

	//todo := Todo{1, 2, "Alışveriş yapılacak", false}
	//jsonTodo, err := json.Marshal(todo)

	/*myurl := "https://4com.manas.com.tr/login?ajax=true"
	data := url.Values{}
	data.Add("username", "atlasyazilim")
	data.Add("password", "123654a")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	content, _ := ioutil.ReadAll(response.Body)
	bodyString := string(content)
	fmt.Println(bodyString)
	//fmt.Println(content)*/

	/*
		response, err := http.DefaultClient.Do(http.Request{
				Method: "POST",
				URL:    &url.URL{Scheme: "https", Host: "4com.manas.com.tr", Path: "/login?ajax=true"},
				Header: map[string][]string{
					"Content-Type": {"application/x-www-form-urlencoded"},
				},
				Body: ioutil.NopCloser(strings.NewReader(data.Encode())),
			})
	*/

	//myurl := "https://4com.manas.com.tr/login?ajax=true"

	for i := 0; i < 1; i++ {
		data := url.Values{}
		data.Add("username", "atlasyazilim")
		data.Add("password", "123654a")
		req, err := http.NewRequest("POST", "https://4com.manas.com.tr/login?ajax=true", strings.NewReader(data.Encode()))
		//req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Length", "394")
		req.Header.Set("Host", "4com.manas.com.tr")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Cache-Control", "no-cache")
		//req.Header.Set("User-Agent", "PostmanRuntime/7.29.0")
		//req.Header.Set("Postman-Token", "14c11a3c-39e8-468b-8f55-1db185271e06")

		if err != nil {
			fmt.Println(err)
		}
		client := &http.Client{}
		response, err := client.Do(req)
		defer response.Body.Close()
		fmt.Println(response.StatusCode)
		fmt.Println(response.Status)
		fmt.Println(response.Header.Get("Set-Cookie"))
		content, _ := ioutil.ReadAll(response.Body)
		bodyString := string(content)
		fmt.Println(bodyString)

		cookie := response.Header.Get("Set-Cookie")
		fmt.Println("cookiemin siki: ", cookie)
		fmt.Println("----------")
		client2 := &http.Client{}
		req2, err := http.NewRequest("GET", "https://4com.manas.com.tr//definitions/meter/fetch?serial_nr=40017372", nil)
		req2.Header.Set("Cookie", cookie)
		response2, err := client2.Do(req2)
		defer response.Body.Close()
		content2, _ := ioutil.ReadAll(response2.Body)
		bodyString2 := string(content2)
		fmt.Println(bodyString2)

	}

	//fmt.Println(content)

	/*response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)*/

}
