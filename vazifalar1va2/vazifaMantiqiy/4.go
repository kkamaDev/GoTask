//	15.02.2026
//
// muallif : Kamron
// sharti = a va b sonlarni berilgan jumlani rostlikka tekshiring: a soni b soniga teng
package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("a ni kiriting : ")
	fmt.Scan(&a)

	fmt.Print("b ni kiritng : ")
	fmt.Scan(&b)

	v := a == b
	fmt.Println("a soni b ga tengmi ? ", v)
}
