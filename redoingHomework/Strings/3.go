// Berilgan so’zda qatnashgan unli va undosh harflar sonini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {

	var soz string = "kamron"

	countr := 0
	countr2 := 0
	bytes := []byte(soz)
	for i := 0; i < len(bytes); i++ {

		switch bytes[i] {
		case 'a', 'u', 'e', 'o', 'i':
			countr += 1
		default:
			countr2 += 1
		}
	}
	fmt.Println("unlilar soni", countr)
	fmt.Println("undoshlar soni ", countr2)

}
