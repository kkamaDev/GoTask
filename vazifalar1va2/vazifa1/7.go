// 14.02.2026
// muallif : Kamron

// sharti = Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!

package main

import "fmt"

func main() {
	var num int
	num = 45
	onlik := num / 10
	birlik := num % 10
	change := birlik*10 + onlik

	fmt.Println(change)
}
