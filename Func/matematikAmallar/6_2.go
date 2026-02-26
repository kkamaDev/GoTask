// Berilgan uch xonali sonni raqamlari yig’ndisni toping!
package main

import "fmt"

func main() {

	result := threesumnum(555)
	fmt.Println(result)
}

func threesumnum(num int) int {
	yuzlik := num / 100
	onlik := num / 10 % 10
	birlik := num % 10
	result := yuzlik + onlik + birlik

	return result
}
