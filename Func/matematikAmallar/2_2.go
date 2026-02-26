// sharti : Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
// ( s = pi * r * r, L = 2 * pi * r) pi = 3.14
package main

import "fmt"

func main() {

	x, y := doirayuzi(5)
	fmt.Println("uzunligi :", x)
	fmt.Println("yuzi:", y)

}

func doirayuzi(R float32) (float32, float32) {

	PI := float32(3.14)
	uzunlik := 2 * PI * R
	yuza := PI * R * R
	return uzunlik, yuza

}
