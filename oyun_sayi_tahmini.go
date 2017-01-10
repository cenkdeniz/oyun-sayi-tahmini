package main

import "fmt"

const MAX = 20
const MIN = 1

var sayi int
var win int

func main() {
	win = 10
	oyun()
}

func oyun() {

	fmt.Printf("%d ile %d arasında bir sayı söyle\n", MIN, MAX)
	fmt.Scanln(&sayi)

	if (sayi == win) {
		fmt.Printf("Tebrikler tuttuğum sayı %d\n", sayi)

	} else if (sayi > win) {
		fmt.Printf("%d tuttuğum sayıdan büyük\n", sayi)
		oyun()

	} else if (sayi < win) {
		fmt.Printf("%d tuttuğum sayıdan küçük\n", sayi)
		oyun()

	}
}


