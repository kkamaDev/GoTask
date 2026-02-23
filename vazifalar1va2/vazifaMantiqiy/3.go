//  15.02.2026
// muallif : Kamron

// sharti = a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b sonidan katta

package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("a sonini kiritng > ")
	fmt.Scan(&a)

	fmt.Print("b sonini kiritng ? ")
	fmt.Scan(&b)

	x := a > b
	fmt.Println("a soni b dan kattami > ", x)
}
