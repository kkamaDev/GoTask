// Array yoki slice elementlari ichidan eng kattasini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	kattason := arr[0]

	for i := 0; len(arr) > i; i++ {
		if kattason < arr[i] {
			kattason = arr[i]
		}
	}
	fmt.Println(kattason)

}
