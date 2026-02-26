// sharti : Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
// ( s = pi * r * r, L = 2 * pi * r) pi = 3.14

package main

import "fmt"

func main() {
	var num int = 4
	DoiraYuzi(float32(num))

}

func DoiraYuzi(R float32) {
	uzunlik := 2 * 3.14 * R
	yuzi := 3.14 * R * R
	fmt.Println("doiraning uzunligi :", uzunlik)
	fmt.Println("doiraning yuzi :", yuzi)
}
