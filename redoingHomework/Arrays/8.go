// Array yoki slice elementlari ichidagi musbat va manfiy sonlarni alohida
//  slicelarga yig’uvchi dastur tuzing topuvchi dastur tuzing

package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, -5, -8, -99, 12, 23, 4}
	musbat := []int{}
	manfiy := []int{}

	for i := 0; len(arr) > i; i++ {
		if arr[i] > 0 {
			musbat = append(musbat, arr[i])
		} else {
			manfiy = append(manfiy, arr[i])
		}
	}

	fmt.Println(musbat)
	fmt.Println(manfiy)

}
