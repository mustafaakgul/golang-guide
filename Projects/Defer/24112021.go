package main

func main() {
	y := 6
	defer disconnect(y)
	defer sondefer()
	
	connect()
	println("islemler")

}

//Geri donus parametrelerimi metodumu olustururken isimlendirdigim icin asagida bir daha tanimlamam gerek kalmiyor.
func split(sum int) (x int, y int) {
	x = sum * 4
	y = sum - x
	return x, y
}

func connect() {
	println("baglandi")
}

func disconnect(x int) {
	x=5
	println("baglanti koptu")
}

func sondefer() {
	println("son defer")
}
