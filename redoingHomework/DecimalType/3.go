// L uzunlik berilgan sm larda! Shu uzunlik ichidagi metrlar sonini topuvchi dastur tuzing!
//  (100sm = 1m)

package main

import "fmt"

func main() {

	var uzunlik float32

	fmt.Scan(&uzunlik)

	change := uzunlik / 100

	fmt.Println(uzunlik, change)
	// fmt.Printf("%f{sm}  %f{m}", uzunlik, change)

}
