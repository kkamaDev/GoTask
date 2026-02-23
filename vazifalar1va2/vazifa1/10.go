// 14.02.2026
// muallif : Kamron

// sharti = Berilgan 4 xonali sonning raqamlarini teskari tartibda yozish orqali xosil bo’ladigan sonni toping!

package main

import "fmt"

func main() {
	var num int
	num = 1234

	th := num / 1000
	hn := num / 100 % 10
	ny := num / 10 % 10
	one := num % 10

	natija := one*1000 + ny*100 + hn*10 + th
	fmt.Println(natija)
}
