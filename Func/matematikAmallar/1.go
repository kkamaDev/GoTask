package main

import "fmt"

func main() {

	var x, y int
	fmt.Print("a ni kiritng : ")
	fmt.Scan(&x)
	fmt.Print("b ni kiritng : ")
	fmt.Scan(&y)

	yuzVaPerimetr(x, y)

}

func yuzVaPerimetr(a, b int) {

	peremetri := (a + b) * 2
	yuzi := a * b
	fmt.Println("to'g'ri to'rt burchakni peremetri : ", peremetri)
	fmt.Println("to'g'ri to'rt burchakni yuzasi : ", yuzi)

}
