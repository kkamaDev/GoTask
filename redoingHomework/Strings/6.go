// Berilgan so’zning barcha harflarini agar kichik harflardan iborat bo’lsa katta harfga
// aylantiruvchi aks holda kichik harflarga aylantiruvchi dastur tuzing

package main

import "fmt"

func main() {
	var soz string = "KAMRON"
	bytes := []byte(soz)

	for i := 0; i < len(bytes); i++ {

		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] = bytes[i] + 32
		} else if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = bytes[i] - 32
		}
	}

	fmt.Println(string(bytes))

}
