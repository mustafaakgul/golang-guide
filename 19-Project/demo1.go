package _9_Project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil { //hata varsa
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("getall")

	defer response.Body.Close() //yaptıgımız istedi her türlü kapatmamız gerekiyor defer ile

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)  //aynı adresi koyarak direk atama yaparız bu sayede bisey degisince sonucta degisir
	//&koymasaydık degiisnce api yeni alan accagından dolayı ramde eskisi kalcaktı o yüzden pointer ile gncelde tutarız
	fmt.Println(products)
	return products, nil
}

func AddProduct() (Product, error) {
	product := Product{ProductName: "Microphone 3", CategoryId: 1, UnitPrice: 2000.0}  //id yazmaaya grek yok
	jsonProduct, err := json.Marshal(product)  //json a cevirme

	response, err := http.Post("http://localhost:3000/products",
		"application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println("error occurred")
		fmt.Println(err)
		return Product{}, err
	}

	defer response.Body.Close()   //acılam response kapamak

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)
	fmt.Println("kaydedild", productResponse)
	return productResponse, nil
}

func hello()  {
	fmt.Println("hellotest")

}