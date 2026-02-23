// sana : 17.02.2026
// muallif : Kamronbek

// sharti = a, b va c sonlari berilgan. Bu sonlarning eng kichigini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("a ,b , c sonlarni kiriting >> ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if a <= b && a <= c {
		fmt.Println(a, "eng kichigi ")
	} else if b <= a && b <= c {
		fmt.Println(b, "eng kichigi ")
	} else {
		fmt.Println(c, "eng kichigi ")
	}

}
