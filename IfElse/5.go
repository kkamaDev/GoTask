// sana : 17.02.2026
// muallif : Kamronbek

// sharti =  soni berilgan! Agar a soni 3 ga va 5 ga bo’linsa FIZZBUZZ, agar faqat 3 ga bo’linsa FIZZ,
//  agar faqat 5 ga bo’linssa BUZZ so’zini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {
	var son int
	fmt.Print("son kiritng >> ")
	fmt.Scan(&son)

	if son%3 == 0 && son%5 == 0 {
		fmt.Println(son, "FIZZBUZZ")
	} else if son%3 == 0 {
		fmt.Println(son, "FIZZ")
	} else if son%5 == 0 {
		fmt.Println(son, "Buzz")
	}

}
