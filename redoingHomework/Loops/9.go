// kkilik sanoq tizimida berilgan sonni o’nlik sanoq tizimiga o’tkazuchi dastur tuzing
package main

import "fmt"

func main() {
	var a int

	fmt.Print("a = ")
	fmt.Scan(&a)

	result := change(a)

	fmt.Println(result)
}
func change(ikkilikSon int) int {
	onlik := 0
	daraja := 1
	for ikkilikSon > 0 {
		qoldiq := ikkilikSon % 10

		onlik += daraja * qoldiq
		daraja = daraja * 2

		ikkilikSon = ikkilikSon / 10
	}

	return onlik

}
