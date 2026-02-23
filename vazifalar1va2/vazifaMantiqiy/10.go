//	15.02.2026
// muallif : Kamron

// sharti = To’rt xonali a soni berilgan! Jumlani rostilkka tekshiring!
// A sonining juft o’rindagi raqamlari yigindisi va toq o’rindiadi ramlari yig’indisining ayrimasi 0 ga teng
// va a soning xar-bir raqami takrorlanmas

package main

import "fmt"

func main() {

	var num int
	num = 7513

	t1 := num / 1000
	j1 := num / 100 % 10
	t2 := num / 10 % 10
	j2 := num % 10

	natija := t1 != j1 && j1 != t2 && t2 != j2 && t1 != t2 && (j1+j2)-(t1+t2) == 0

	fmt.Println(natija)
}
