package variables
//normalde package name ustteki kalsr adini aliyor

import "fmt"

//camelCase
//PascalCase
func Demo1() {
	var metin string = "Merhaba Dünya!"
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	//integer
	var kdv int = -10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25.2 //:= olarak yazılırsa veri tipini kendi farkediyor yada esittirli kllncaksan tipini verceksn
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	var durum bool
	var state bool = true

	fmt.Println(state)

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	//true false
	durum = metin1 != metin2 //metin1 metin 2 esit degilmi demek oyleyse durum false veya true olcak
	//metin1==metin2 olsaydı metin 1 metin 2 ye esitmiye bakcak false oldugundan drum false olcak

	fmt.Println(durum)
	fmt.Println(!durum)

	var name string
	name = "mustafa"
	var age int
	age = 59
	fmt.Println(name, age)

	var first, last string = "mstafa", "akful"
	fmt.Println(first, last)

	name1 := "Arin"
	age1 := 40
	isMarried := true

	age1 = 41

	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(isMarried)
}
