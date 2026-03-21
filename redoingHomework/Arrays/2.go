// Array yoki slice elementlari yig’indisni topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var array []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]

	}
	fmt.Println(sum)

}

