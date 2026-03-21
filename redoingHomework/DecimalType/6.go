// Berilgan uch xonali sonni raqamlari yig’ndisni toping!

package main

import "fmt"

func main() {

	var num int = 123

	yuzlar := num / 100
	onlar := num / 10 % 10
	birlar := num % 10

	result := yuzlar + onlar + birlar

	fmt.Println(result)
}
