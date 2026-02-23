//  15.02.2026
// muallif : Kamron

// sharti = Berilgan a soni juftligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {

	var a int
	fmt.Print("Sonimizni juftligini tekshiramiz (son kiriting ) > ")
	fmt.Scan(&a)

	juft := a%2 == 0

	fmt.Println("sonimiz juiftmi ? :", juft)
}
