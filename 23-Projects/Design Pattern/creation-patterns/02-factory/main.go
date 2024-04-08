package main

import "fmt"

func main() {
	news1, _ := newPublication("newspaper", "Yazılımcı Türünün Yaşamı", 546, "Doğada Hayatta Kalmak")
	// news2, _ := newPublication("newspaper", "Yazılımcılar Yemek Yer mi?", 456, "The Newyork Times")
	mag1, _ := newPublication("magazine", "Hacker'lar da Sever", 123, "Murtaza Yayınevi")
	// mag2, _ := newPublication("magazine", "Hatasız Kod Olmaz!", 120, "Cabbar Yayınları")

	printPublicationDetails(news1)
	// printPublicationDetails(news2)
	printPublicationDetails(mag1)
	// printPublicationDetails(mag2)
}

func printPublicationDetails(pub iPublication) {
	fmt.Printf("---------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("---------------\n")
}
