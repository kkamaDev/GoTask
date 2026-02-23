// // 14.02.2026
// // muallif : Kamron
// shart = Berilgan uch xonali sonni raqamlari yig’ndisni toping!

package main

import "fmt"

func main() {

	var num int
	num = 123
	yuzlik := num / 100
	onlik := num / 10 % 10
	birlik := num % 10

	sum := yuzlik + onlik + birlik

	fmt.Println(sum)
}
