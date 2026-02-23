// // 14.02.2026
// // muallif : Kamron
// shart = Berilgan ikki xonali sonni raqamlari yig’indisini toping!

package main

import "fmt"

func main() {
	var num int
	fmt.Print("ikki honali son kiritng > ")
	fmt.Scan(&num)

	onlik := num / 10
	birlik := num % 10
	sum := onlik + birlik

	fmt.Println(sum)

}
