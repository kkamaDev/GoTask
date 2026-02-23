//	15.02.2026
// muallif : Kamron

// sharti = Uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juf son

package main

import "fmt"

func main() {

	var num int
	num = 212
	a1 := num / 100
	a3 := num % 10

	polindrom := a1 == a3 && num%2 == 0

	fmt.Println(polindrom)
}
