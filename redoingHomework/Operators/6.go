// a va b sonlari berilgan! Bu sonlarning yig’indisi musbatligini va toqligini aniqlovchi dastur tuzing!
// Input: 	 a = 5		Input: 	 a = 5
//  b = -4			 b = -9
// 		Output: true		Output: false

package main

import "fmt"

func main() {
	var a, b int = 5, 4
	fmt.Println(a+b > 0 && a+b%2 != 0)
}
