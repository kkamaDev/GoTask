// Array yoki slice elementlari uchida polindromlarini alohida slicega yig’uvchi dastur tuzing!

package main

import "fmt"

func main() {

	var arr []int = []int{11, 25, 414, 252, 5, 77}
	newarr := []int{}

	for i := 0; i < len(arr); i++ {
		if isPolindrome(arr[i]) {
			newarr = append(newarr, arr[i])

		}
	}
	fmt.Println(newarr)
}

func isPolindrome(son int) bool {

	absoluteNume := son
	reverse := 0
	ispolindrome := false

	for son > 0 {
		qoldiq := son % 10
		reverse = reverse*10 + qoldiq
		son = son / 10
		if reverse == absoluteNume {
			ispolindrome = true
		} else {
			ispolindrome = false
		}

	}
	return ispolindrome
}
