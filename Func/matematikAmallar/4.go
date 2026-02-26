// a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching

package main

import "fmt"

func main() {
	change(4, 5)

}

func change(a, b int) {
	fmt.Println(a, b)

	c := a
	a = b
	b = c

	fmt.Println(a, b)

}
