// Berilgan n sonini poindrom yoki polindrom emasligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {

	var son int = 41514

	if isPolindrome(son) {
		fmt.Println(son, "polindrome")
	} else {
		fmt.Println(son, "poldrome emas")
	}

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
