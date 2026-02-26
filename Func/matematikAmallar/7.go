// Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!

package main

import "fmt"

func main() {
	ChangeNum(45)

}

func ChangeNum(num int) {
	onlik := num / 10
	birlik := num % 10
	result := birlik*10 + onlik
	fmt.Println(result)
}
