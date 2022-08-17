package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://4com.manas.com.tr/login?ajax=true"
	method := "POST"

	payload := strings.NewReader(`{
    "username": "atlasyazilim",
    "password": "123654a"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	cookie := res.Header.Get("Set-Cookie")
	fmt.Println("-------------------------------------------------------")

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	// Second request
	client2 := &http.Client{}
	req2, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req2.Header.Add("Content-Type", "application/json")
	req2.Header.Add("Cookie", cookie)
	fmt.Println("-------------------------------------------------------")

	res2, err := client2.Do(req2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res2.Body.Close()

	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body2))
}
