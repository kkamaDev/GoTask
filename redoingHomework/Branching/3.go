// a, b va c sonlari berilgan. Bu sonlarning o’rachasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b, c int = 1, 2, 3
	if a > b && a < c {
		fmt.Printf("%d soni o'rtacha son\n", a)
	} else if b > a && b < c {
		fmt.Printf(" o'rtacha son %d\n  ", b)
	} else {
		fmt.Printf("o'rtacha son %d\n", c)
	}
}
