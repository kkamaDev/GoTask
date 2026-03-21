// Berilgan ikki xonali sonni raqamlari yig’indisini toping!

package main

import "fmt"

func main() {
	var num int = 45

	onlar := num / 10
	birlar := num % 10

	result := onlar + birlar

	fmt.Println(result)

}
