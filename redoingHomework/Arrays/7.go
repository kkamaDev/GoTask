// Array yoki slice elementlari ichidagi juft va toq sonlarni alohida slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newArr := []int{}
	newArr2 := []int{}

	for _, value := range arr {
		if value%2 == 0 {
			newArr = append(newArr, value)
		} else {
			newArr2 = append(newArr2, value)
		}
	}

	fmt.Println(newArr)
	fmt.Println(newArr2)

}
