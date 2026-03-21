// Berilgan n sonini raqamlarini teskari tartib yozishda xosil bo’lgan sonni topuvchi dastur tuzing

package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)

	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}
