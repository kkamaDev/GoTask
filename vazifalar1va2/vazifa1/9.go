// 14.02.2026
// muallif : Kamron

// sharti = Berilgan 4 xonali sonning juft o’rindagi raqamlari yi’gindisidan toq o’rindagi raqamlari yig’indising ayrimasini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var num int
	num = 1234

	t1 := num / 1000
	j1 := num / 100 % 10
	t2 := num / 10 % 10
	j2 := num % 10

	natija := (j1 + j2) - (t1 + t2)

	fmt.Println(natija)
}
