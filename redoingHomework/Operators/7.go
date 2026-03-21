// Ikki xonal a soni berilgan! Jumlani rostlikka tekshiring: a soni polindrom va toq son

package main

import "fmt"

func main() {
	var num int = 33

	polindrome := num%11 == 0
	toq := num%2 != 0

	result := polindrome && toq

	fmt.Println(result)

}
