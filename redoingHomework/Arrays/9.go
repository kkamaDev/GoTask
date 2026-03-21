// Array yoki slice e’lementlari 0 dan boshlab bittadan n gacha sonlardan iborat
// faqat bitta son tushib qoladi! Tushib qolgan sonni topuvchi dastur tuzing

package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	uzunlik := len(numbers) + 1
	sum := 0
	aslqiymat := (uzunlik * (uzunlik + 1)) / 2

	for i := 0; len(numbers) > i; i++ {
		sum += numbers[i]
	}
	result := aslqiymat - sum
	fmt.Println(result)
}
