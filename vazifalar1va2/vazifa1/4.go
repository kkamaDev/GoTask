// // 14.02.2026
// // muallif : Kamron

// shart = a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching)

package main

import "fmt"

// // 1- usul
// func main() {
// 	var a, b int
// 	a = 3
// 	b = 9

// 	fmt.Println(a, b)

// 	x := a
// 	a = b
// 	b = x

// 	fmt.Println(a, b)

// }

// 2- usul

func main() {
	var a, b int
	a = 5
	b = 9
	fmt.Println(a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

}
