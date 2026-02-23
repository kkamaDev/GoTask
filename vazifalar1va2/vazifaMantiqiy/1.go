// 15.02.2026
// muallif : Kamron

// sharti = Berilgan a soni musbatligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {
	var a int
	fmt.Print("sonni musbatligini tekshiramiz (son kiritng ) : ")
	fmt.Scan(&a)

	musbatNUm := a > 0
	fmt.Println("Sonimiz musbatmi ? >", musbatNUm)

}
