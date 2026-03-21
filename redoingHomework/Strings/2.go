// Berilgan so’z necha bo’g’indan iborat ekanligini aniqlovchi dastur tuzing

package main

import "fmt"

func main() {

	var soz string = "kamron"

	countr := 0

	bytes := []byte(soz)
	for i := 0; i < len(bytes); i++ {

		switch bytes[i] {
		case 'a', 'u', 'e', 'o', 'i':
			countr += 1

		}
	}
	fmt.Println(countr)

}
