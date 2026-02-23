// // 14.02.2026
// // muallif : Kamron

// // Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing !

package main

import "fmt"

func main() {
	var radius float32
	var PI float32 = 3.14
	fmt.Print("aylana radiusini kiriting > ")
	fmt.Scan(&radius)

	yuzi := PI * radius
	uzunlik := 2 * PI * radius

	fmt.Println("aylana yuzasi: ", yuzi)
	fmt.Println("aylana uzunligi :", uzunlik)

}
