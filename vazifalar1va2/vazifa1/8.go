// 14.02.2026
// muallif : Kamron

// sharti = Berilgan uch xonali sonni onlar va yuzlar xonasidagi raqamlari o’rnini alishtirishdan xosil bo’lgan sonni toping

// 345 = 435

package main

import "fmt"

func main() {
	var num int
	num = 345
	yuzlar := num / 100
	onlar := num / 10 % 10
	birlar := num % 10

	natija := onlar*100 + yuzlar*10 + birlar

	fmt.Println(natija)
}
