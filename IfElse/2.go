// sana : 17.02.2026
// muallif : Kamronbek

// sharti : a, b va c sonlari berilgan. Bu sonlarning eng kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a > b && a > c {
		fmt.Println(a, "soni  eng katta ")
	} else if b > a && b > c {
		fmt.Println(b, "soni k eng atta")
	} else if c > a && c > b {
		fmt.Println(c, "soni eng katta ")
	} else {
		fmt.Println("sonlar teng ")
	}

}
