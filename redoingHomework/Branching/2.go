// a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	Katta(2, 3, 5)

}

func Katta(a, b, c int) {

	if a > b && a > c {
		fmt.Println(a, "katta ")
	} else if b > a && b > c {
		fmt.Printf("%d soni katta \n ", b)
	} else {
		fmt.Printf("%d soni katta \n  ", c)
	}

}
