// sana : 22.02.2026
// muallif = Kamron

// sharti : = Son kiritiladi (1, 10) kiriltilgan sonni so’zlar orqali ifodalovchi dastur tuzing

package main

import "fmt"

func main() {
	var num int
	fmt.Print("son kiriting ")
	fmt.Scan(&num)

	switch num {
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
