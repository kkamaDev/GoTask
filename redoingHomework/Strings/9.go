// Berilgan gapning ichidan berilgan so’z necha marta takrorlanganligini aniqlovchi dastur tuzing

package main

import (
	"fmt"
	"strings"
)

func main() {
	var gap string = "kamron uyga ketti ketti va ketti"
	var soz string = "ketti"

	sanoq := takrorson(gap, soz)

	if sanoq > 0 {
		fmt.Printf("Gapda '%s' so'zi %d marta qatnashgan\n", soz, sanoq)
	} else {
		fmt.Println("Bu so'z gapda umuman yo'q")
	}
}

func takrorson(gapp, soz string) int {
	solitt := strings.Fields(gapp)
	counter := 0

	for i := 0; i < len(solitt); i++ {
		if solitt[i] == soz {
			counter++
		}
	}
	return counter
}
