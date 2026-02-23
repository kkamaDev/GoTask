// sana : 17.02.2026
// muallif : Kamronbek

// sharti : a, b va c sonlari berilgan. Bu sonlarning o’rachasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b, c float32 = 4, 5, 9
	if a > b && a < c {
		fmt.Println(a, "o'rtacha ")
	} else if b > a && b < c {
		fmt.Println(b, "o'rtacha ")
	} else if c > a && c < b {
		fmt.Println(c, "o'rtacha ")
	}

}
