// Array yoki slice elementlari ichidan eng kichigini topuvchi dastur tuzing

package main

import "fmt"

func main() {
	var arr []int = []int{2, 3, 5, 54, 548, 7, 5}
	kichik := arr[0]

	for i := 0; len(arr) > i; i++ {
		if kichik > arr[i] {
			kichik = arr[i]
		}
	}

	fmt.Println(kichik)

}
