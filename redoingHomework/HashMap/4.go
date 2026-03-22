// Slicennig birinchi bo’lib ikki marta takrorlangan elementini toping

package main

import "fmt"

func main() {

	slays := []int{10, 5, 3, 4, 5, 6}

	son, topildi := BirinchiSon(slays)

	if topildi {
		fmt.Println(son)
	} else {
		fmt.Println("yo'q")
	}

}

func BirinchiSon(nums []int) (int, bool) {

	korilgan := map[int]bool{}

	for _, son := range nums {
		if korilgan[son] {
			return son, true
		}
		korilgan[son] = true
	}

	return 0, false
}
