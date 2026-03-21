// a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	kichik(5, 3, 5)

}

func kichik(a, b, c int) {

	if a < b && a < c {
		fmt.Printf("%d soni kichik\n ", a)
	} else if b < a && b < c {
		fmt.Printf("%d soni kichik \n ", b)
	} else {
		fmt.Printf("%d soni kichik \n  ", c)
	}

}
