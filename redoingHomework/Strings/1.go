// berilgan so’zni birinchi harfini katta harfga o’giruvchi dastur tuzing

package main

import (
	"fmt"
)

func main() {

	var soz string = "kamron"
	bytes := []byte(soz)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = bytes[i] - 32
			break
		}
	}
	fmt.Println(string(bytes))
}
