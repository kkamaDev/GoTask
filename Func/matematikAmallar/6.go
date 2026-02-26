// Berilgan uch xonali sonni raqamlari yig’ndisni toping!

package main

import "fmt"

func main() {
	ThreeNumsum(123)

}

func ThreeNumsum(num int) {

	yuzlik := num / 100
	onlik := num / 10 % 10
	birlik := num % 10

	result := yuzlik + onlik + birlik
	fmt.Println(result)
}
