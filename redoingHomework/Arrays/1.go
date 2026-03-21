// Array yoki slice elementlarini teskari tartibda dastur tuzing

package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
