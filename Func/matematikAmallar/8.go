// Berilgan uch xonali sonni onlar va yuzlar xonasidagi raqamlari o’rnini alishtirishdan xosil bo’lgan sonni toping

package main

import "fmt"

func main() {
	ExchangeNUm(456)

}

func ExchangeNUm(num int) {
	yuzlar := num / 100
	onlar := num / 10 % 10
	birlar := num % 10
	result := onlar*100 + yuzlar*10 + birlar
	fmt.Println(result)
}
