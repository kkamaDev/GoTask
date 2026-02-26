package main

import "fmt"

func main() {
	var x, y int = 4, 2

	m, k := YuziPeremetr(x, y)
	fmt.Println("yuzi : ", m)
	fmt.Println("peremetri: ", k)

}

func YuziPeremetr(a, b int) (int, int) {

	yuzi := (a + b) * 2
	peremetr := a * b

	return yuzi, peremetr

}
