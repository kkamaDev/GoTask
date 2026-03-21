// Berilgan gapning ichidan berilgan so’z bor yoki yo’qligini topuvchi dastur tuzing

package main

import (
	"fmt"
	"strings"
)

func main() {

	var gap string = "kamron dars qilmadi"
	var soz string = "dars"
	if qaytarilgan(gap, soz) {
		fmt.Println("bor", soz)
	} else {
		fmt.Println("yoq")
	}
}
func qaytarilgan(gapp, soz string) bool {
	solitt := strings.Split(gapp, " ")
	bormi := false

	for i := 0; i < len(solitt); i++ {
		if solitt[i] == soz {
			bormi = true
			break
		}

	}
	return bormi

}
