// 14.02.2026
// muallif : Kamron

// // To’gri to’rtburchakni a va b tomonlari berilgan uning yuzini va perimetrini topuvchi dastur tuzing!

package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("to'g'ri to'rt burchakning a tomonini kiriting >> ")
	fmt.Scan(&a)

	fmt.Print("to'g'ri to'rtburchakni b tomonini kiritng >> ")
	fmt.Scan(&b)

	returnS := a * b
	returnP := 2 * (a + b)

	fmt.Println("Peremetri", returnP)
	fmt.Println("yuzasi :", returnS)

}
