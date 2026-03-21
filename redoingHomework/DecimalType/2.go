// Doiraning radiusi berilgan r, uning aylanasi uzunligi va yuzini hisblovchi dastur tuzing
// ( s = pi * r * r, L = 2 * pi * r) pi = 3.14

package main

import "fmt"

func main() {
	const PI float32 = 3.1413
	var raduis float32

	fmt.Scan(&raduis)

	uzunlik := 2 * PI * raduis
	yuzi := PI * raduis * raduis

	fmt.Printf("aylana uzunligi %f yuzi : %f\n", uzunlik, yuzi)

}
