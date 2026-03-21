// Kiritilgan so’zni polindrom yoki polindrom emasligini aniqlovchi dastur tuzing

package main

import (
	"fmt"
)

func main() {

	var soz string = "aka"
	reverse(soz)

}

func reverse(s string) {

	left := 0
	right := len(s) - 1
	runes := []rune(s)
	polindrommi := true

	for left < right {
		if runes[left] != runes[right] {
			polindrommi = false
			break
		}
		left++
		right--

	}
	if polindrommi {
		fmt.Println("polindrome", s)
	} else {
		fmt.Println("poldrome emas", s)
	}

}
