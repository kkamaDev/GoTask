// Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!

package main

import "fmt"

func main() {
	result := changenum(87)
	fmt.Println(result)

}

func changenum(num int) int {

	onlik := num / 10
	birlik := num % 10
	result := birlik*10 + onlik
	return result
}
