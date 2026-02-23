// sana : 17.02.2026
// muallif : Kamronbek

// sharti : a va b sonlari berilgan. Bu sonlarning kattasini topuvchi dastur tuzing!

package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		fmt.Println(a, "soni ", b, "sonidan katta")
	} else if a < b {
		fmt.Println(b, "soni ", a, "sonidan katta")
	} else {
		fmt.Println("sonlar teng")
	}

}
