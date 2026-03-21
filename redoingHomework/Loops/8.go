// Berilgan n sonini tub yoki tub emasligini aniqlovchi dastur tuzing!

package main

import "fmt"

func main() {

	var son int = 20
	if istub(son) {
		fmt.Println(son, "tub")
	} else {
		fmt.Println(son, "tub emas")
	}

}

func istub(son int) bool {
	istub := false
	if son < 2 {
		return false
	}

	for i := 2; i < son+1; i++ {
		if son%i == 0 {
			istub = false
		}
	}
	return istub
}
