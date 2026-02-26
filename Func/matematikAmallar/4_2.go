// a va b o’zgaruvchilar berilgan. Ularning qiymatlarni o’rini almashtiruvhi dastur tuzing! (2 xil usulda yeching

package main

import "fmt"

func main() {
	var x, y int = 4, 5
	fmt.Println(x, y)
	result, result2 := change2(x, y)
	fmt.Println(result, result2)

}

func change2(a, b int) (int, int) {

	c := a
	a = b
	b = c
	return a, b

}
