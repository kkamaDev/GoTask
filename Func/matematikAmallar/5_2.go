// Berilgan ikki xonali sonni raqamlari yig’indisini toping!
package main

import "fmt"

func main() {
	natija := twosumNum(44)
	fmt.Println(natija)

}

func twosumNum(num int) int {
	onlik := num / 10
	birlik := num % 10
	result := onlik + birlik

	return result
}
