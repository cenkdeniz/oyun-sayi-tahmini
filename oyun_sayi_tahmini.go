package main

import "fmt"

const MAX = 20
const MIN = 1
const NUMBER  = 10

func main() {
	oyun()
}

func oyun() {

	var sayi int

	fmt.Printf("%d ile %d arasında bir sayı söyle\n", MIN, MAX)
	fmt.Scanln(&sayi)

	if (sayi > NUMBER) {
		fmt.Printf("%d tuttuğum sayıdan büyük\n", sayi)
		oyun()

	} else if (sayi < NUMBER) {
		fmt.Printf("%d tuttuğum sayıdan küçük\n", sayi)
		oyun()
	} else {
		fmt.Printf("Tebrikler tuttuğum sayı %d\n", sayi)
	}
}


