// // 14.02.2026
// // muallif : Kamron

// // masala = L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing!

package main

import "fmt"

func main() {
	var uzunlik float32

	fmt.Print("(sm)da uzunlik kiritng (m)ga o'tqizamiz :")
	fmt.Scan(&uzunlik)

	metr := uzunlik / 100

	fmt.Println("natija : ", metr, "metr")

}
