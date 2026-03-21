// a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching)
// a = 5, b = 7
// b = 7  a = 5

package main

import "fmt"

func main() {

	var (
		a int = 9
		b int = 8
	)
	fmt.Println(a, b)
	c := a
	a = b
	b = c

	fmt.Println(a, b)

}
