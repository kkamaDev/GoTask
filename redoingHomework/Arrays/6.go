// Array yoki slice elementlari ichidan juftlarini topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newArr := []int{}

	for _, value := range arr {
		if value%2 == 0 {
			newArr = append(newArr, value)
		}
	}

	fmt.Println(newArr)

}
