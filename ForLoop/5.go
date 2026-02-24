// sana ; 22,02,2026
// maullif = Kamron

// sharti = 1 dan 10 gacha bo’lgan sonlar ichidan juftlarini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {

	var num int = 1

	for num < 10 {
		if num%2 == 0 {
			fmt.Println(num, "juft ")
		}
		num++

	}
}
