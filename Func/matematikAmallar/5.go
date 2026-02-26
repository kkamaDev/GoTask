// Berilgan ikki xonali sonni raqamlari yig’indisini toping!

package main

import "fmt"

func main() {

	twoNumSum(99)

}

func twoNumSum(num int) {

	onlik := num / 10
	birlik := num % 10
	result := onlik + birlik
	fmt.Println(result)

}
