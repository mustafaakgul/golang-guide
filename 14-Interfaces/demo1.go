 package _4_Interfaces

import (
	"fmt"
	"math"
)

//bu yapıda yeni bir sekl eklennce sadece onu ekleyip diger satırları hic bozmadan devam edebiliriz

//ortak fonk interfacelern altına koyablrz
type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {  // rectangle icin yapıyoruz demek syruct metodu oldugunu anlıyoruz
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {  //parametre olarak bir interface istemisiz bize bir shape interface ver
	fmt.Println("Şeklin alanı : ", s.area())
//buradaki shape area arıyor icinde area olan her seyi alabilirsin gndereblrsn
}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r) //syle yrmluyoruz gonderilen r rectangle tipinde bir struct s.area yı arıyor rectagle icin var ve onu ypıor
	//shape imzasına uyan bir operasyon yolla
	//bnlar pointer degeri ile calısıyor o yuzden algılıyor bunları
	//rec yolladın o yuzden rect ın alanı calisiyor

	c := circle{radius: 10}
	school(c)
}
