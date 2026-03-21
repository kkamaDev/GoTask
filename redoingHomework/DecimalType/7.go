// Berilgan ikki xonali soni raqamlari o’rini almashtirishdan xosil bo’lgan sonni toping!

package main

import "fmt"

func main() {
	var num int = 12

	onlar := num / 10
	birlar := num % 10

	result := birlar*10 + onlar

	fmt.Println(result)
}
