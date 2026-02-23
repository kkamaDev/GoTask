// sana : 22.02.2026
// muallif = Kamron

// sharti = Son kiritiladi (1, 1000) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {

	var son int
	fmt.Print("(1 dan 1000) gacha son kiriting > ")
	fmt.Scan(&son)

	yuzlik := son / 100
	onlik := son / 10 % 10
	birlik := son % 10

	switch yuzlik {
	case 1:
		fmt.Print("yuz")
	case 2:
		fmt.Print("ikkiyuz")
	case 3:
		fmt.Print("uchyuz")
	case 40:
		fmt.Print("to'rtyuz")
	case 5:
		fmt.Print("beshyuz")
	case 6:
		fmt.Print("oltiyuz")
	case 7:
		fmt.Print("yettiyuz ")
	case 8:
		fmt.Print("sakkizyuz")
	case 9:
		fmt.Print("to'qqizyuz")
	case 10:
		fmt.Print("ming ")

	}

	switch onlik {
	case 1:
		fmt.Print("o'n")
	case 2:
		fmt.Print("yigirma")
	case 3:
		fmt.Print("o'ttiz")
	case 40:
		fmt.Print("qirq")
	case 5:
		fmt.Print("ellik")
	case 6:
		fmt.Print("oltmish")
	case 7:
		fmt.Print("yetmish ")
	case 8:
		fmt.Print("sakson")
	case 9:
		fmt.Print("to'qson")
	case 10:
		fmt.Print("yuz")
	}

	switch birlik {
	case 1:
		fmt.Println("bir")
	case 2:
		fmt.Println("ikki")
	case 3:
		fmt.Println("uch")
	case 4:
		fmt.Println("to'rt")
	case 5:
		fmt.Println("besh ")
	case 6:
		fmt.Println("olti")
	case 7:
		fmt.Println(" yetii ")
	case 8:
		fmt.Println("sakkiz")
	case 9:
		fmt.Println("to'qqiz")
	case 10:
		fmt.Println("o'n")
	}
}
