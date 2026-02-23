//	15.02.2026
// muallif : Kamron

// sharti = Ikki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son

package main

import "fmt"

func main() {

	var num int
	num = 33

	onlik := num / 10
	birlik := num % 10

	polindrom := onlik == birlik && num%2 != 0

	fmt.Println(polindrom)

}
