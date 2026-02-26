// L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing! (100sm = 1m)

package main

import "fmt"

func main() {
	ToMetr(123)

}

func ToMetr(L float32) {

	uzunlik := float32(L / 10)
	fmt.Println(uzunlik, "m")
}
