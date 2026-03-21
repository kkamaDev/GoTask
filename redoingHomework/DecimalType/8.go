// Berilgan uch xonali sonni onlar va yuzlar xonasidagi raqamlari o’rnini alishtirishdan xosil bo’lgan sonni toping

package main

import "fmt"

func main() {
	var nums int = 123
	yuzlar := nums / 100
	onlar := nums / 10 % 10
	birlar := nums % 10

	result := onlar*100 + yuzlar*10 + birlar

	fmt.Println(result)
}
