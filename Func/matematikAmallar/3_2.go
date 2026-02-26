// L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing! (100sm = 1m)

package main

import "fmt"

func main() {
	result := tometr(123)

	fmt.Println(result, "m")

}

func tometr(sm float32) float32 {

	return sm / 10
}
