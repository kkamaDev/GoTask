// 2 - qismi

package main

import "fmt"

func main() {

	var (
		a int = 9
		b int = 8
	)
	fmt.Println(a, b)
	c := a + b
	a = c - a
	b = c - b

	fmt.Println(a, b)

}
