// 1 dan 10 gacha bo’lgan sonlar ichidan toqlarini ekranga chiqaruchi dastur tuzing!

package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
