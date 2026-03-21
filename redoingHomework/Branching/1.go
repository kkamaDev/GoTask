// a va b sonlari berilgan. Bu sonlarning kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {

	var (
		a int = 9
		b int = 10
	)
	if a > b {
		fmt.Printf(" %d soni katta \n", a)
	} else {
		fmt.Printf("%d soni katta \n ", b)
	}
}
