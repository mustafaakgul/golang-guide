package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
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
	req2, err := http.NewRequest("POST", "https://4com.manas.com.tr/login?ajax=true", strings.NewReader(data.Encode()))
	req2.Header.Set("Cookie", cookie)
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Set("Content-Length", "394")
	req2.Header.Set("Host", "4com.manas.com.tr")
	req2.Header.Set("Connection", "keep-alive")
	req2.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req2.Header.Set("Accept", "*/*")
	req2.Header.Set("Cache-Control", "no-cache")
	response2, err := client2.Do(req2)
	defer response.Body.Close()
	content2, _ := ioutil.ReadAll(response2.Body)
	bodyString2 := string(content2)
	fmt.Println(bodyString2)
}
