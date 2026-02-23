//	15.02.2026
// muallif : Kamron

// sharti = a va b sonlari berilgan! Bu sonlarning yig’indisi musbatligini va toqligini aniqlovchi dastur tuzing!

package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("a ni kiriting : ")
	fmt.Scan(&a)

	fmt.Print("b ni kiritng : ")
	fmt.Scan(&b)

	sum := a+b > 0 && a+b%2 != 0

	fmt.Println(sum)

}
