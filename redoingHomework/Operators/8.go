// uch xonali a soni berilgan jumlani rostlikka tekshiring: a soni polindrom va juf son

package main

import "fmt"

func main() {

	var son int = 414

	polindrome := son/100 == son%10
	juft := son%2 == 0

	result := polindrome && juft

	fmt.Println(result)

}
