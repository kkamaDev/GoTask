// Berilgan 4 xonali sonning raqamlarini teskari tartibda yozish orqali xosil bo’ladigan sonni toping!

package main

import "fmt"

func main() {

	natija := exchangefourNum(1234)
	fmt.Println(natija)

}

func exchangefourNum(son int) int {

	a1 := son / 1000
	a2 := son / 100 % 10
	a3 := son / 10 % 10
	a4 := son % 10

	natija := a4*1000 + a3*100 + a2*10 + a1
	return natija

}
